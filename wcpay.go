package wcpay

import (
	"crypto/x509"
	"log"

	"golang.org/x/crypto/pkcs12"
)

var (
	Appid     string
	MchId     string
	AppSecret string
	Key       string

	ApiclientCert *x509.Certificate
	ApiclientKey  interface{}
)

type UnifiedOrderResult struct {
	XMLName    xml.Name `xml:"xml"`
	ReturnCode string   `xml:"return_code"`
	ReturnMsg  string   `xml:"return_msg"`
	Appid      string   `xml:"appid"`
	MchId      string   `xml:"mch_id"`
	NonceStr   string   `xml:"nonce_str"`
	Openid     string   `xml:"openid"`
	Sign       string   `xml:"sign"`
	TradeType  string   `xml:"trade_type"`
	PrepayId   string   `xml:"prepay_id"`
	ResultCode string   `xml:"result_code"`
	CodeUrl    string   `xml:"code_url"`
}
type RefundResult struct {
	XMLName       xml.Name `xml:"xml"`
	ReturnCode    string   `xml:"return_code"`
	ReturnMsg     string   `xml:"return_msg"`
	Appid         string   `xml:"appid"`
	MchId         string   `xml:"mch_id"`
	NonceStr      string   `xml:"nonce_str"`
	Sign          string   `xml:"sign"`
	ResultCode    string   `xml:"result_code"`
	TransactionId string   `xml:"transaction_id"`
	OutTradeNo    string   `xml:"out_trade_no"`
	OutRefundNo   string   `xml:"out_refund_no"`
	RefundId      string   `xml:"refund_id"`
	RefundFee     string   `xml:"refund_fee"`
}
type QueryOrderResult struct {
	XMLName       xml.Name `xml:"xml"`
	ReturnCode    string   `xml:"return_code"`
	ReturnMsg     string   `xml:"return_msg"`
	Appid         string   `xml:"appid"`
	MchId         string   `xml:"mch_id"`
	DeviceInfo    string   `xml:"device_info"`
	NonceStr      string   `xml:"nonce_str"`
	Sign          string   `xml:"sign"`
	ResultCode    string   `xml:"result_code"`
	Openid        string   `xml:"openid"`
	IsSubscribe   string   `xml:"is_subscribe"`
	TradeType     string   `xml:"trade_type"`
	BankType      string   `xml:"bank_type"`
	TotalFee      string   `xml:"total_fee"`
	FeeType       string   `xml:"fee_type"`
	TransactionId string   `xml:"transaction_id"`
	OutTradeNo    string   `xml:"out_trade_no"`
	Attach        string   `xml:"attach"`
	TimeEnd       string   `xml:"time_end"`
	TradeState    string   `xml:"trade_state"`
}
type CloseOrderResult struct {
	XMLName    xml.Name `xml:"xml"`
	ReturnCode string   `xml:"return_code"`
	ReturnMsg  string   `xml:"return_msg"`
	Appid      string   `xml:"appid"`
	MchId      string   `xml:"mch_id"`
	NonceStr   string   `xml:"nonce_str"`
	Sign       string   `xml:"sign"`
	ResultCode string   `xml:"result_code"`
	ResultMsg  string   `xml:"result_msg"`
}

func SetApiclientByPkcs12(pfxData *[]byte, pass string) {
	ApiclientKey, ApiclientCert, err := pkcs12.Decode(*pfxData, pass)
	if err != nil {
		log.Println(err)
	}
}
