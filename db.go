package joomla

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() error {

	port := Config.GetString("port", "3306")
	host := Config.GetString("host")
	user := Config.GetString("user")
	pass := Config.GetString("password")
	dbname := Config.GetString("db")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, dbname)

	var derr error

	DB, derr = sql.Open("mysql", dsn)

	if derr != nil {
		return derr
	}

	pingErr := DB.Ping()
	
	if pingErr != nil {
		return pingErr
	}

	return nil

}

func Prefix() string {

	return Config.GetString("dbprefix")

}
