package mysqlConnect

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"trojan-go-manage/util"
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
		dbuser: "admin",
		dbPWD:  "keoSt2DVSSIj0c0G6fog",
		dburl:  "yudou.ceef9olzngin.ap-northeast-2.rds.amazonaws.com",
		dbport: 3306,
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

// Getuserlist 查询用户列表
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
	var userlist []Userconfig
	for rows.Next() {
		err := rows.Scan(&id, &username, &passwd, &quota, &download, &upload)
		if err != nil {
			log.Fatal(err)
		}
		userlist = append(userlist, Userconfig{
			id:       id,
			Username: username,
			passwd:   passwd,
			Quota:    quota,
			Download: download,
			Upload:   upload,
		})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return userlist
}

// Insetuser 新增用户函数
func Insetuser() {
	insql := "INSERT INTO `trojan`.`users` (`username`, `password`,`quota`) VALUES (?,?,?)"
	value := [2]string{"zhao", "dsadas"}
	value[1] = util.GetSha224(value[1])

	quota := 0
	r, err := db.Exec(insql, value[0], value[1], quota)

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
