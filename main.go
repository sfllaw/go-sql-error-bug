package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var ErrInvalid = errors.New("not an ASCII string")

type ASCII []byte

func (a ASCII) Scan(src interface{}) error {
	// Elide code that parses src into a.
	return ErrInvalid
}

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var a ASCII
	if err := db.QueryRow(`SELECT "世界"`).Scan(&a); err != nil {
		if !errors.Is(err, ErrInvalid) {
			panic(fmt.Errorf("unexpected error: %w", err))
		}
		fmt.Printf("err is ErrInvalid: %v\n", err)
	}
}
