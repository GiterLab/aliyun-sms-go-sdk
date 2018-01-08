// Package dysms Copyright 2016 The GiterLab Authors. All rights reserved.
package dysms

import (
	"encoding/json"
	"errors"
)

// SendSmsResponse 发送短信接口服务器响应
type SendSmsResponse struct {
	ErrorMessage
	BizID *string `json:"BizId,omitempty"` // 发送回执ID,可根据该ID查询具体的发送状态
}

// GetBizID 发送回执ID,可根据该ID查询具体的发送状态
func (s *SendSmsResponse) GetBizID() string {
	if s != nil && s.BizID != nil {
		return *s.BizID
	}
	return ""
}

// String 序列化成JSON字符串
func (s SendSmsResponse) String() string {
	body, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(body)
}

// SendSmsRequest 发送短信接口请求
type SendSmsRequest struct {
	Request *Request
}

// SetOutID 设置外部流水扩展字段
func (s *SendSmsRequest) SetOutID(outID string) {
	if s != nil && s.Request != nil {
		s.Request.Put("OutId", outID)
	}
}

// GetOutID 获取外部流水扩展字段
func (s *SendSmsRequest) GetOutID(outID string) string {
	if s != nil && s.Request != nil {
		return s.Request.Get("OutId")
	}
	return ""
}

// SetSignName 设置短信签名
func (s *SendSmsRequest) SetSignName(signName string) {
	if s != nil && s.Request != nil {
		s.Request.Put("SignName", signName)
	}
}

// GetSignName 获取短信签名
func (s *SendSmsRequest) GetSignName() string {
	if s != nil && s.Request != nil {
		return s.Request.Get("SignName")
	}
	return ""
}

// SetResourceOwnerID 来源于python，未知参数
func (s *SendSmsRequest) SetResourceOwnerID(resourceOwnerID string) {
	if s != nil && s.Request != nil {
		s.Request.Put("ResourceOwnerId", resourceOwnerID)
	}
}

// GetResourceOwnerID 来源于python，未知参数
func (s *SendSmsRequest) GetResourceOwnerID() string {
	if s != nil && s.Request != nil {
		return s.Request.Get("ResourceOwnerId")
	}
	return ""
}

// SetOwnerID 来源于python，未知参数
func (s *SendSmsRequest) SetOwnerID(ownerID string) {
	if s != nil && s.Request != nil {
		s.Request.Put("OwnerId", ownerID)
	}
}

// GetOwnerID 来源于python，未知参数
func (s *SendSmsRequest) GetOwnerID() string {
	if s != nil && s.Request != nil {
		return s.Request.Get("OwnerId")
	}
	return ""
}

// SetTemplateCode 短信模板ID
func (s *SendSmsRequest) SetTemplateCode(templateCode string) {
	if s != nil && s.Request != nil {
		s.Request.Put("TemplateCode", templateCode)
	}
}

// GetTemplateCode 获取短信模板ID
func (s *SendSmsRequest) GetTemplateCode() string {
	if s != nil && s.Request != nil {
		return s.Request.Get("TemplateCode")
	}
	return ""
}

// SetPhoneNumbers 短信接收号码。
// 支持以逗号分隔的形式进行批量调用，批量上限为1000个手机号码,批量调用相对于单条调用及时性稍有延迟,
// 验证码类型的短信推荐使用单条调用的方式
func (s *SendSmsRequest) SetPhoneNumbers(phoneNumbers string) {
	if s != nil && s.Request != nil {
		s.Request.Put("PhoneNumbers", phoneNumbers)
	}
}

// GetPhoneNumbers 获取短信接收号码。
func (s *SendSmsRequest) GetPhoneNumbers() string {
	if s != nil && s.Request != nil {
		return s.Request.Get("PhoneNumbers")
	}
	return ""
}

// SetResourceOwnerAccount 来源于python，未知参数
func (s *SendSmsRequest) SetResourceOwnerAccount(resourceOwnerAccount string) {
	if s != nil && s.Request != nil {
		s.Request.Put("ResourceOwnerAccount", resourceOwnerAccount)
	}
}

// GetResourceOwnerAccount 来源于python，未知参数
func (s *SendSmsRequest) GetResourceOwnerAccount() string {
	if s != nil && s.Request != nil {
		return s.Request.Get("ResourceOwnerAccount")
	}
	return ""
}

// SetTemplateParam 短信模板变量替换JSON串,
// 友情提示:如果JSON中需要带换行符,请参照标准的JSON协议对换行符的要求,
// 比如短信内容中包含\r\n的情况在JSON中需要表示成\r\n,否则会导致JSON在服务端解析失败
func (s *SendSmsRequest) SetTemplateParam(templateParam string) {
	if s != nil && s.Request != nil {
		s.Request.Put("TemplateParam", templateParam)
	}
}

// GetTemplateParam 获取短信模板变量替换JSON串,
func (s *SendSmsRequest) GetTemplateParam() string {
	if s != nil && s.Request != nil {
		return s.Request.Get("TemplateParam")
	}
	return ""
}

// DoActionWithException 发起HTTP请求
func (s *SendSmsRequest) DoActionWithException() (resp *SendSmsResponse, err error) {
	if s != nil && s.Request != nil {
		resp := &SendSmsResponse{}
		body, httpCode, err := s.Request.Do("SendSms")
		resp.SetHTTPCode(httpCode)
		if err != nil {
			return resp, err
		}
		err = json.Unmarshal(body, resp)
		if err != nil {
			return resp, err
		}
		if httpCode != 200 {
			return resp, errors.New(resp.GetCode())
		}
		return resp, nil
	}
	return nil, errors.New("SendSmsRequest is nil")
}

// SendSms 发送短信接口
// businessID 设置业务请求流水号，必填。
// phoneNumbers 短信发送的号码列表，必填。 多手机号使用,分割
// signName 短信签名
// templateCode 申请的短信模板编码,必填
// templateParam 短信模板变量参数
func SendSms(businessID, phoneNumbers, signName, templateCode, templateParam string) *SendSmsRequest {
	req := newRequset()
	req.Put("Version", "2017-05-25")
	req.Put("Action", "SendSms")

	r := &SendSmsRequest{Request: req}
	r.SetOutID(businessID)          // 设置业务请求流水号，必填
	r.SetSignName(signName)         // 短信签名
	r.SetPhoneNumbers(phoneNumbers) // 短信发送的号码列表，必填。
	r.SetTemplateCode(templateCode) // 短信模板
	if templateParam != "" {
		r.SetTemplateParam(templateParam) // 短信模板参数
	}
	return r
}
