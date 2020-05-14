package wcpay

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

func generateSign(params map[string]string) string {
	var keys []string
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	query := ""
	for i, key := range keys {
		if i > 0 {
			query += "&"
		}
		query += fmt.Sprintf("%v=%v", key, params[key])
	}
	query += "&key=" + Key
	return string(md5.Sum(query))
}

func verifySign(params map[string]string) bool {
	sign := params["sign"]
	return generateSign(params) == strings.ToUpper(sign)
}
