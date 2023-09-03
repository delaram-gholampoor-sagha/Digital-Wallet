package http

import (
	"context"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/config"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type (
	Server struct {
		echo    *echo.Echo
		address string
	}

	ServerConfig struct {
		Config               config.HTTP
		Logger               *zap.SugaredLogger
		UserService          protocol.User
		BankService          protocol.Bank
		BankBranchService    protocol.BankBranch
		FinancialCardService protocol.FinancialCard
		JWTSecret            string
	}
)

func New(sc ServerConfig) *Server {
	e := echo.New()

	// Echo config
	e.HideBanner = true
	e.Server.ReadTimeout = sc.Config.ReadTimeout
	e.Server.WriteTimeout = sc.Config.WriteTimeout
	e.Server.IdleTimeout = sc.Config.IdleTimeout
	e.Server.ErrorLog = zap.NewStdLog(sc.Logger.Desugar())
	e.HTTPErrorHandler = errorHandler(sc.Logger)

	// Middlewares
	e.Use(echomw.RecoverWithConfig(echomw.RecoverConfig{
		StackSize:         sc.Config.Recover.StackSize << 10,
		DisableStackAll:   sc.Config.Recover.DisableStackAll,
		DisablePrintStack: sc.Config.Recover.DisablePrintStack,
	}))
	e.Use(echomw.TimeoutWithConfig(echomw.TimeoutConfig{
		Timeout: sc.Config.Timeout,
	}))
	e.Use(echomw.CORSWithConfig(echomw.CORSConfig{
		AllowOrigins:     sc.Config.CORS.AllowedOrigins,
		AllowMethods:     sc.Config.CORS.AllowedMethods,
		AllowHeaders:     sc.Config.CORS.AllowedHeaders,
		AllowCredentials: sc.Config.CORS.AllowCredentials,
		ExposeHeaders:    sc.Config.CORS.ExposedHeaders,
		MaxAge:           sc.Config.CORS.MaxAge,
	}))
	e.Use(echomw.BodyLimit(sc.Config.BodyLimitSize))

	server := &Server{
		echo:    e,
		address: sc.Config.Address,
	}

	server.register(
		sc.JWTSecret,
		sc.UserService,
		sc.BankService,
		sc.BankBranchService,
		sc.FinancialCardService,
	)

	return server
}

func (s *Server) Start() error {
	return s.echo.Start(s.address)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

func (s *Server) Close() error {
	return s.echo.Close()
}
