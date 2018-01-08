// Package dysms Copyright 2016 The GiterLab Authors. All rights reserved.
package dysms

import (
	"encoding/json"
	"errors"
)

// SmsSendDetailDTO 短信发送记录信息
type SmsSendDetailDTO struct {
	PhoneNum     string `json:"PhoneNum"`     // 手机号码
	SendStatus   int    `json:"SendStatus"`   // 发送状态 1：等待回执，2：发送失败，3：发送成功
	ErrCode      string `json:"ErrCode"`      // 运营商短信错误码
	TemplateCode string `json:"TemplateCode"` // 模板ID
	Content      string `json:"Content"`      // 短信内容
	SendDate     string `json:"SendDate"`     // 发送时间
	ReceiveDate  string `json:"ReceiveDate"`  // 接收时间
	OutID        string `json:"OutId"`        // 外部流水扩展字段
}

// SmsSendDetailDTOs 短信发送记录查询列表
type SmsSendDetailDTOs struct {
	SmsSendDetailDTO []SmsSendDetailDTO `json:"SmsSendDetailDTO"`
}

// QuerySendDetailsResponse 短信发送记录查询接口服务器响应
type QuerySendDetailsResponse struct {
	ErrorMessage
	TotalCount        *int               `json:"TotalCount,omitempty"`        // 发送总条数
	TotalPage         *int               `json:"TotalPage,omitempty"`         // 总页数
	SmsSendDetailDTOs *SmsSendDetailDTOs `json:"SmsSendDetailDTOs,omitempty"` // 发送明细结构体
}

// GetTotalCount 发送总条数
func (q *QuerySendDetailsResponse) GetTotalCount() int {
	if q != nil && q.TotalCount != nil {
		return *q.TotalCount
	}
	return 0
}

// GetTotalPage 总页数
func (q *QuerySendDetailsResponse) GetTotalPage() int {
	if q != nil && q.TotalPage != nil {
		return *q.TotalPage
	}
	return 0
}

// GetSmsSendDetailDTOs 获取短信发送记录
func (q *QuerySendDetailsResponse) GetSmsSendDetailDTOs() *SmsSendDetailDTOs {
	if q != nil && q.SmsSendDetailDTOs != nil {
		return q.SmsSendDetailDTOs
	}
	return nil
}

// String 序列化成JSON字符串
func (q QuerySendDetailsResponse) String() string {
	body, err := json.Marshal(q)
	if err != nil {
		return ""
	}
	return string(body)
}

// QuerySendDetailsRequest 短信发送记录查询接口请求
type QuerySendDetailsRequest struct {
	Request *Request
}

// SetSendDate 设置短信发送日期 必须
// 短信发送日期格式yyyyMMdd,支持最近30天记录查询
func (q *QuerySendDetailsRequest) SetSendDate(sendDate string) {
	if q != nil && q.Request != nil {
		q.Request.Put("SendDate", sendDate)
	}
}

// GetSendDate 获取短信发送日期
func (q *QuerySendDetailsRequest) GetSendDate() string {
	if q != nil && q.Request != nil {
		return q.Request.Get("SendDate")
	}
	return ""
}

// SetPageSize 设置页大小 必须
// 页大小Max=50
func (q *QuerySendDetailsRequest) SetPageSize(pageSize string) {
	if q != nil && q.Request != nil {
		q.Request.Put("PageSize", pageSize)
	}
}

// GetPageSize 获取设置页大小
func (q *QuerySendDetailsRequest) GetPageSize() string {
	if q != nil && q.Request != nil {
		return q.Request.Get("PageSize")
	}
	return ""
}

// SetResourceOwnerID 来源于python，未知参数
func (q *QuerySendDetailsRequest) SetResourceOwnerID(resourceOwnerID string) {
	if q != nil && q.Request != nil {
		q.Request.Put("ResourceOwnerId", resourceOwnerID)
	}
}

// GetResourceOwnerID 来源于python，未知参数
func (q *QuerySendDetailsRequest) GetResourceOwnerID() string {
	if q != nil && q.Request != nil {
		return q.Request.Get("ResourceOwnerId")
	}
	return ""
}

