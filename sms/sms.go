// Package sms Copyright 2016 The GiterLab Authors. All rights reserved.
package sms

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/GiterLab/urllib"
	"github.com/tobyzxj/uuid"
)

// HTTPDebugEnable http调试开关
var HTTPDebugEnable = false

// Param 短信发送所需要的参数
type Param struct {
	// 系统参数
	AccessKeyID      string // 阿里云颁发给用户的访问服务所用的密钥ID
	Timestamp        string // 格式为：yyyy-MM-dd’T’HH:mm:ss’Z’；时区为：GMT
	Format           string // 没传默认为JSON，可选填值：XML
	SignatureMethod  string // 建议固定值：HMAC-SHA1
	SignatureVersion string // 建议固定值：1.0
	SignatureNonce   string // 用于请求的防重放攻击，每次请求唯一
	Signature        string // 最终生成的签名结果值

	// 业务参数
	Action       string // API的命名，固定值，如发送短信API的值为：SendSms
	Version      string // API的版本，固定值，如短信API的值为：2017-05-25
	RegionID     string // API支持的RegionID，如短信API的值为：cn-hangzhou
	RecNum       string // 手机号
	SignName     string // 短信签名
	TemplateCode string // 短信模板
	ParamString  string // 短信版本中的参数
	OutID        string // 外部流水扩展字段
}

// SetAccessKeyID 设置密钥ID
func (p *Param) SetAccessKeyID(accesskeyid string) {
	p.AccessKeyID = accesskeyid
}

// GetAccessKeyID 获取密钥ID
func (p *Param) GetAccessKeyID() string {
	return p.AccessKeyID
}

// SetTimestamp 设置时间戳
func (p *Param) SetTimestamp(timestamp string) {
	p.Timestamp = timestamp
}

// GetTimestamp 获取时间戳
func (p *Param) GetTimestamp() string {
	return p.Timestamp
}

// SetFormat 设置返回格式，JSON/XML
func (p *Param) SetFormat(format string) {
	p.Format = format
}

// GetFormat 获取返回格式
func (p *Param) GetFormat() string {
	return p.Format
}

// SetSignatureMethod 设置签名方法
func (p *Param) SetSignatureMethod(signaturemethod string) {
	p.SignatureMethod = signaturemethod
}

// GetSignatureMethod 获取签名方法
func (p *Param) GetSignatureMethod() string {
	return p.SignatureMethod
}

// SetSignatureVersion 设置签名版本
func (p *Param) SetSignatureVersion(signatureversion string) {
	p.SignatureVersion = signatureversion
}

// GetSignatureVersion 获取签名版本
func (p *Param) GetSignatureVersion() string {
	return p.SignatureVersion
}

// SetSignatureNonce 设置每一次请求的唯一序列
func (p *Param) SetSignatureNonce(signaturenonce string) {
	p.SignatureNonce = signaturenonce
}

// GetSignatureNonce 获取当前请求的序列
func (p *Param) GetSignatureNonce() string {
	return p.SignatureNonce
}

// SetSignature 设置最终的签名结果
func (p *Param) SetSignature(signature string) {
	p.Signature = signature
}

// GetSignature 获取签名结果
func (p *Param) GetSignature() string {
	return p.Signature
}

// SetAction 设置API请求方法参数
func (p *Param) SetAction(action string) {
	p.Action = action
}

// GetAction 获取API请求方法参数
func (p *Param) GetAction() string {
	return p.Action
}

// SetVersion 设置API版本
func (p *Param) SetVersion(version string) {
	p.Version = version
}

// GetVersion 获取API版本
func (p *Param) GetVersion() string {
	return p.Version
}

// SetRegionID 设置API的RegionID
func (p *Param) SetRegionID(regioniD string) {
	p.RegionID = regioniD
}

// GetRegionID 获取API的RegionID
func (p *Param) GetRegionID() string {
	return p.RegionID
}

// SetRecNum 设置短信接收的手机号
func (p *Param) SetRecNum(RecNum string) {
	p.RecNum = RecNum
}

// GetRecNum 获取短信接收的手机号
func (p *Param) GetRecNum() string {
	return p.RecNum
}

// SetSignName 设置签名参数
func (p *Param) SetSignName(signname string) {
	p.SignName = signname
}

// GetSignName 获取签名参数
func (p *Param) GetSignName() string {
	return p.SignName
}

// SetTemplateCode 设置短信模板
func (p *Param) SetTemplateCode(templatecode string) {
	p.TemplateCode = templatecode
}

// GetTemplateCode 获取短信模板
func (p *Param) GetTemplateCode() string {
	return p.TemplateCode
}

