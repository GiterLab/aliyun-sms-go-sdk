package sms

import (
	"net/url"
	"testing"
)

func Test_signatureMethod(t *testing.T) {
	stringToSign := `POST&%2F&AccessKeyId%3Dtestid%26Action%3DSingleSendSms%26Format%3DXML%26ParamString%3D%257B%2522name%2522%253A%2522d%2522%252C%2522name1%2522%253A%2522d%2522%257D%26RecNum%3D13098765432%26RegionId%3Dcn-hangzhou%26SignName%3D%25E6%25A0%2587%25E7%25AD%25BE%25E6%25B5%258B%25E8%25AF%2595%26SignatureMethod%3DHMAC-SHA1%26SignatureNonce%3D9e030f6b-03a2-40f0-a6ba-157d44532fd0%26SignatureVersion%3D1.0%26TemplateCode%3DSMS_1650053%26Timestamp%3D2016-10-20T05%253A37%253A52Z%26Version%3D2016-09-27`
	signature := signatureMethod(`testsecret`, stringToSign)
	if url.QueryEscape(signature) != `ka8PDlV7S9sYqxEMRnmlBv%2FDoAE%3D` {
		t.Error("signatureMethod failed")
	}
}

func Test_stringToSign(t *testing.T) {
	c := new(Client)
	c.EndPoint = "https://sms.aliyuncs.com/"
	c.AccessID = "testid"
	c.AccessKey = "testsecret"
	c.Param.SetAction("SingleSendSms")
	c.Param.SetSignName("标签测试")
	c.Param.SetTemplateCode("SMS_1650053")
	c.Param.SetRecNum("13098765432")
	c.Param.SetParamString(`{"name":"d","name1":"d"}`)
	c.Param.SetFormat("XML")
	c.Param.SetVersion("2016-09-27")
	c.Param.SetAccessKeyID("testid")
	c.Param.SetSignatureMethod("HMAC-SHA1")
	c.Param.SetTimestamp("2016-10-20T05:37:52Z")
	c.Param.SetSignatureVersion("1.0")
	c.Param.SetSignatureNonce("9e030f6b-03a2-40f0-a6ba-157d44532fd0")
	c.Param.SetRegionID("cn-hangzhou")
	stringToSign := c.calcStringToSign()
	stringToSignResult := `POST&%2F&AccessKeyId%3Dtestid%26Action%3DSingleSendSms%26Format%3DXML%26ParamString%3D%257B%2522name%2522%253A%2522d%2522%252C%2522name1%2522%253A%2522d%2522%257D%26RecNum%3D13098765432%26RegionId%3Dcn-hangzhou%26SignName%3D%25E6%25A0%2587%25E7%25AD%25BE%25E6%25B5%258B%25E8%25AF%2595%26SignatureMethod%3DHMAC-SHA1%26SignatureNonce%3D` + url.QueryEscape(url.QueryEscape(c.Param.GetSignatureNonce())) + `%26SignatureVersion%3D1.0%26TemplateCode%3DSMS_1650053%26Timestamp%3D` + url.QueryEscape(url.QueryEscape(c.Param.GetTimestamp())) + `%26Version%3D2016-09-27`
	if stringToSign != stringToSignResult {
		t.Error("calcStringToSign failed")
	}
}
