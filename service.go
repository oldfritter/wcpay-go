package wcpay

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	UnifiedOrerWCPayRequest = []string{"appid", "mch_id", "nonce_str", "sign", "body", "out_trade_no", "total_fee", "spbill_create_ip", "notify_url", "trade_type"}
	RefundWCPayRequest      = []string{"appid", "mch_id", "nonce_str", "sign", "transaction_id", "out_refund_no", "total_fee", "refund_fee"}
)

func UnifiedOrder(options map[string]string) (string, err error) {
	options["appid"] = Appid
	options["mch_id"] = MchId
	options["nonce_str"] = nonceStr()
	options["trade_type"] = "NATIVE"
	options["fee_type"] = "CNY"
	options["sign_type"] = "MD5"
	options["sign"] = generateSign(options)
	err := checkRequiredOptions(options, UnifiedOrerWCPayRequest)
	if err != nil {
		log.Println(err)
		return
	}
	ctx, cancelFun := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFun()
	req, _ := http.NewRequest(
		http.MethodPost,
		"https://api.mch.weixin.qq.com/pay/unifiedorder",
		bytes.NewReader([]byte(xmlBody(options))),
	)
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.Println(err)
		return "", err
	}
	var result UnifiedOrderResult
	err := xmlParse(&resp.Body, &result)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return result.CodeUrl, nil
}

func Refund(options map[string]string) (result RefundResult, err error) {
	options["appid"] = Appid
	options["mch_id"] = MchId
	options["nonce_str"] = nonceStr()
	options["sign_type"] = "MD5"
	options["sign"] = generateSign(options)
	err := checkRequiredOptions(options, RefundWCPayRequest)
	if err != nil {
		log.Println(err)
		return
	}
	ctx, cancelFun := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFun()
	req, _ := http.NewRequest(
		http.MethodPost,
		"https://api.mch.weixin.qq.com/secapi/pay/refund",
		bytes.NewReader([]byte(xmlBody(options))),
	)
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.Println(err)
		return result, err
	}
	err := xmlParse(&resp.Body, &result)
	return result, err
}

func checkRequiredOptions(options map[string]string, names []string) (err error) {
	for _, name := range names {
		var found bool
		for k, v := range options {
			if k == name {
				found = true
			}
		}
		if !found {
			err = fmt.Errorf("WCPay Warn: missing required option: %v", name)
			return
		}
	}
}
