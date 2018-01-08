package main

import (
	"fmt"
	"os"

	"github.com/GiterLab/aliyun-sms-go-sdk/dysms"
	"github.com/tobyzxj/uuid"
)

// modify it to yours
const (
	ACCESSID  = "your_accessid"
	ACCESSKEY = "your_accesskey"
)

func main() {
	dysms.HTTPDebugEnable = true
	dysms.SetACLClient(ACCESSID, ACCESSKEY) // dysms.New(ACCESSID, ACCESSKEY)

	// 短信发送
	respSendSms, err := dysms.SendSms(uuid.New(), "1375821****", "多协云", "SMS_22175101", `{"company":"duoxieyun"}`).DoActionWithException()
	if err != nil {
		fmt.Println("send sms failed", err, respSendSms.Error())
		os.Exit(0)
	}
	fmt.Println("send sms succeed", respSendSms.String())

	// 查询短信
	respQuerySendDetails, err := dysms.QuerySendDetails("612710515335092485^0", "1375821****", "10", "1", "20180107").DoActionWithException()
	if err != nil {
		fmt.Println("query sms failed", err, respQuerySendDetails.Error())
		os.Exit(0)
	}
	fmt.Println("query sms succeed", respQuerySendDetails.String())
}
