package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/tecolotedev/stori_back/config"
	"github.com/tecolotedev/stori_back/db/sqlc_code"
)

var Queries *sqlc_code.Queries

func InitDb() {
	ctx := context.Background()

	sslConn := "verify-full"

	if config.EnvVars.IS_LOCAL {
		sslConn = "disable"
	}

	configConn := fmt.Sprintf("host=%s password=%s user=%s dbname=%s sslmode=%s",
		config.EnvVars.DB_HOST,
		config.EnvVars.DB_PASSWORD,
		config.EnvVars.DB_USER,
		config.EnvVars.DB_NAME,
		sslConn,
	)

	conn, err := pgx.Connect(ctx, configConn)
	if err != nil {
		log.Fatal(err)
	}

	Queries = sqlc_code.New(conn)
}
