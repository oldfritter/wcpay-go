package wcpay

import (
	"encoding/xml"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func nonceStr() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func xmlBody(options map[string]string) (xml string) {
	xml = "<xml>"
	for k, v := range options {
		xml += fmt.Sprintf("<%v>%v</%v>", k, v, k)
	}
	xml += "</xml>"
}

func xmlParse(xml *[]byte, options *interface{}) (err error) {
	err = xml.Unmarshal(*xml, &options)
}

func GenerateBatchNo() (batchNo string) {
	t := time.Now()
	batchNo = fmt.Sprintf("%v%v%v", t.Format("20060102150405"), t.UnixNano()%1000000000, rand.Intn(10))
}
