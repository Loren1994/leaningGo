package web

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mine_db?charset=utf8")
	checkErr(err)
	defer db.Close()
	//查询
	selectData(db)
	//更新
	updateData(db)
	//增加
	addData(db)
	//删除
	deleteData(db)
}

func deleteData(db *sql.DB) {
	stmt, err := db.Prepare("delete from user_tb where id=?")
	checkErr(err)
	res, err := stmt.Exec(24)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("删除行数:", affect)
}

func addData(db *sql.DB) {
	stmt, err := db.Prepare("INSERT user_tb set name=?,age=?,address=?")
	checkErr(err)
	res, err := stmt.Exec("goAdd", 16, "株洲路")
	checkErr(err)
	index, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("最新插入的id:", index)
}

func updateData(db *sql.DB) {
	stmt, err := db.Prepare("UPDATE user_tb set name=? where  id = ?")
	checkErr(err)
	res, err := stmt.Exec("go_name", 17)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("更新行数:", affect)
}

func selectData(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM user_tb")
	checkErr(err)
	for rows.Next() {
		var (
			id      int
			name    string
			age     int
			address string
		)
		err = rows.Scan(&id, &name, &age, &address)
		checkErr(err)
		fmt.Println(id, name, age, address)
	}
}
