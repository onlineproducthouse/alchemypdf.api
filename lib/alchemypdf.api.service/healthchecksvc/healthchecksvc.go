package healthchecksvc

import (
	"context"
	"fmt"

	"alchemypdf.api/lib/alchemypdf.api.infrastructure/loglocal"
	"github.com/jackc/pgx/v4/pgxpool"
)

type HealthCheckService struct {
	logger loglocal.ILogger
	db     *pgxpool.Pool
}

type IHealthCheckService interface {
	Ping() error
}

func NewHealthCheckService(logger loglocal.ILogger, db *pgxpool.Pool) HealthCheckService {
	return HealthCheckService{logger, db}
}

func (svc HealthCheckService) Ping() error {
	const op string = "HealthCheckService.Ping"

	svc.logger.Debug(fmt.Sprintf(`start: "%s"`, op))

	svc.logger.Debug("healthcheck:db")
	if err := svc.db.Ping(context.Background()); err != nil {
		svc.logger.Debug("healthcheck:db:not-okay")
		svc.logger.Fatal(err.Error())
	}
	svc.logger.Debug("healthcheck:db:okay")

	return nil
}