// SetOwnerID 来源于python，未知参数
func (q *QuerySendDetailsRequest) SetOwnerID(ownerID string) {
	if q != nil && q.Request != nil {
		q.Request.Put("OwnerId", ownerID)
	}
}

// GetOwnerID 来源于python，未知参数
func (q *QuerySendDetailsRequest) GetOwnerID() string {
	if q != nil && q.Request != nil {
		return q.Request.Get("OwnerId")
	}
	return ""
}

// SetPhoneNumber 设置短信接收号码 必须
func (q *QuerySendDetailsRequest) SetPhoneNumber(phoneNumber string) {
	if q != nil && q.Request != nil {
		q.Request.Put("PhoneNumber", phoneNumber)
	}
}

// GetPhoneNumber 获取短信接收号码
func (q *QuerySendDetailsRequest) GetPhoneNumber() string {
	if q != nil && q.Request != nil {
		return q.Request.Get("PhoneNumber")
	}
	return ""
}

// SetCurrentPage 设置当前页码 必须
func (q *QuerySendDetailsRequest) SetCurrentPage(currentPage string) {
	if q != nil && q.Request != nil {
		q.Request.Put("CurrentPage", currentPage)
	}
}

// GetCurrentPage 获取当前页码
func (q *QuerySendDetailsRequest) GetCurrentPage() string {
	if q != nil && q.Request != nil {
		return q.Request.Get("CurrentPage")
	}
	return ""
}

// SetBizID 设置发送流水号 可选
// 从调用发送接口返回值中获取
func (q *QuerySendDetailsRequest) SetBizID(bizID string) {
	if q != nil && q.Request != nil {
		q.Request.Put("BizId", bizID)
	}
}

// GetBizID 获取发送流水号
func (q *QuerySendDetailsRequest) GetBizID() string {
	if q != nil && q.Request != nil {
		return q.Request.Get("BizId")
	}
	return ""
}

// SetResourceOwnerAccount 来源于python，未知参数
func (q *QuerySendDetailsRequest) SetResourceOwnerAccount(resourceOwnerAccount string) {
	if q != nil && q.Request != nil {
		q.Request.Put("ResourceOwnerAccount", resourceOwnerAccount)
	}
}

// GetResourceOwnerAccount 来源于python，未知参数
func (q *QuerySendDetailsRequest) GetResourceOwnerAccount() string {
	if q != nil && q.Request != nil {
		return q.Request.Get("ResourceOwnerAccount")
	}
	return ""
}

// DoActionWithException 发起HTTP请求
func (q *QuerySendDetailsRequest) DoActionWithException() (resp *QuerySendDetailsResponse, err error) {
	if q != nil && q.Request != nil {
		resp := &QuerySendDetailsResponse{}
		body, httpCode, err := q.Request.Do("QuerySendDetails")
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
	return nil, errors.New("QuerySendDetailsRequest is nil")
}

// QuerySendDetails 短信发送记录查询接口
// bizID 可选 - 流水号
// phoneNumber 查询的手机号码
// pageSize 必填 - 页大小
// currentPage 必填 - 当前页码从1开始计数
// sendDate 必填 - 发送日期 支持30天内记录查询，格式yyyyMMdd
func QuerySendDetails(bizID, phoneNumber, pageSize, currentPage, sendDate string) *QuerySendDetailsRequest {
	req := newRequset()
	req.Put("Version", "2017-05-25")
	req.Put("Action", "QuerySendDetails")

	r := &QuerySendDetailsRequest{Request: req}
	r.SetPhoneNumber(phoneNumber) // 查询的手机号码
	if bizID != "" {
		r.SetBizID(bizID) // 可选 - 流水号
	}
	r.SetSendDate(sendDate)       // 必填 - 发送日期 支持30天内记录查询，格式yyyyMMdd
	r.SetCurrentPage(currentPage) // 必填 - 当前页码从1开始计数
	r.SetPageSize(pageSize)       // 必填 - 页大小
	return r
}
