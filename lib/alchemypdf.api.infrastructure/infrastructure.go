package infrastructure

import (
	"context"
	"fmt"

	"alchemypdf.api/lib/alchemypdf.api.infrastructure/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/onlineproducthouse/alchemypdf.api.logger/loggylog"
)

type (
	IInfrastructure interface {
		Config() config.TConfig
		Logger() loggylog.LoggyLog
		DBConn() *pgxpool.Pool
	}

	Infrastructure struct {
		config config.TConfig
		logger loggylog.LoggyLog
		dbconn *pgxpool.Pool
	}
)

func NewInfrastructure() Infrastructure {
	cfg := config.Config()
	logr := loggylog.New("Info", 0)
	dbConn := db(cfg, logr)

	return Infrastructure{
		config: cfg,
		logger: logr,
		dbconn: dbConn,
	}
}

func db(c config.IConfig, l loggylog.ILoggyLog) *pgxpool.Pool {
	l.Info("opening connection to database")

	connStr := c.DbConnectionString()
	// if c.EnvName() != constant.ENV_LOCAL {
	// 	path, err := os.Getwd()
	// 	if err != nil {
	// 		l.Fatal(fmt.Sprintf("%v", err))
	// 	}

	// 	connStr = fmt.Sprintf("%s?%s%s%s", connStr, "ssl=true&sslmode=verify-ca&sslrootcert=", path, "/root.cert")
	// }

	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		l.Fatal(fmt.Sprintf("%v", err))
	}

	if err := db.Ping(context.Background()); err != nil {
		l.Fatal(err.Error())
	}

	l.Info("connection opened to database")
	return db
}

func (infra Infrastructure) Config() config.TConfig {
	return infra.config
}

func (infra Infrastructure) Logger() loggylog.LoggyLog {
	return infra.logger
}

func (infra Infrastructure) DBConn() *pgxpool.Pool {
	return infra.dbconn
}
