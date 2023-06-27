package main

import (
	"fmt"
	"go.knocknote.io/octillery"
	"go.knocknote.io/octillery/connection/adapter"
	"go.knocknote.io/octillery/connection/adapter/plugin"
	"go.knocknote.io/octillery/database/sql"
	"go.knocknote.io/octillery/debug"
	"go.knocknote.io/octillery/path"
	_ "go.knocknote.io/octillery/plugin"
	"path/filepath"
)

type Member struct {
	ID      int64  `db:"id"`
	Number  int64  `db:"number"`
	Name    string `db:"name"`
	IsValid bool   `db:"is_valid"`
}

const SQL_CREATE = `
CREATE TABLE IF NOT EXISTS members(
    id integer NOT NULL PRIMARY KEY,
    number integer NOT NULL,
    name varchar(255),
    is_valid tinyint(1) NOT NULL
)`

func main() {
	adapter.Register("mysql", &plugin.MySQLAdapter{})
	debug.SetDebug(true)
	if err := octillery.LoadConfig(filepath.Join(path.ThisDirPath(), "databases.yml")); err != nil {
		panic(err)
	}
	db, err := sql.Open("mysql", "")
	if err != nil {
		panic(err)
	}
	if db != nil {
		if _, err := db.Exec(SQL_CREATE); err != nil {
			panic(err)
		}
	}

	result, err := db.Exec("select * from product limit 1")
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
