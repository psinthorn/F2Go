package users_db

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/psinthorn/F2Go/configs"
)

const (
	// mysql_users_username = "mysql_users_username"
	// mysql_users_password = "mysql_users_password"
	// mysql_users_host     = "mysql_users_host"
	// mysql_f2go_schema    = "mysql_f2go_schema"

	CLEARDB_DATABASE_URL    = "CLEARDB_DATABASE_URL"
	CLEARDB_DATABASE_USER   = "CLEARDB_DATABASE_USER"
	CLEARDB_DATABASE_PASS   = "CLEARDB_DATABASE_PASS"
	CLEARDB_DATABASE_SCHEMA = "CLEARDB_DATABASE_SCHEMA"
)

var (
	Client *sql.DB

	// // Local host
	// username = os.Getenv(mysql_users_username)
	// password = os.Getenv(mysql_users_password)
	// host         = os.Getenv(mysql_users_host)
	// schema = os.Getenv(mysql_f2go_schema)

	// // ClearDB
	username = os.Getenv(CLEARDB_DATABASE_USER)
	password = os.Getenv(CLEARDB_DATABASE_PASS)
	host     = os.Getenv(CLEARDB_DATABASE_URL)
	schema   = os.Getenv(CLEARDB_DATABASE_SCHEMA)
)

func init() {

	envPort := configs.Server.PortRunning("8089")

	switch envPort {
	case "8089":

		// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		// 	username, password, host, schema,
		// )

		// ClearDB
		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
			username, password, host, schema,
		)

		//fmt.Println(dataSourceName)
		var err error
		Client, err = sql.Open("mysql", dataSourceName)
		if err != nil {
			panic(err)
		}
		if err = Client.Ping(); err != nil {
			panic(err)
		}
		log.Println("MySQL is successfully figured")
		return
	}
}
