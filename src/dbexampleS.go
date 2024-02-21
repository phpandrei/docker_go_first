package main

import (
  "database/sql"
  "fmt"
  "log"
  "github.com/lib/pq"
  "github.com/jmoiron/sqlx"
)

const (
    host     = "golang_postgres"
    port     = 5432
    user     = "postgres"
    password = "qwerty"
    dbname   = "postgres"
)

var schema = `
CREATE TABLE ak (
    name text
)`

type Myname struct {
    Name sql.NullString `db:"name"` 
}

func main() {
    db, err := sqlx.Connect("postgres", "host=golang_postgres user=postgres dbname=postgres password=qwerty port=5432 sslmode=disable")
    if err != nil {
        log.Fatalln(err)
    }
  
    //db.MustExec(schema)

    fmt.Println("Successfully connected!")

    rows, err := db.Queryx("SELECT name FROM ak")
    for rows.Next() {
        myname := &Myname{}

        err := rows.StructScan(&myname)
        if err != nil {
            log.Fatalln(err)
        } 
        fmt.Printf("%#v\n", myname.Name)
        fmt.Printf("%#v\n", myname)
        //fmt.Printf("%#v\n", rows.)
    }

    //rows := []ttt{}
    //db.Select(&rows, "SELECT name FROM ak")
    /*myttt := Myname{}
    rows, err := db.Queryx("SELECT * FROM ak")
    for rows.Next() {
        err := rows.StructScan(&myttt)
        if err != nil {
            log.Fatalln(err)
        } 
        fmt.Printf("%#v\n", myttt)
    }*/
    /*
    //rows, err := db.Queryx("SELECT name FROM ak WHERE age = ?", 1)
    rows, err := db.Queryx("SELECT name FROM ak")
    if err != nil {
        log.Fatal(err)
    }
        
    // defer rows.Close()
    
    names := make([]string, 0)
    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
            log.Fatal(err)
        }
        names = append(names, name)
    }
    
    // Check for errors from iterating over rows.
    if err := rows.Err(); err != nil {
      log.Fatal(err)  
    }*/
  }