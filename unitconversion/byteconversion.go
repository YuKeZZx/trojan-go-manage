package unitconversion

import (
	"fmt"
	"log"
)

// Byteconversion 字节转换
func Byteconversion(bt float64) (ctnum string) {
	if bt == 0 {
		log.Print("传入值为空")
		ctnum = "0Byte"
		return ctnum
	}
	units := [...]string{"KB", "MB", "GB"}
	//转换单位
	size := 1024.0
	for index := range units {
		if (bt / size) < 1 {
			ctnum = fmt.Sprintf("%.2f%s", bt, units[index])
		}
		bt = bt / size
		ctnum = fmt.Sprintf("%.2f%s", bt, units[index])
	}
	return ctnum
}
