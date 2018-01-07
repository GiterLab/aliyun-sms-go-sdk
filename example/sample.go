package main

import (
	"fmt"
	"os"

	"github.com/GiterLab/aliyun-sms-go-sdk/sms"
)

// modify it to yours
const (
	ACCESSID  = "your_accessid"
	ACCESSKEY = "your_accesskey"
)

func main() {
	// 2017年12月20日至2018年1月21日 消息服务中的短信功能和云市场（阿里短信服务）将迁移至云通信短信服务
	// 为了尽快使用更专业的服务，还请您确认迁移后尽快下载正确的SKD和API代码
	// 此测试接口过时，请勿再使用
	sms.HTTPDebugEnable = true
	c := sms.New(ACCESSID, ACCESSKEY)
	// send to one person
	e, err := c.SendOne("1375821****", "多协云", "SMS_22175101", `{"company":"duoxieyun"}`)
	if err != nil {
		fmt.Println("send sms failed", err, e.Error())
		os.Exit(0)
	}
	// send to more than one person
	e, err = c.SendMulti([]string{"1375821****", "1835718****"}, "多协云", "SMS_22175101", `{"company":"duoxieyun"}`)
	if err != nil {
		fmt.Println("send sms failed", err, e.Error())
		os.Exit(0)
	}
	fmt.Println("send sms succeed", e.GetRequestID())
}
