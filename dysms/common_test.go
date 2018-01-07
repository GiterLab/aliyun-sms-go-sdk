package dysms

import (
	"testing"
)

func Test_stringToSign(t *testing.T) {
	c := new(Client)
	c.EndPoint = "http://dysmsapi.aliyuncs.com/"
	c.AccessID = "testId"
	c.AccessKey = "testSecret"

	req := newRequset()
	// 1. 系统参数
	req.Put("SignatureMethod", "HMAC-SHA1")
	req.Put("SignatureNonce", "45e25e9b-0a6f-4070-8c85-2956eda1b466")
	req.Put("AccessKeyId", c.AccessID)
	req.Put("SignatureVersion", "1.0")
	req.Put("Timestamp", "2017-07-12T02:42:19Z")
	req.Put("Format", "XML")
	// 2. 业务API参数
	req.Put("Action", "SendSms")
	req.Put("Version", "2017-05-25")
	req.Put("RegionId", "cn-hangzhou")
	req.Put("PhoneNumbers", "15300000001")
	req.Put("SignName", "阿里云短信测试专用")
	req.Put("TemplateParam", "{\"customer\":\"test\"}")
	req.Put("TemplateCode", "SMS_71390007")
	req.Put("OutId", "123")
	stringToSign := req.CalcStringToSign("GET")
	stringToSignResult := `GET&%2F&AccessKeyId%3DtestId%26Action%3DSendSms%26Format%3DXML%26OutId%3D123%26PhoneNumbers%3D15300000001%26RegionId%3Dcn-hangzhou%26SignName%3D%25E9%2598%25BF%25E9%2587%258C%25E4%25BA%2591%25E7%259F%25AD%25E4%25BF%25A1%25E6%25B5%258B%25E8%25AF%2595%25E4%25B8%2593%25E7%2594%25A8%26SignatureMethod%3DHMAC-SHA1%26SignatureNonce%3D45e25e9b-0a6f-4070-8c85-2956eda1b466%26SignatureVersion%3D1.0%26TemplateCode%3DSMS_71390007%26TemplateParam%3D%257B%2522customer%2522%253A%2522test%2522%257D%26Timestamp%3D2017-07-12T02%253A42%253A19Z%26Version%3D2017-05-25`
	if stringToSign != stringToSignResult {
		t.Error("calcStringToSign failed")
	}
}
