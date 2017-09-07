package providers

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type DB struct {
	Config *Config `inject:"config"`
	Client *sql.DB
}

func NewDB() *DB {
	fmt.Println("NewDB")
	db, _ := sql.Open("mysql", "root:910123@/phalcon")
	return &DB{Client:db}
}
