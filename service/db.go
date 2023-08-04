package service

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	UserName     string = "root"
	Password     string = "passw0rd"
	Address      string = "localhost"
	Port         int    = 3306
	Database     string = "GoFirst"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

func Context() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb&parseTime=True", UserName, Password, Address, Port, Database)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("mysql conn fail")
	}

	context, err := conn.DB()
	if err != nil {
		panic("get db fail")
	}

	context.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	context.SetMaxIdleConns(MaxIdleConns)
	context.SetMaxOpenConns(MaxOpenConns)

	return context
}
