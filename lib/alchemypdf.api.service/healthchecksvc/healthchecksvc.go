package healthchecksvc

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/onlineproducthouse/alchemypdf.api.logger/loggylog"
)

type HealthCheckService struct {
	logger loggylog.ILoggyLog
	db     *pgxpool.Pool
}

type IHealthCheckService interface {
	Ping() error
}

func NewHealthCheckService(logger loggylog.ILoggyLog, db *pgxpool.Pool) HealthCheckService {
	return HealthCheckService{logger, db}
}

func (svc HealthCheckService) Ping() error {
	const op string = "HealthCheckService.Ping"

	if err := svc.db.Ping(context.Background()); err != nil {
		svc.logger.Debug("healthcheck:db:not-okay")
		svc.logger.Fatal(err.Error())
	}

	svc.logger.Debug("healthcheck:db:okay")

	return nil
}
