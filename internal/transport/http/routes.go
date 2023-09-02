package http

import (
	"fmt"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/transport/http/handler"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/transport/http/middleware"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/log"
	"go.uber.org/zap/zapcore"
)

func (s *Server) register(secret string, userService protocol.User, bankService protocol.Bank, bankBranchService protocol.BankBranch) {

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

}
