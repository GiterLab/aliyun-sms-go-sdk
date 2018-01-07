# aliyun-sms-go-sdk
Aliyun SMS SDK for golang

[![wercker status](https://app.wercker.com/status/5ef19ea6b2a854db200521592d0d7b2e/m/master "wercker status")](https://app.wercker.com/project/byKey/5ef19ea6b2a854db200521592d0d7b2e)

[![Build Status](https://travis-ci.org/GiterLab/aliyun-sms-go-sdk.svg?branch=master)](https://travis-ci.org/GiterLab/aliyun-sms-go-sdk)
[![GoDoc](https://godoc.org/github.com/GiterLab/aliyun-sms-go-sdk/sms?status.svg)](https://godoc.org/github.com/GiterLab/aliyun-sms-go-sdk/sms)

## About
短信服务（Short Message Service）是阿里云为用户提供的一种通信服务的能力，支持快速发送短信验证码、短信通知等。 完美支撑双11期间2亿用户，发送6亿短信，8万并发量。三网合一专属通道，与工信部携号转网平台实时互联。电信级运维保障，实时监控自动切换，到达率高达99%。

<font color=#FF0000 size=4>
注意: 
2017年12月20日至2018年1月21日 消息服务中的短信功能和云市场（阿里短信服务）将迁移至云通信短信服务
为了尽快使用更专业的服务，还请您确认迁移后尽快下载正确的SKD和API代码</font>
## Install

	$ go get -u -v github.com/GiterLab/aliyun-sms-go-sdk

	`github.com/GiterLab/aliyun-sms-go-sdk/sms` 将停止维护
	`github.com/GiterLab/aliyun-sms-go-sdk/dysms` 为迁移至云通信后的新SDK

## Usage

[使用帮助](https://github.com/GiterLab/aliyun-sms-go-sdk/blob/master/doc/tips.md)

**已过时示例**

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

**迁移后的例子：**

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
	
		// send to one person
		respSendSms, err := dysms.SendSms(uuid.New(), "1375821****", "多协云", "SMS_22175101", `{"company":"duoxieyun"}`).DoActionWithException()
		if err != nil {
			fmt.Println("send sms failed", err, respSendSms.Error())
			os.Exit(0)
		}
		fmt.Println("send sms succeed", respSendSms.GetRequestID())
	}




## Links 
- [Short Message Service，SMS(短信服务)](https://www.aliyun.com/product/sms)
- [HTTP协议及签名](https://help.aliyun.com/document_detail/56189.html?spm=5176.doc56189.6.576.JIUq2i)
- [API签名机制](https://help.aliyun.com/document_detail/56189.html?spm=5176.product44282.6.576.VQczaW)

## License

This project is under the Apache Licence, Version 2.0. See the [LICENSE](https://github.com/GiterLab/aliyun-sms-go-sdk/blob/master/LICENSE) file for the full license text.