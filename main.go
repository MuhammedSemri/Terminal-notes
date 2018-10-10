package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

// file where the notes are saved
var memPath string
var dir string
var DB *sql.DB

func init() {
	
	initDB()
}

func main() {
	initCmd()
	defer DB.Close()
}

func initCmd(){
	
	RootCmd.AddCommand(TitleCmd,ContCmd,GetAllCmd, DeleteIDCmd,DelteAllCmd)
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	
}

func initDB(){
	db, err := sql.Open("sqlite3", "./mem.db")
	if err != nil {
		log.Println(err)
	}
	
	sqlStmt := `
	create table IF NOT EXISTS Notes (id integer not null primary key AUTOINCREMENT, title text, content text NOT NULL);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
	DB = db
}