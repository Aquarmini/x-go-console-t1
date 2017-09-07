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

func (this *DB)Query(sql string, args ...interface{}) (map[string]string) {
	stmt, err := this.Client.Prepare(sql)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmt.Close()

	rows, _ := stmt.Query(args...)
	// Get column names
	columns, err := rows.Columns()
	fmt.Println(columns)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println(len(columns))

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))
	fmt.Println(values)

	//
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	result := make(map[string]string)
	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
			result[i] = value
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return result
}
