package mysqlConnect

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 连接mysql使用

type Mysqlconfig struct {
	dbuser string
	dbPWD  string
	dburl  string
	dbport int
	dbname string
}

// users表字段
var (
	id       int
	username string
	passwd   string
	quota    int64
	download int64
	upload   int64
)

func MysqlOpen() {
	m := Mysqlconfig{
		dbuser: "root",
		dbPWD:  "trojan",
		dburl:  "ydzx.club",
		dbport: 1035,
		dbname: "trojan",
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.dbuser, m.dbPWD, m.dburl, m.dbport, m.dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select id,username,password,quota,download,upload from users")
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	//println("内存", rows)
	for rows.Next() {
		err := rows.Scan(&id, &username, &passwd, &quota, &download, &upload)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nid:%d\n用户:%s\n密码:%s\n限制速度:%d\n总上传使用:%d\n总下载使用:%d", id, username, passwd, quota, upload, download)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
