// Copyright 2016 The GiterLab Authors. All rights reserved.

package sms

import (
	"compress/gzip"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/GiterLab/urllib"
	"github.com/tobyzxj/uuid"
)

var HttpDebugEnable bool = false

type SmsParam struct {
	Action           string
	SignName         string
	TemplateCode     string
	RecNum           string
	ParamString      string
	Format           string
	Version          string
	AccessKeyId      string
	SignatureMethod  string
	Timestamp        string
	SignatureVersion string
	SignatureNonce   string
	RegionId         string
}

func (this *SmsParam) SetAction(action string) {
	this.Action = action
}

func (this *SmsParam) GetAction() string {
	return this.Action
}

func (this *SmsParam) SetSignName(signname string) {
	this.SignName = signname
}

func (this *SmsParam) GetSignName() string {
	return this.SignName
}

func (this *SmsParam) SetTemplateCode(templatecode string) {
	this.TemplateCode = templatecode
}

func (this *SmsParam) GetTemplateCode() string {
	return this.TemplateCode
}

func (this *SmsParam) SetRecNum(recnum string) {
	this.RecNum = recnum
}

func (this *SmsParam) GetRecNum() string {
	return this.RecNum
}

func (this *SmsParam) SetParamString(paramstring string) {
	this.ParamString = paramstring
}

func (this *SmsParam) GetParamString() string {
	return this.ParamString
}

func (this *SmsParam) SetFormat(format string) {
	this.Format = format
}

func (this *SmsParam) GetFormat() string {
	return this.Format
}

func (this *SmsParam) SetVersion(version string) {
	this.Version = version
}

func (this *SmsParam) GetVersion() string {
	return this.Version
}

func (this *SmsParam) SetAccessKeyId(accesskeyid string) {
	this.AccessKeyId = accesskeyid
}

func (this *SmsParam) GetAccessKeyId() string {
	return this.AccessKeyId
}

func (this *SmsParam) SetSignatureMethod(signaturemethod string) {
	this.SignatureMethod = signaturemethod
}

func (this *SmsParam) GetSignatureMethod() string {
	return this.SignatureMethod
}

func (this *SmsParam) SetTimestamp(timestamp string) {
	this.Timestamp = timestamp
}

func (this *SmsParam) GetTimestamp() string {
	return this.Timestamp
}

func (this *SmsParam) SetSignatureVersion(signatureversion string) {
	this.SignatureVersion = signatureversion
}

func (this *SmsParam) GetSignatureVersion() string {
	return this.SignatureVersion
}

func (this *SmsParam) SetSignatureNonce(signaturenonce string) {
	this.SignatureNonce = signaturenonce
}

func (this *SmsParam) GetSignatureNonce() string {
	return this.SignatureNonce
}

func (this *SmsParam) SetRegionId(regionid string) {
	this.RegionId = regionid
}

func (this *SmsParam) GetRegionId() string {
	return this.RegionId
}

// 短信服务器返回的错误信息
type ErrorMessage struct {
	HttpCode  int     `json"-"`
	Model     *string `json:"Model,omitempty"`
	RequestId *string `json:"RequestId,omitempty"`
	Message   *string `json:"Message,omitempty"`
	Code      *string `json:"Code,omitempty"`
}

func (e *ErrorMessage) GetHttpCode() int {
	return e.HttpCode
}

func (e *ErrorMessage) SetHttpCode(code int) {
	e.HttpCode = code
}

func (e *ErrorMessage) GetModel() string {
	if e != nil && e.Model != nil {
		return *e.Model
	}
	return ""
}

func (e *ErrorMessage) GetRequestId() string {
	if e != nil && e.RequestId != nil {
		return *e.RequestId
	}
	return ""
}

func (e *ErrorMessage) GetMessage() string {
	if e != nil && e.Message != nil {
		return *e.Message
	}
	return ""
}

func (e *ErrorMessage) GetCode() string {
	if e != nil && e.Code != nil {
		return *e.Code
	}
	return ""
}

func (e *ErrorMessage) Error() string {
	body, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(body)
}

