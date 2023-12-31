package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/config"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/database/postgres"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/repository"
	accounttransaction "github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/service/account_transaction"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/service/bank"
	bankbranch "github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/service/bank_branch"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/service/currency"
	financialaccount "github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/service/financial_account"
	financialcard "github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/service/financial_card"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/service/user"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/transport/http"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/log"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/utils"
	"github.com/urfave/cli/v2"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
)

var (
	apiCommand = &cli.Command{
		Name:        "api",
		Description: "serving http api",
		Action:      api,
	}

	httpServer protocol.HTTP
)

func api(_ *cli.Context) (err error) {
	// App Starting
	fmt.Println("starting service")
	defer fmt.Println("shutdown complete")

	fmt.Println("loading application config")
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("loading application config : %w", err)
	}

	// GOMAXPROCS
	// Set the correct number of threads for the service
	// based on what is available either by the machine or quotas.
	if _, err := maxprocs.Set(); err != nil {
		return fmt.Errorf("cant set maxprocs: %w", err)
	}

	fmt.Printf("GOMAXPROCS = %d \n", runtime.GOMAXPROCS(0))

	fmt.Println("initial logger")
	logger, err := log.New("digital-wallet", log.Config{
		OutputPaths:       cfg.Logger.OutputPaths,
		ErrorOutputPaths:  cfg.Logger.ErrorOutputPaths,
		DisableStacktrace: cfg.Logger.DisableStacktrace,
		Level:             cfg.Logger.Level,
	})
	if err != nil {
		return fmt.Errorf("initial log: %w", err)
	}

	defer func(logger *zap.SugaredLogger) {
		fmt.Println("syncing logger start")
		defer fmt.Println("syncing logger complete")
		derr := logger.Sync()
		if derr != nil && err != nil {
			err = fmt.Errorf("sync logger: %w", err)
		}
	}(logger)

	// Create connectivity to the database.
	fmt.Println("connecting to postgres databases...")
	postgresDB, err := postgres.New(cfg.Postgres)
	if err != nil {
		return fmt.Errorf("connecting to database: %w", err)
	}
	fmt.Println("connected to postgres databases")

	defer func() {
		fmt.Println("closing postgres connections start")
		defer fmt.Println("closing postgres connections complete")

		if derr := postgresDB.Close(); derr != nil && err == nil {
			err = fmt.Errorf("closing postgres connections: %w", err)
		}
	}()

	userRepo := repository.NewUser(postgresDB)
	bankRepo := repository.NewBank(postgresDB)
	bankBranchRepo := repository.NewBankBranch(postgresDB)
	financialCardRepo := repository.NewFinancialCard(postgresDB)
	currencyRepo := repository.NewCurrency(postgresDB)
	financialAccountRepo := repository.NewFinancialAccount(postgresDB)
	accountTransactionRepo := repository.NewAccountTransaction(postgresDB)

	// Create instances of BcryptHasher and JWTTokenGenerator
	hasher := utils.BcryptHasher{}
	tokenGenerator := utils.JWTTokenGenerator{}
	userService := user.New(cfg.JWT, logger, userRepo, hasher, tokenGenerator)
	bankService := bank.New(cfg.JWT, logger, bankRepo, tokenGenerator)
	currencyService := currency.New(cfg.JWT, logger, tokenGenerator, currencyRepo)
	bankBranchService := bankbranch.New(cfg.JWT, logger, bankBranchRepo, tokenGenerator, bankService)
	financialAccountService := financialaccount.New(cfg.JWT, logger,
		tokenGenerator,
		financialAccountRepo,
		bankService,
		bankBranchService,
		userService,
		currencyService)
	accountTransactionService := accounttransaction.New(cfg.JWT, logger, tokenGenerator, financialAccountService, accountTransactionRepo)
	financialCardService := financialcard.New(cfg.JWT, logger, financialCardRepo, tokenGenerator, financialAccountService)

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// server init
	serverConfig := http.ServerConfig{
		Logger:               logger,
		Config:               cfg.HTTP,
		UserService:          userService,
		BankService:          bankService,
		BankBranchService:    bankBranchService,
		FinancialCardService: financialCardService,
		CurrencyService:      currencyService,
		FinancialAccount:     financialAccountService,
		AccountTransaction:   accountTransactionService,
	}
	httpServer = http.New(serverConfig)

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	httpServerErrors := make(chan error, 1)
	go func() {
		httpServerErrors <- httpServer.Start()
	}()

	select {
	case err := <-httpServerErrors:
		return fmt.Errorf("http server error: %w\n", err)
	case <-shutdown:
		fmt.Println("application shutdown start")

		fmt.Println("http server shutdown start")
		defer fmt.Println("http server shutdown complete")
		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), cfg.HTTP.ShutdownTimeout)
		defer cancel()

		// Asking listener to shut down and shed load.
		if err := httpServer.Shutdown(ctx); err != nil {
			func(httpServer protocol.HTTP) {
				err := httpServer.Close()
				if err != nil {
					logger.Errorw("closing http server", "error", err)
				}
			}(httpServer)

			logger.Errorw("could not stop server gracefully: %w", err)
			return err
		}
	}

	return nil
}
