package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

var conn *sql.DB

func getC

func Insert(cities []Town) {
    conn, err := sql.Open("mysql", "scott:tiger@/poke")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()
    sqlStr := "INSERT INTO CITY(web_name, url) VALUES "
    vars := []interface{}{}
    for _, city := range cities {
        sqlStr += "(?, ?),"
        vars = append(vars, city.Name, city.Url)
    }
    sqlStr = sqlStr[0:len(sqlStr)-1]
    stmt, _ := conn.Prepare(sqlStr)
    result, err := stmt.Exec(vars...)
    if err != nil {
        fmt.Println(err)
        return
    }
    nRow, err := result.RowsAffected()
    fmt.Println(nRow, " rows inserted")
}

func getAllTowns() []Town {

}
