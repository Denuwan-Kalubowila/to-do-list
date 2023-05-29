package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type MysqlDb struct {
	Username          string
	Password          string
	Host              string
	Port              string
	Database          string
	MaxOpenConnection int
	MaxIdleConnection int
}

func ConnectToMysql() error {
	var err error
	var database MysqlDb

	database.Username = os.Getenv("USER_NAME")
	database.Password = os.Getenv("PASSWORD")
	database.Host = os.Getenv("HOST")
	database.Port = os.Getenv("DB_PORT")
	database.Database = os.Getenv("DATABASE")
	database.MaxIdleConnection = 10
	database.MaxIdleConnection = 5

	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", database.Username, database.Password, database.Host, database.Port, database.Database))
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(database.MaxOpenConnection)
	db.SetMaxIdleConns(database.MaxIdleConnection)

	_, err = db.Conn(context.Background())
	fmt.Println("Connected.")

	return err
}

func GetDbConnection() *sql.DB {
	if db == nil {
		var err error
		var database MysqlDb

		database.Username = os.Getenv("USER_NAME")
		database.Password = os.Getenv("PASSWORD")
		database.Host = os.Getenv("HOST")
		database.Port = os.Getenv("DB_PORT")
		database.Database = os.Getenv("DATABASE")
		database.MaxIdleConnection = 10
		database.MaxIdleConnection = 5

		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", database.Username, database.Password, database.Host, database.Port, database.Database))
		if err != nil {
			return nil
		}

		db.SetMaxOpenConns(database.MaxOpenConnection)
		db.SetMaxIdleConns(database.MaxIdleConnection)
		return db
	} else {
		return db
	}
}
