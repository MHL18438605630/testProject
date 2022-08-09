package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, _ := sql.Open("mysql", "root:12346@(127.0.0.1:3306)/goblog")
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Println("数据库连接失败")
		return
	}

	rows, _ := db.Query("select * from blog_category") //获取所有数据
	var cid int
	var name string
	var create_at string
	var update_at string
	for rows.Next() { //循环显示所有的数据
		rows.Scan(&cid, &name, &create_at, &update_at)
		fmt.Println(cid, "--", name)
	}

}
