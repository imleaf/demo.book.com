package conf

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

/*
前面是含义，后面是 go 的表示值,多种表示,逗号","分割
年　 06,2006
月份 1,01,Jan,January
日　 2,02,_2
时　 3,03,15,PM,pm,AM,am
分　 4,04
秒　 5,05
周几 Mon,Monday
时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
时区字母缩写 MST
您看出规律了么！哦是的，你发现了，这里面没有一个是重复的，所有的值表示都唯一对应一个时间部分。
并且涵盖了很多格式组合。
*/
// 时间格式化字符串
const SysTimeform string = "2006-01-02 15:04:05"
const SysTimeformShort string = "2006-01-02"

// 中国时区
var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")

//自定义配置文件地址
const SysWebconfigPath string = "./conf/web.config"

//全局配置文件map
var SysConfMap map[string]string

//初始函数
func init() {
	ReLoad()
}
func ReLoad() {
	config := make(map[string]string)

	f, err := os.Open(SysWebconfigPath)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		//如果前两位是//，则视为注释
		if len(s) >= 2 && s[0:2] == "//" {
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	fmt.Println("读取配置文件end", SysWebconfigPath)
	SysConfMap = config
}
