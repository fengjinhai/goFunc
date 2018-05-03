package main

import (
	"database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
)

var (
    dbhostip = "127.0.0.1:3306"
    dbusername = "root"
    dbpassword = "123456"
    dbname = "test"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func insert() {
    db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/Test?charset=utf8")
    checkErr(err)
    stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
    checkErr(err)
    res, err := stmt.Exec("fjh", "research", "2015")
    checkErr(err)
    id, err := res.LastInsertId()
    checkErr(err)
    fmt.Println(id)
    db.Close()
}

func update() {
    db, err := sql.Open("mysql", "root:@/test?charset=utf-8")
    checkErr(err)
    stmt, err = db.Prepare("update userinfo set username=? where uid=?")
    checkErr(err)
    res, err = stmt.Exec("astaxieupdate", id)
    checkErr(err)
    affect, err := res.RowsAffected()
    fmt.Println(affect)
    db.Close()
}

func select() {
    db, err := sql.Open("mysql", "root:@/test?charset=utf-8")
    checkErr(err)
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)
    for rows.Next() {
        var uid int
        var username string
        err = rows.Scan(&uid, &username)
        check(err)
        fmt.Println(uid)
        fmt.Println(username)
    }
    db.Close()
}


func delete(){
    db, err := sql.Open("mysql", "root:@/test?charset=utf-8")
    checkErr(err)
    stmt, err = db.Prepare("delete from userinfo where uid=?")
    checkErr(err)
    res, err = stmt.Exec(id) 
    affect, err = res.RowsAffected()
    checkErr(err)
    fmt.Println(affect)
    db.Close()
}