// SetParamString 设置短信模板参数
func (p *Param) SetParamString(ParamString string) {
	p.ParamString = ParamString
}

// GetParamString 获取短信模板参数
func (p *Param) GetParamString() string {
	return p.ParamString
}

// SetOutID 设置外部流水扩展字段
func (p *Param) SetOutID(outid string) {
	p.OutID = outid
}

// GetOutID 获取外部流水扩展字段
func (p *Param) GetOutID() string {
	return p.OutID
}

// ErrorMessage 短信服务器返回的错误信息
type ErrorMessage struct {
	HTTPCode  int     `json:"-"`
	Model     *string `json:"Model,omitempty"`
	RequestID *string `json:"RequestId,omitempty"`
	Message   *string `json:"Message,omitempty"`
	Code      *string `json:"Code,omitempty"`
}

// GetHTTPCode 获取HTTP请求的错误码
func (e *ErrorMessage) GetHTTPCode() int {
	return e.HTTPCode
}

// SetHTTPCode 设置HTTP错误码
func (e *ErrorMessage) SetHTTPCode(code int) {
	e.HTTPCode = code
}

// GetModel get model
func (e *ErrorMessage) GetModel() string {
	if e != nil && e.Model != nil {
		return *e.Model
	}
	return ""
}

// GetRequestID 获取请求的ID序列
func (e *ErrorMessage) GetRequestID() string {
	if e != nil && e.RequestID != nil {
		return *e.RequestID
	}
	return ""
}

// GetMessage 获取错误信息
func (e *ErrorMessage) GetMessage() string {
	if e != nil && e.Message != nil {
		return *e.Message
	}
	return ""
}

// GetCode 获取请求的错误码
func (e *ErrorMessage) GetCode() string {
	if e != nil && e.Code != nil {
		return *e.Code
	}
	return ""
}

// Error 序列化成字符串
func (e *ErrorMessage) Error() string {
	body, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(body)
}

// Client HTTP请求配置信息
type Client struct {
	// SMS服务的地址，默认为（https://sms.aliyuncs.com）
	EndPoint string
	// 访问SMS服务的accessid，通过官方网站申请或通过管理员获取
	AccessID string
	// 访问SMS服务的accesskey，通过官方网站申请或通过管理员获取
	AccessKey string
	// 连接池中每个连接的Socket超时，单位为秒，可以为int或float。默认值为30
	SocketTimeout int

	// 其他参数
	Param Param
	param map[string]string
}

// SetEndPoint 设置短信服务器
func (c *Client) SetEndPoint(endPoint string) {
	c.EndPoint = endPoint
}

// SetAccessID 设置短信服务的accessid，通过官方网站申请或通过管理员获取
func (c *Client) SetAccessID(accessid string) {
	c.AccessID = accessid
}

// SetAccessKey 设置短信服务的accesskey，通过官方网站申请或通过管理员获取
func (c *Client) SetAccessKey(accesskey string) {
	c.AccessKey = accesskey
}

// SetSocketTimeout 设置短信服务的Socket超时，单位为秒，可以为int或float。默认值为30
func (c *Client) SetSocketTimeout(sockettimeout int) {
	if sockettimeout == 0 {
		sockettimeout = 30
	}
	c.SocketTimeout = sockettimeout
}

func (c *Client) calcStringToSign() string {
	c.param = make(map[string]string)
	c.param["SignatureMethod"] = c.Param.GetSignatureMethod()
	c.param["SignatureNonce"] = uuid.New()
	// sync c.Param.SignatureNonce
	c.Param.SetSignatureNonce(c.param["SignatureNonce"])
	c.param["AccessKeyId"] = c.Param.GetAccessKeyID()
	c.param["SignatureVersion"] = c.Param.GetSignatureVersion()
	c.param["Timestamp"] = time.Now().UTC().Format(time.RFC3339)
	// sync c.Param.Timestamp
	c.Param.SetTimestamp(c.param["Timestamp"])
	c.param["Format"] = c.Param.GetFormat()

	c.param["Action"] = c.Param.GetAction()
	c.param["Version"] = c.Param.GetVersion()
	c.param["RegionId"] = c.Param.GetRegionID()
	c.param["RecNum"] = c.Param.GetRecNum()
	c.param["SignName"] = c.Param.GetSignName()
	c.param["ParamString"] = c.Param.GetParamString()
	c.param["TemplateCode"] = c.Param.GetTemplateCode()

	strslice := make([]string, len(c.param))
	i := 0
	for k, v := range c.param {
		data := url.Values{}
		data.Add(k, v)
		strslice[i] = data.Encode()
		strslice[i] = percentEncodeBefore(strslice[i])
		i++
	}
	sort.Strings(strslice)
	return "POST&" + percentEncode("/") + "&" + percentEncode(strings.Join(strslice, "&"))
}

