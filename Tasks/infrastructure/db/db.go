package db

import (
	"context"
	"database/sql"
	"fmt"
	"knovel/tasks/domain/common"
	"knovel/tasks/infrastructure/config"
	"time"
)

type database struct {
	pg *sql.DB
}

var _ common.DbContext = (*database)(nil)

func NewDBContext() common.DbContext {
	pgCfg := config.GetConfig().Db
	// build connect string:
	connect := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=%s", pgCfg.Host, pgCfg.Port, pgCfg.User, pgCfg.Password, pgCfg.DB, pgCfg.SearchPath, pgCfg.SSLMode)
	pg, err := sql.Open("pgx", connect)
	if err != nil {
		panic(fmt.Sprintf("could not connect database %v", err.Error()))
	}
	pg.SetMaxOpenConns(pgCfg.MaxOpenConn)
	pg.SetMaxIdleConns(pgCfg.MaxIdleConn)
	pg.SetConnMaxLifetime(time.Duration(pgCfg.ConnLifeTime) * time.Minute)
	return &database{pg: pg}
}

func (db *database) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return db.pg.ExecContext(ctx, query, args...)
}

func (db *database) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return db.pg.QueryContext(ctx, query, args...)
}

func (db *database) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return db.pg.QueryRowContext(ctx, query, args...)
}
