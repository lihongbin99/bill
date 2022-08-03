package utils

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)
}

func init() {
	sql.Register("mysql", &mysql.MySQLDriver{})
}

func test() {
	db, err := sql.Open("mysql", "root:root@tcp(ip:3306)/database?charset=utf-8")
}
