package delivery

import (
	"fmt"
	"payment-apps-backend/config"
	"payment-apps-backend/delivery/controller"
	"payment-apps-backend/manager"
	"payment-apps-backend/utils/exceptions"

	"github.com/gin-gonic/gin"
)

type Server struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initController() {
	controller.NewCustomerController(s.engine, s.useCaseManager.CustomerUseCase())
	controller.NewAuthController(s.engine, s.useCaseManager.AuthUseCase())
	controller.NewMerchantController(s.engine, s.useCaseManager.MerchantUseCase())
	controller.NewTransactionController(s.engine, s.useCaseManager.TransactionUseCase())
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	exceptions.CheckErr(err)
	infraManager, _ := manager.NewInfraManager(cfg)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)

	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		useCaseManager: useCaseManager,
		engine:         engine,
		host:           host,
	}
}
