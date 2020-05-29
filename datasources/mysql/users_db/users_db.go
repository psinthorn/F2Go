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
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_f2go_schema    = "mysql_f2go_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	schema   = os.Getenv(mysql_f2go_schema)
)

func init() {

	envPort := configs.Server.PortRunning("8089")

	switch envPort {
	case "8089":
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