// SendOne 发送给一个手机号
func (c *Client) SendOne(RecNum, signname, templatecode, ParamString string) (e *ErrorMessage, err error) {
	var body []byte

	e = &ErrorMessage{}
	c.Param.SetSignName(signname)
	c.Param.SetTemplateCode(templatecode)
	c.Param.SetParamString(ParamString)
	c.Param.SetRecNum(RecNum)
	signature := signatureMethod(c.AccessKey, c.calcStringToSign())

	req := urllib.Post(c.EndPoint)
	if HTTPDebugEnable {
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
		reader, errGzip := gzip.NewReader(resp.Body)
		if errGzip != nil {
			return nil, errGzip
		}
		body, err = ioutil.ReadAll(reader)
	} else {
		body, err = ioutil.ReadAll(resp.Body)
	}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, e)
	e.SetHTTPCode(resp.StatusCode)
	if HTTPDebugEnable {
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

// SendMulti 发送给多个手机号, 最多100个
func (c *Client) SendMulti(RecNum []string, signname, templatecode, ParamString string) (e *ErrorMessage, err error) {
	var body []byte

	e = &ErrorMessage{}
	if len(RecNum) > 100 {
		return nil, errors.New("number of RecNum should be less than 100")
	}
	c.Param.SetSignName(signname)
	c.Param.SetTemplateCode(templatecode)
	c.Param.SetParamString(ParamString)
	c.Param.SetRecNum(strings.Join(RecNum, ","))
	signature := signatureMethod(c.AccessKey, c.calcStringToSign())

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
		reader, errGzip := gzip.NewReader(resp.Body)
		if errGzip != nil {
			return nil, errGzip
		}
		body, err = ioutil.ReadAll(reader)
	} else {
		body, err = ioutil.ReadAll(resp.Body)
	}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, e)
	e.SetHTTPCode(resp.StatusCode)
	if HTTPDebugEnable {
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

// New 创建一个短信发送客户端
func New(accessid, accesskey string) (c *Client) {
	c = new(Client)
	if c.EndPoint == "" {
		c.EndPoint = "https://sms.aliyuncs.com/"
	}
	c.AccessID = accessid
	c.AccessKey = accesskey
	c.Param.SetSignatureMethod("HMAC-SHA1")
	c.Param.SetSignatureNonce(uuid.New())
	c.Param.SetAccessKeyID(accessid)
	c.Param.SetSignatureVersion("1.0")
	c.Param.SetTimestamp(time.Now().UTC().Format(time.RFC3339))
	c.Param.SetFormat("JSON")

	c.Param.SetAction("SingleSendSms")
	c.Param.SetVersion("2016-09-27")
	c.Param.SetRegionID("cn-hangzhou")
	c.Param.SetRecNum("your_RecNum")
	c.Param.SetSignName("your_signname")
	c.Param.SetParamString("your_ParamString")
	c.Param.SetTemplateCode("your_templatecode")

	if urllib.GetDefaultSetting().Transport == nil {
		// set default setting for urllib
		trans := &http.Transport{
			MaxIdleConnsPerHost: 500,
			Dial: (&net.Dialer{
				Timeout: time.Duration(15) * time.Second,
			}).Dial,
		}

		urlSetting := urllib.HttpSettings{
			ShowDebug:        false,            // ShowDebug
			UserAgent:        "GiterLab",       // UserAgent
			ConnectTimeout:   15 * time.Second, // ConnectTimeout
			ReadWriteTimeout: 30 * time.Second, // ReadWriteTimeout
			TlsClientConfig:  nil,              // TlsClientConfig
			Proxy:            nil,              // Proxy
			Transport:        trans,            // Transport
			EnableCookie:     false,            // EnableCookie
			Gzip:             true,             // Gzip
			DumpBody:         true,             // DumpBody
		}
		if c.SocketTimeout != 0 {
			urlSetting.ConnectTimeout = time.Duration(c.SocketTimeout) * time.Second
			urlSetting.ReadWriteTimeout = time.Duration(c.SocketTimeout) * time.Second
		}
		if HTTPDebugEnable {
			urlSetting.ShowDebug = true
		} else {
			urlSetting.ShowDebug = false
		}
		urllib.SetDefaultSetting(urlSetting)
	}

	return c
}
