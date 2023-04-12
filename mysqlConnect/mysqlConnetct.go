package mysqlConnect

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

// 连接mysql使用

type mysqlconfig struct {
	dbuser string
	dbPWD  string
	dburl  string
	dbport int
	dbname string
}

// 数据库指针
var db *sqlx.DB

// 数据库初始化
func init() {
	m := mysqlconfig{
		dbuser: "root",
		dbPWD:  "trojan",
		dburl:  "ydzx.club",
		dbport: 1035,
		dbname: "trojan",
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.dbuser, m.dbPWD, m.dburl, m.dbport, m.dbname)
	datebase, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	db = datebase
}

// Userconfig 用户表
type Userconfig struct {
	id       int
	Username string
	passwd   string
	Quota    int64
	Download int64
	Upload   int64
}

// 查询用户列表
func Getuserlist() []Userconfig {
	// 表字段
	// users表字段
	var (
		id       int
		username string
		passwd   string
		quota    int64
		download int64
		upload   int64
	)
	rows, err := db.Query("select id,username,password,quota,download,upload from users")
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		err := rows.Scan(&id, &username, &passwd, &quota, &download, &upload)
		if err != nil {
			log.Fatal(err)
		}
	}
	var userlist []Userconfig
	userlist = append(userlist, Userconfig{
		id:       id,
		Username: username,
		passwd:   passwd,
		Quota:    quota,
		Download: download,
		Upload:   upload,
	})
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return userlist
}

func Insetuser() {
	sql := "INSERT INTO `trojan`.`users` (`username`, `password`,`quota`) VALUES (?,?,?)"
	value := [2]string{"zhao", "dsadas"}
	//value[1] = crypto.SHA224.String("sdad")
	quota := 0
	r, err := db.Exec(sql, value[0], value[1], quota)

	if err != nil {
		log.Fatal(err)
		return
	}

	id, err := r.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("增加成功,id为:", id)
}