type SMSClient struct {
	// SMS服务的地址，默认为（https://sms.aliyuncs.com）
	EndPoint string
	// 访问SMS服务的accessid，通过官方网站申请或通过管理员获取
	AccessId string
	// 访问SMS服务的accesskey，通过官方网站申请或通过管理员获取
	AccessKey string
	// 连接池中每个连接的Socket超时，单位为秒，可以为int或float。默认值为50
	SocketTimeout int

	// 其他参数
	Param SmsParam
	param map[string]string
}

// 设置短信服务器
func (c *SMSClient) SetEndPoint(end_point string) {
	c.EndPoint = end_point
}

// 设置短信服务的accessid，通过官方网站申请或通过管理员获取
func (c *SMSClient) SetAccessId(accessid string) {
	c.AccessId = accessid
}

// 设置短信服务的accesskey，通过官方网站申请或通过管理员获取
func (c *SMSClient) SetAccessKey(accesskey string) {
	c.AccessKey = accesskey
}

// 设置短信服务的Socket超时，单位为秒，可以为int或float。默认值为50
func (c *SMSClient) SetSocketTimeout(sockettimeout int) {
	if sockettimeout == 0 {
		sockettimeout = 50
	}
	c.SocketTimeout = sockettimeout
}

// 发送给多个手机号, 最多100个
func (c *SMSClient) SendMulti(recnum []string, signname, templatecode, paramstring string) (e *ErrorMessage, err error) {
	var body []byte

	e = &ErrorMessage{}
	if len(recnum) > 100 {
		return nil, errors.New("number of recnum should be less than 100")
	}
	c.Param.SetSignName(signname)
	c.Param.SetTemplateCode(templatecode)
	c.Param.SetParamString(paramstring)
	c.Param.SetRecNum(strings.Join(recnum, ","))
	signature := signature_method(c.AccessKey, c.calc_string_to_sign())

	req := urllib.Post(c.EndPoint)
	for k, v := range c.param {
		req.Param(k, v)
	}
	req.Param("Signature", signature)
	resp, err := req.Response()
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		body, err = ioutil.ReadAll(reader)
	} else {
		body, err = ioutil.ReadAll(resp.Body)
	}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, e)
	e.SetHttpCode(resp.StatusCode)
	if HttpDebugEnable {
		fmt.Println("C-->S:", req.DumpRequestString())
		fmt.Println("S-->C:", e.Error())
	}
	if err != nil {
		return e, err
	}
	if e.GetCode() != "" {
		return e, errors.New(e.GetCode())
	}
	return e, nil
}

// 发送给一个手机号
func (c *SMSClient) SendOne(recnum, signname, templatecode, paramstring string) (e *ErrorMessage, err error) {
	var body []byte

	e = &ErrorMessage{}
	c.Param.SetSignName(signname)
	c.Param.SetTemplateCode(templatecode)
	c.Param.SetParamString(paramstring)
	c.Param.SetRecNum(recnum)
	signature := signature_method(c.AccessKey, c.calc_string_to_sign())

	req := urllib.Post(c.EndPoint)
	if HttpDebugEnable {
		req.Debug(true)
	}
	for k, v := range c.param {
		req.Param(k, v)
	}
	req.Param("Signature", signature)
	resp, err := req.Response()
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		body, err = ioutil.ReadAll(reader)
	} else {
		body, err = ioutil.ReadAll(resp.Body)
	}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, e)
	e.SetHttpCode(resp.StatusCode)
	if HttpDebugEnable {
		fmt.Println("C-->S:", req.DumpRequestString())
		fmt.Println("S-->C:", e.Error())
	}
	if err != nil {
		return e, err
	}
	if e.GetCode() != "" {
		return e, errors.New(e.GetCode())
	}
	return e, nil
}

