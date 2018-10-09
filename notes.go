package main

import (
	"fmt"
	"log"
	"time"
	_ "github.com/mattn/go-sqlite3"
	
)

var Notes *[]Note
var title string

type Note struct{
	Title string `json:"title"`
	Cont string `json:"content"`
}

// Update updates note content 
func (nt Note) Update(_title string,_cont string){
	nt.Title = _title
	nt.Cont = _cont
}

func (nt Note) Save(){
	tx, err := DB.Begin()
	if err != nil {
		fmt.Printf("tx %v",err)
	}
	stmt, err := tx.Prepare("insert into Notes(title, content) values(?, ?)")
	if err != nil {
		fmt.Printf("tx %v",err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(nt.Title,nt.Cont)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}

func (nt Note)GetAll(){
	rows, err := DB.Query("select id, title, content from Notes")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var content string
		err = rows.Scan(&id, &title, &content)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d | %s:\n",id, title)
		fmt.Println(content)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
func getDate() string{
	currentTime := time.Now()
	return currentTime.String()
}