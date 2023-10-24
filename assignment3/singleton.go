package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

type DatabaseManager struct {
	db *sql.DB
}

var instance *DatabaseManager
var once sync.Once

func GetDatabaseManager() *DatabaseManager {
	once.Do(func() {
		db, err := sql.Open("sqlite3", "test.db")
		if err != nil {
			fmt.Println(err)
		}
		instance = &DatabaseManager{db}
		fmt.Println("Creating single instance now.")
	})
	fmt.Println("Single instance already created.")
	return instance
}

func main() {
	for i := 0; i < 3; i++ {
		go GetDatabaseManager()
	}

	fmt.Scanln()
}
