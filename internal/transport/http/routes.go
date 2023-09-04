package http

import (
	"fmt"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/transport/http/handler"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/transport/http/middleware"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/log"
	"go.uber.org/zap/zapcore"
)

func (s *Server) register(secret string, userService protocol.User,
	bankService protocol.Bank,
	bankBranchService protocol.BankBranch,
	financialCardService protocol.FinancialCard,
	currencyService protocol.Currency,
) {

	logConfig := log.Config{
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		DisableStacktrace: false,
		Level:             zapcore.DebugLevel,
	}
	logger, err := log.New("bank-branch-handler", logConfig)
	if err != nil {
		// Handle error, possibly exit the application
		fmt.Println("Could not initialize logger:", err)
		return
	}

	bankBranchHandler := handler.NewBranchHandler(logger, bankBranchService)
	financialCardHandler := handler.NewFinancialCardHandler(logger, financialCardService)
	currencyHandler := handler.NewCurrencyHandler(logger, currencyService)

	auth := s.echo.Group("/auth")
	auth.POST("/sign-up", handler.SignUpHandler(userService))
	auth.POST("/sign-in", handler.SignInHandler(userService))
	auth.POST("/refresh", handler.RefreshTokenHandler(userService), middleware.JWT(secret))

	user := s.echo.Group("/account", middleware.JWT(secret))
	user.GET("profile", handler.GetProfileHandler(userService))
	user.PUT("profile", handler.EditProfileHandler(userService))

	bank := s.echo.Group("/bank", middleware.JWT(secret))
	bank.POST("/register", handler.RegisterBankHandler(bankService))
	bank.GET("/id/:id", handler.GetBankByIDHandler(bankService))
	bank.GET("/code/:code", handler.GetBankByCodeHandler(bankService))
	bank.GET("/name/:name", handler.GetBankByNameHandler(bankService))
	bank.PUT("/update", handler.UpdateBankDetailsHandler(bankService))
	bank.GET("/list", handler.ListAllBanksHandler(bankService))
	bank.GET("/status/:status", handler.ListBanksByStatusHandler(bankService))

	// Adding new group for bank-branch
	branch := s.echo.Group("/branch", middleware.JWT(secret))
	branch.POST("/add", bankBranchHandler.AddBranchHandler(bankBranchService))
	branch.GET("/id/:id", bankBranchHandler.GetBranchByIDHandler(bankBranchService))
	branch.GET("/name/:name", bankBranchHandler.GetBranchByNameHandler(bankBranchService))
	branch.GET("/code/:code", bankBranchHandler.GetBranchByCodeHandler(bankBranchService))
	branch.PUT("/update", bankBranchHandler.UpdateBranchHandler(bankBranchService))
	branch.DELETE("/delete/:id", bankBranchHandler.DeleteBranchHandler(bankBranchService))
	branch.GET("/list", bankBranchHandler.ListAllBranchesHandler(bankBranchService))
	branch.GET("/status/:status", bankBranchHandler.ListBranchesByStatusHandler(bankBranchService))
	branch.GET("/listByBank/:id", bankBranchHandler.ListBranchesByBankIDHandler(bankBranchService))

	// Adding new group for financial-card
	card := s.echo.Group("/card", middleware.JWT(secret))
	card.POST("/register", financialCardHandler.RegisterCardHandler)
	card.PUT("/update", financialCardHandler.UpdateCardHandler)
	card.DELETE("/delete/:id", financialCardHandler.DeleteCardHandler)
	card.GET("/id/:id", financialCardHandler.GetCardByIDHandler)
	card.GET("/listByAccount/:id", financialCardHandler.ListCardsByAccountIDHandler)
	card.GET("/listByType/:type", financialCardHandler.ListCardsByTypeHandler)

	// Adding new group for currency
	currency := s.echo.Group("/currency", middleware.JWT(secret))
	currency.POST("/add", currencyHandler.AddCurrencyHandler)
	currency.PUT("/update", currencyHandler.UpdateCurrencyHandler)
	currency.DELETE("/delete/:id", currencyHandler.DeleteCurrencyHandler)
	currency.GET("/id/:id", currencyHandler.GetCurrencyByIDHandler)
	currency.GET("/name/:name", currencyHandler.GetCurrencyByNameHandler)
	currency.GET("/list", currencyHandler.ListCurrenciesHandler)
	currency.GET("/exchangeRate/:fromCode/:toCode", currencyHandler.GetExchangeRateHandler)
	currency.PUT("/bulkUpdateExchangeRates", currencyHandler.BulkUpdateExchangeRatesHandler)
	currency.GET("/search/:query", currencyHandler.SearchCurrenciesHandler)
	currency.GET("/convert/:fromCode/:toCode/:amount", currencyHandler.ConvertAmountHandler)
	currency.GET("/compare/:firstCode/:secondCode", currencyHandler.CompareCurrenciesHandler)
	currency.GET("/trends/:code/:duration", currencyHandler.GetCurrencyTrendsHandler)
	currency.GET("/strongest", currencyHandler.GetStrongestCurrencyHandler)
	currency.GET("/weakest", currencyHandler.GetWeakestCurrencyHandler)
	currency.PUT("/notifyUsersOnExchangeRateChange/:threshold", currencyHandler.NotifyUsersOnExchangeRateChangeHandler)
	currency.GET("/countries/:code", currencyHandler.GetCountriesUsingCurrencyHandler)

}
