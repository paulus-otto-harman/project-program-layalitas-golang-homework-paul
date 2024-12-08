package infra

import (
	"go.uber.org/zap"
	"homework/config"
	"homework/database"
	"homework/handler"
	"homework/log"
	"homework/repository"
	"homework/service"
)

type ServiceContext struct {
	Cfg config.Config
	Ctl handler.Handler
	Log *zap.Logger
}

func NewServiceContext(migrateDb bool, seedDb bool) (*ServiceContext, error) {
	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	// instance config
	appConfig, err := config.LoadConfig(migrateDb, seedDb)
	if err != nil {
		return handlerError(err)
	}

	// instance looger
	appLogger, err := log.InitZapLogger(appConfig)
	if err != nil {
		return handlerError(err)
	}

	// instance database
	db, err := database.ConnectDB(appConfig)
	if err != nil {
		return handlerError(err)
	}

	// instance repository
	repositories := repository.NewRepository(db)

	// instance service
	services := service.NewService(repositories)

	// instance controller
	Ctl := handler.NewHandler(services, appLogger)

	return &ServiceContext{Cfg: appConfig, Ctl: *Ctl, Log: appLogger}, nil
}