func (c *SMSClient) calc_string_to_sign() string {
	c.param = make(map[string]string)
	c.param["Action"] = c.Param.GetAction()
	c.param["SignName"] = c.Param.GetSignName()
	c.param["TemplateCode"] = c.Param.GetTemplateCode()
	c.param["RecNum"] = c.Param.GetRecNum()
	c.param["ParamString"] = c.Param.GetParamString()
	c.param["Format"] = c.Param.GetFormat()
	c.param["Version"] = c.Param.GetVersion()
	c.param["AccessKeyId"] = c.Param.GetAccessKeyId()
	c.param["SignatureMethod"] = c.Param.GetSignatureMethod()
	c.param["Timestamp"] = time.Now().UTC().Format(time.RFC3339)
	// sync c.Param.Timestamp
	c.Param.SetTimestamp(c.param["Timestamp"])
	c.param["SignatureVersion"] = c.Param.GetSignatureVersion()
	c.param["SignatureNonce"] = uuid.New()
	// sync c.Param.SignatureNonce
	c.Param.SetSignatureNonce(c.param["SignatureNonce"])
	c.param["RegionId"] = c.Param.GetRegionId()

	strslice := make([]string, len(c.param))
	i := 0
	for k, v := range c.param {
		data := url.Values{}
		data.Add(k, v)
		strslice[i] = data.Encode()
		strslice[i] = aliyun_sms_encode_over(strslice[i])
		i++
	}
	sort.Strings(strslice)
	return "POST&" + percent_encode("/") + "&" + percent_encode(strings.Join(strslice, "&"))
}

func signature_method(key, string_to_sign string) string {
	// The signature method is supposed to be HmacSHA1
	// A switch case is required if there is other methods available
	mac := hmac.New(sha1.New, []byte(key+"&"))
	mac.Write([]byte(string_to_sign))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// 一般支持 URL 编码的库（比如 Java 中的 java.net.URLEncoder）都是按照“application/x-www-form-urlencoded”的MIME类型的规则进行编码的。
// 实现时可以直接使用这类方式进行编码，
// 把编码后的字符串中加号（+）替换成%20、星号（*）替换成%2A、%7E 替换回波浪号（~）, 即可得到所需要的编码字符串
func percent_encode(s string) string {
	s = url.QueryEscape(s)
	s = strings.Replace(s, "+", "%20", -1)
	s = strings.Replace(s, "*", "%2A", -1)
	s = strings.Replace(s, "%7E", "~", -1)

	return s
}

// 把编码后的字符串中加号（+）替换成%20、星号（*）替换成%2A、%7E 替换回波浪号（~）, 即可得到所需要的编码字符串
func aliyun_sms_encode_over(s string) string {
	s = strings.Replace(s, "+", "%20", -1)
	s = strings.Replace(s, "*", "%2A", -1)
	s = strings.Replace(s, "%7E", "~", -1)

	return s
}

// 创建一个短信发送客户端
func New(accessid, accesskey string) (c *SMSClient) {
	c = new(SMSClient)
	if c.EndPoint == "" {
		c.EndPoint = "https://sms.aliyuncs.com/"
	}
	c.AccessId = accessid
	c.AccessKey = accesskey
	c.Param.SetAction("SingleSendSms")
	c.Param.SetSignName("your_signname")
	c.Param.SetTemplateCode("your_templatecode")
	c.Param.SetRecNum("your_recnum")
	c.Param.SetParamString("your_paramstring")
	c.Param.SetFormat("JSON")
	c.Param.SetVersion("2016-09-27")
	c.Param.SetAccessKeyId(accessid)
	c.Param.SetSignatureMethod("HMAC-SHA1")
	c.Param.SetTimestamp(time.Now().UTC().Format(time.RFC3339))
	c.Param.SetSignatureVersion("1.0")
	c.Param.SetSignatureNonce(uuid.New())
	c.Param.SetRegionId("cn-hangzhou")

	// set default setting for urllib
	url_setting := urllib.HttpSettings{
		false,            // ShowDebug
		"GiterLab",       // UserAgent
		50 * time.Second, // ConnectTimeout
		50 * time.Second, // ReadWriteTimeout
		nil,              // TlsClientConfig
		nil,              // Proxy
		nil,              // Transport
		false,            // EnableCookie
		true,             // Gzip
		true,             // DumpBody
	}
	if c.SocketTimeout != 0 {
		url_setting.ConnectTimeout = time.Duration(c.SocketTimeout) * time.Second
		url_setting.ReadWriteTimeout = time.Duration(c.SocketTimeout) * time.Second
	}
	if HttpDebugEnable {
		url_setting.ShowDebug = true
	} else {
		url_setting.ShowDebug = false
	}
	urllib.SetDefaultSetting(url_setting)

	return c
}
