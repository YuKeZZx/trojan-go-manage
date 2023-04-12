package main

import (
	"fmt"
	"log"
	"trojan-go-manage/mysqlConnect"
	"trojan-go-manage/unitconversion"
)

func main() {
	fmt.Println("开始")
	//mysqlConnect.MysqlOpen()
	userlist := mysqlConnect.Getuserlist()
	var user mysqlConnect.Userconfig
	user = userlist[0]
	totol := fmt.Sprintf("用户:%s\t限制流量%s\t使用流量:%s", user.Username, unitconversion.Byteconversion(float64(user.Quota)),
		unitconversion.Byteconversion(float64(user.Upload+user.Download)))
	log.Print(totol)
	mysqlConnect.Insetuser()
	//num := 21474836480.0
	//log.Print("\t", cruutnum)
}
