package main

import (
	"database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
)


var (
    dbhostip = "ip:port"
    dbusername = "user"
    dbpassword = "passwd"
    dbname = "database"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func insertDb(recordDic map[string]string) bool {
    db, err := sql.Open("mysql", "user:passwd@tcp(ip:port)/database?charset=utf8")
    checkErr(err)
    stmt, err := db.Prepare("INSERT tbl_role SET name=?,create_time=?")
    checkErr(err)
    res, err := stmt.Exec(recordDic["name"], recordDic["create_time"])
    checkErr(err)
    id, err := res.LastInsertId()
    checkErr(err)
    fmt.Println(id)
    db.Close()
    return true
}

func updateDb(id int, name string) bool {
    db, err := sql.Open("mysql", "user:passwd@tcp(ip:port)/database?charset=utf8")
    checkErr(err)
    stmt, err := db.Prepare("update tbl_role set name=? where id=?")
    checkErr(err)
    res, err := stmt.Exec(name, id)
    checkErr(err)
    affect, err := res.RowsAffected()
    fmt.Println(affect)
    db.Close()
    return true
}

func selectDb() map[int]map[string]string{
    results := make(map[int]map[string]string)
    itemDic := make(map[string]string)
    db, err := sql.Open("mysql", "user:passwd@tcp(ip:port)/database?charset=utf8")
    checkErr(err)
    rows, err := db.Query("SELECT id, name FROM tbl_role")
    checkErr(err)
    i := 0
    for rows.Next() {
        var id int
        var name string
        err = rows.Scan(&id, &name)
        checkErr(err)
        itemDic["id"] = string(id)
        itemDic["name"] = name 
        results[i] = itemDic
        i = i + 1
    }
    db.Close()
    return results 
}

func deleteDb(id int) bool {
    db, err := sql.Open("mysql", "user:passwd@tcp(ip:port)/database?charset=utf8")
    checkErr(err)
    stmt, err := db.Prepare("delete from tbl_role where id=?")
    checkErr(err)
    res, err := stmt.Exec(id) 
    checkErr(err)
    affect, err := res.RowsAffected()
    checkErr(err)
    fmt.Println(affect)
    db.Close()
    return true
}

func main() {
    //insert
    insertDic := make(map[string]string)
    insertDic["name"] = "role3"
    insertDic["create_time"] = "1399064577"
    ret := insertDb(insertDic)
    if (! ret) {
        fmt.Println("insert error")
    }
    //update
    ret := updateDb(27, "role8")
    if (! ret) {
        fmt.Println("update error")
    }
    //select
    results := selectDb()
    fmt.Println(results)
    //delete
    ret := deleteDb(25)
    if (! ret) {
        fmt.Println("delete error")
    }
}
