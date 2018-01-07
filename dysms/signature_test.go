package dysms

import (
	"net/url"
	"testing"
)

func Test_signatureMethod(t *testing.T) {
	stringToSign := `GET&%2F&AccessKeyId%3DtestId%26Action%3DSendSms%26Format%3DXML%26OutId%3D123%26PhoneNumbers%3D15300000001%26RegionId%3Dcn-hangzhou%26SignName%3D%25E9%2598%25BF%25E9%2587%258C%25E4%25BA%2591%25E7%259F%25AD%25E4%25BF%25A1%25E6%25B5%258B%25E8%25AF%2595%25E4%25B8%2593%25E7%2594%25A8%26SignatureMethod%3DHMAC-SHA1%26SignatureNonce%3D45e25e9b-0a6f-4070-8c85-2956eda1b466%26SignatureVersion%3D1.0%26TemplateCode%3DSMS_71390007%26TemplateParam%3D%257B%2522customer%2522%253A%2522test%2522%257D%26Timestamp%3D2017-07-12T02%253A42%253A19Z%26Version%3D2017-05-25`
	signature := signatureMethod(`testSecret`, stringToSign)
	if url.QueryEscape(signature) != `zJDF%2BLrzhj%2FThnlvIToysFRq6t4%3D` {
		t.Error("signatureMethod failed")
	}
}
