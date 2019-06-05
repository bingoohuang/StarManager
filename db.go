package main

import (
	"database/sql"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

func FixMySQLTableOptions(db *gorm.DB) *gorm.DB {
	return db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")
}

func SetConnectionPool(db *sql.DB) {
	// 1. https://making.pusher.com/production-ready-connection-pooling-in-go/
	// 2. http://go-database-sql.org/connection-pool.html
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(10 * time.Second)
}

func IsMySQLDriver(dbDriver string) bool {
	return dbDriver == "mysql"
}

func FixMySQLURIParameters(dbURI string) string {
	// user:pass@tcp(192.168.136.90:3307)/db?charset=utf8mb4&parseTime=true&loc=Local
	// refer
	// 1. https://github.com/go-sql-driver/mysql
	// 2. https://gorm.io/docs/connecting_to_the_database.html
	// 3. https://stackoverflow.com/questions/40527808/setting-tcp-timeout-for-sql-connection-in-go
	attachParameters := AttachParameter(dbURI, "charset", "utf8mb4")
	attachParameters += AttachParameter(dbURI, "parseTime", "true")
	attachParameters += AttachParameter(dbURI, "loc", "Local")
	attachParameters += AttachParameter(dbURI, "timeout", "10s")
	attachParameters += AttachParameter(dbURI, "writeTimeout", "10s")
	attachParameters += AttachParameter(dbURI, "readTimeout", "10s")

	if attachParameters != "" && !strings.Contains(dbURI, "?") {
		attachParameters = "?" + attachParameters[1:]
	}

	return dbURI + attachParameters
}

func AttachParameter(dbURI, key, value string) string {
	if strings.Contains(dbURI, key+"=") {
		return ""
	}

	return "&" + key + "=" + value
}
