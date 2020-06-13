package config

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// NewDB DBと接続する
func NewDB() *sql.DB {
	DbConnection, err := sql.Open("sqlite3", "./server.sql")

	// fmt.Println("NewDB DbConnection", DbConnection)

	if err != nil {
		fmt.Println("NewDB err", err)
	}
	// defer DbConnection.Close()

	return DbConnection
}
