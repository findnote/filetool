package p3000

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"gopkg.in/resty.v1"

	log "github.com/pion/ion-log"
	"github.com/tidwall/gjson"
)

// 创建http client
var client = resty.New()

func EnableDebug(enable bool) {
	client.SetDebug(enable)
}

func NewP3000Client() *P3000Conn {
	host := cfg.P3000.Host
	port := cfg.P3000.Port

	url := host + ":" + strconv.Itoa(port)

	baseUrl := "http://" + url

	// log.Infof("NewP3000Client conn, baseUrl=%v", baseUrl)

	return &P3000Conn{
		conn: baseUrl,
	}
}

func (c *P3000Conn) PostOrders(orjson OrdersPostReq) (string, int, error) {
	var orderUrl = c.conn + "/api/1.0/middle/dw/dlopt/storeDl"

	qjson, _ := json.MarshalIndent(orjson, "", " ")

	log.Infof(">>PostOrders url=%v, qjson=%v", orderUrl, string(qjson))
	log.Infof(">>PostOrders qjson: %v", qjson)

	resp, err := c.post(orderUrl, qjson)
	if err != nil {
		log.Infof("(c *P3000Conn) PostOrders Failed error=%v", err)
		return "", -1, err
	}

	jsonstr := resp.String()

	log.Infof(">>PostOrders respstr=%v", jsonstr)

	errcode := gjson.Get(jsonstr, "errNo").Int()
	respond := gjson.Get(jsonstr, "respond").String()

	// log.Infof(">>errcode: %v, respond: %v", errcode, respond)

	if errcode != 0 {
		return respond, int(errcode), err
	}

	return respond, 0, nil
}

// post请求封装
func (c *P3000Conn) post(url string, qjson interface{}) (r *resty.Response, e error) {
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(qjson).
		Post(url)

	statusCode := resp.StatusCode()
	if statusCode == 0 {
		return resp, ErrNotValidP3000
	}

	return resp, err
}

func (c *P3000Conn) dumpResp(r *resty.Response, e error) {
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", e)
	fmt.Println("  Status Code:", r.StatusCode())
	fmt.Println("  Status     :", r.Status())
	fmt.Println("  Time       :", r.Time())
	fmt.Println("  Received At:", r.ReceivedAt())
	fmt.Println("  Error      :", r.Error())
	fmt.Println("  Header    :", r.Header())
	fmt.Println("  IsSuccess  :", r.IsSuccess())
	fmt.Println()

	// pretty print json
	body := r.Body()

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, body, "", "\t")
	if error != nil {
		log.Infof("JSON parse error: %v", error)
		return
	}

	log.Infof("CSP Violation: %v\n", prettyJSON.String())
}
