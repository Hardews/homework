package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	Id int
	username string
	password string
}




func main() {
	var db,err = sql.Open("mysql","root:lmh123@tcp(127.0.0.1:3306)/user")
	if err!=nil{
		log.Fatal(err)
	}
	PrepareInsertDemo(db)
	PrepareQueryDemo(db)
}

func Check(db *sql.DB) {
	//检查数据库是否可用
	err := db.Ping()
	if err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println("successful!")
	}
}

func QueryRowDemo(db *sql.DB)  {
	/*
		单行查询db.QueryRow() 执行一次查询，并期望返回最多一行数据
		它总是返回非nil值，直到返回值被scan方法调用，才会返回被延迟错误。
		一定要有scan
	*/
	sqlStr := "select id,username from userData where id >= ?"
	var u User
	err := db.QueryRow(sqlStr,1).Scan(&u.Id,&u.username)
	if err != nil{
		fmt.Println("scan failed err :",err)
		return
	}
	fmt.Println(u)
}

func Query(db *sql.DB)  {

	sqlStr := "select id,username from userData where id >= ?"
	rows, err := db.Query(sqlStr,0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n",err)
		return
	}

	//非常重要的 关闭rows 释放持有的数据库链接
	defer rows.Close()

	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id,&u.username)
		if err != nil{
			fmt.Printf("query failed, err:%v\n",err)
			return
		}
		fmt.Printf("id : %d  name: %s \n",u.Id,u.username)
	}
}

func insertRowDemo(db *sql.DB)  {
	//插入数据
	sqlStr := "insert into userData(username) values (?)"
	ret, err := db.Exec(sqlStr,"lsy")
	if err != nil{
		fmt.Println("insert failed, err :",err)
		return
	}
	theID,err := ret.LastInsertId() //获取插入数据的ID
	if err!= nil{
		fmt.Println("get the ID failed",err)
		return
	}
	fmt.Println("insert success, the Id is",theID)
}

func updateRowDemo(db *sql.DB)  {
	//更新操作
	sqlStr := "update userData set username=? where id = ?"
	ret,err := db.Exec(sqlStr,"lsy",2)
	if err != nil{
		fmt.Println("update failed, err :",err)
		return
	}
	n,err := ret.RowsAffected() //操作影响的行数
	if err != nil{
		fmt.Println("get RowsAffected failed, err :",err)
		return
	}
	fmt.Println("update success, affected rows :",n)
}

func deleteRowDemo(db *sql.DB)  {
	// 删除操作
	 sqlStr := "delete from userData where id = ?"
	 ret, err := db.Exec(sqlStr,1)
	 if err!=nil{
		 fmt.Println("delete failed, err:",err)
		 return
	 }
	 n,err := ret.RowsAffected() //影响操作的行数
	 if err != nil{
		 fmt.Println("get error :",err)
		 return
	 }
	 fmt.Println("success, affected rows :",n)
}


/*
  预处理

  预处理查询
*/

func PrepareQueryDemo(db *sql.DB)  {
	sqlStr := "select id,username from userData where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed, err :",err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err!=nil{
		fmt.Println("query failed, err :",err)
		return
	}
	defer rows.Close()

	//循环读取结果集中的数据
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id,&u.username)
		if err != nil{
			fmt.Println("scan failed, err",err)
			return
		}
		fmt.Println(u.Id,u.username)
	}
}

//插入

func PrepareInsertDemo(db *sql.DB)  {
	sqlStr := "insert into userData (username) values (?)"
	stmt , err := db.Prepare(sqlStr)
	if err != nil{
		fmt.Println("prepare failed, err : ",err)
		return
	}
	defer stmt.Close()
	_ , err = stmt.Exec("Zl")
	if err != nil {
		fmt.Println("insert failed ,err :",err)
		return
	}
	_ , err = stmt.Exec("sb")
	if err != nil {
		fmt.Println("insert failed ,err :",err)
		return
	}
	fmt.Println("insert success")
}