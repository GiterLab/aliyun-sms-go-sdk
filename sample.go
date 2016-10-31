package main

import (
	"fmt"
	"os"

	"github.com/GiterLab/aliyun-sms-go-sdk/sms"
)

// modify it to yours
const (
	ENDPOINT  = "https://sms.aliyuncs.com/"
	ACCESSID  = "your_accessid"
	ACCESSKEY = "your_accesskey"
)

func main() {
	sms.HttpDebugEnable = true
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
	fmt.Println("send sms succeed", e.GetRequestId())
}
