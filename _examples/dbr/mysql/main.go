package main

import (
	"errors"
	"github.com/gocraft/dbr"
	//"go.knocknote.io/octillery"
	"go.knocknote.io/octillery/connection/adapter"
	"go.knocknote.io/octillery/connection/adapter/plugin"
)

type Member struct {
	ID      int64  `db:"id"`
	Number  int64  `db:"number"`
	Name    string `db:"name"`
	IsValid bool   `db:"is_valid"`
}

const SQL_CREATE = `
CREATE TABLE IF NOT EXISTS members(
    id integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
    number integer NOT NULL,
    name varchar(255),
    is_valid tinyint(1) NOT NULL
)`

func main() {
	adapter.Register("mysql", &plugin.MySQLAdapter{})
	//debug.SetDebug(true)
	//if err := octillery.LoadConfig(filepath.Join(path.ThisDirPath(), "databases.yml")); err != nil {
	//	panic(err)
	//}
	conn, err := dbr.Open("mysql", "root:root@tcp(localhost:13306)/test", nil)
	if err != nil {
		panic(err)
	}
	sess := conn.NewSession(nil)
	if conn.DB != nil {
		if _, err := conn.DB.Exec(SQL_CREATE); err != nil {
			panic(err)
		}
	}
	if _, err := sess.DeleteFrom("members").Exec(); err != nil {
		panic(err)
	}

	result, err := sess.InsertInto("members").
		Columns("id", "number", "name", "is_valid").
		Values(1, 10, "Bob", true).
		Exec()
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	if count != 1 {
		panic(errors.New("cannot insert row"))
	}

	member := &Member{2, 9, "Ken", false}

	sess.InsertInto("members").
		Columns("id", "number", "name", "is_valid").
		Record(member).
		Exec()

	var members []Member
	sess.Select("*").From("members").Load(&members)

	if len(members) != 2 {
		panic(errors.New("cannot get members"))
	}

	attrsMap := map[string]interface{}{"number": 13, "name": "John"}
	if _, err := sess.Update("members").
		SetMap(attrsMap).
		Where("id = ?", members[0].ID).
		Exec(); err != nil {
		panic(err)
	}

	var m Member
	if _, err := sess.Select("*").
		From("members").
		Where("id = ?", members[0].ID).
		Load(&m); err != nil {
		panic(err)
	}

	if m.Name != "John" {
		panic(errors.New("cannot update row"))
	}
}
