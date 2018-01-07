// Package dysms Copyright 2016 The GiterLab Authors. All rights reserved.
package dysms

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"strings"
)

// 计算 HMAC 值。
//     按照 RFC2104 的定义，使用得到的签名字符串计算签名 HMAC 值。
//     注意：计算签名时使用的 Key 就是您持有的 Access Key Secret 并加上一个 “&” 字符（ASCII:38），使用的哈希算法是 SHA1。
// 计算签名值。
//     按照 Base64 编码规则 把步骤 3 中的 HMAC 值编码成字符串，即得到签名值（Signature）。
func signatureMethod(key, stringToSign string) string {
	// The signature method is supposed to be HmacSHA1
	// A switch case is required if there is other methods available
	mac := hmac.New(sha1.New, []byte(key+"&"))
	mac.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// 把编码后的字符串中加号（+）替换成%20、星号（*）替换成%2A、%7E 替换回波浪号（~）, 即可得到所需要的编码字符串
func percentEncodeBefore(s string) string {
	s = strings.Replace(s, "+", "%20", -1)
	s = strings.Replace(s, "*", "%2A", -1)
	s = strings.Replace(s, "%7E", "~", -1)

	return s
}

// 一般支持 URL 编码的库（比如 Java 中的 java.net.URLEncoder）都是按照“application/x-www-form-urlencoded”的MIME类型的规则进行编码的。
// 实现时可以直接使用这类方式进行编码，
func percentEncode(s string) string {
	s = url.QueryEscape(s)
	return percentEncodeBefore(s)
}
