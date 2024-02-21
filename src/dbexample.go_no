package main

import (
  "database/sql"
  "fmt"
  
  _ "github.com/lib/pq"
  "github.com/jmoiron/sqlx"
)

const (
    host     = "golang_postgres"
    port     = 5432
    user     = "postgres"
    password = "qwerty"
    dbname   = "postgres"
)

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()
  
    err = db.Ping()
    if err != nil {
        panic(err)
    }
  
    fmt.Println("Successfully connected!")
    
    type ttt struct {
      myname string `db:"name"`      
    }

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