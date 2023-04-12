package main

import (
	"fmt"
	"trojan-go-manage/mysqlConnect"
)

func main() {
	fmt.Println("开始")
	mysqlConnect.MysqlOpen()
}
