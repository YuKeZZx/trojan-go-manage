package main

import (
	"fmt"
	"log"
	"trojan-go-manage/mysqlConnect"
	"trojan-go-manage/util"
)

func main() {
	fmt.Println("开始")
	mysqlConnect.Insetuser()
	//mysqlConnect.MysqlOpen()
	userlist := mysqlConnect.Getuserlist()
	var user mysqlConnect.Userconfig

	for i := 0; i < len(userlist); i++ {
		user = userlist[i]
		totol := fmt.Sprintf("用户:%s\t限制流量%s\t使用流量:%s", user.Username, util.Byteconversion(float64(user.Quota)),
			util.Byteconversion(float64(user.Upload+user.Download)))
		log.Print(totol)
	}

	//num := 21474836480.0
	//log.Print("\t", cruutnum)
	//mysqlConnect.GetSha224("yukezx0901")
}
