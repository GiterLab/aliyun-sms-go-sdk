开发阿里云短信服务注意事项
=======================

### 名词解释 ###

以下几个变量名称在调用 `SendOne` 和 `SendMulti` 函数时使用）

`recnum`: 用户接收的手机号

`signname`: 签名名称

`templatecode`: 模板CODE

`paramstring`: 参数字符串

### 准备工作 ###

[直达帮助页面](https://help.aliyun.com/document_detail/44346.html?spm=5176.doc44348.6.103.z0JAmF)

1. 新建短信签名

	用户需要先`新建短信签名`, 阿里云审核通过后会得到一个`签名名称`, 此`签名名称`即为`signname`

2. 新建模板

	- 用户需要先`新建模板`, 阿里云审核通过后会得到一个`模板CODE`, 此`模板CODE`即为`templatecode`

	- 用户在创建模板的时候，会在模板中添加变量（一个或多个），如下为一个例子：

			尊敬的用户，您的${device_id}(${devicename})设备已离线
		
		上面的`device_id`和`devicename`既是模板变量，用户在使用此SDK时，需要把这两个变量转换为参数字符串`paramstring`，此`paramstring`是一个json对象，如下所示：

			{"device_id":"T0000001","devicename":"测试设备"}
		
		以上字符串即为 `paramstring`.

		下面提供一个golang方式转换方法：
		
			A. 为每一个模板CODE建立一个模板参数结构体，实现一个String()方法

				type Alarm_Offline_SMS_22120102 struct {
					DeviceId   string `json:"device_id"`
					DeviceName string `json:"devicename"`
				}
				
				func (this Alarm_Offline_SMS_22120102) String() string {
					body, err := json.Marshal(this)
					if err != nil {
						return ""
					}
					return string(body)
				}

			B. 在设置paramstring参数时，直接使用String()方法产生：

				o := new(Alarm_Offline_SMS_22120102)
				o.DeviceId = "T0000001"
				o.DeviceName = "测试设备"
				paramstring ：= o.String()

				// 这里paramstring将会是符合要求的结果：{"device_id":"T0000001","devicename":"测试设备"}
				fmt.Println(paramstring)

			C. 例子sample.go，可以改写成如下形式：

				package main
				
				import (
					"encoding/json"
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
				
				type Register_SMS_22175101 struct {
					CompanyName string `json:"company"`
				}
				
				func (this *Register_SMS_22175101) String() string {
					body, err := json.Marshal(this)
					if err != nil {
						return ""
					}
					return string(body)
				}
				
				func main() {
					sms.HttpDebugEnable = true
					c := sms.New(ACCESSID, ACCESSKEY)
				
					// create a paramstring object
					o := new(Register_SMS_22175101)
					o.CompanyName = "duoxieyun"
				
					// send to one person
					e, err := c.SendOne("1375821****", "多协云", "SMS_22175101", o.String())
					if err != nil {
						fmt.Println("send sms failed", err, e.Error())
						os.Exit(0)
					}
					// send to more than one person
					e, err = c.SendMulti([]string{"1375821****", "1835718****"}, "多协云", "SMS_22175101", o.String())
					if err != nil {
						fmt.Println("send sms failed", err, e.Error())
						os.Exit(0)
					}
					fmt.Println("send sms succeed", e.GetRequestId())
				}

				


		

	

	