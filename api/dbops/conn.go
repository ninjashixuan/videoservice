package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	connDB *sql.DB
	err    error
)

func init() {
	connDB, err = sql.Open("mysql", "root:zsxzsx669094@tcp(127.0.0.1:3306)/video?parseTime=true&loc=Local")
	if err != nil {
		panic(err.Error())
	}
}



