package postgres_db

import (
	"database/sql"
	"os"
)

const (
	postgres_users_username = "postgres_users_username"
	postgres_users_password = "postgres_users_password"
	postgres_users_host     = "postgres_users_host"
	postgres_f2go_schema    = "postgres_f2go_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(postgres_users_username)
	password = os.Getenv(postgres_users_password)
	host     = os.Getenv(postgres_users_host)
	schema   = os.Getenv(postgres_f2go_schema)
)
