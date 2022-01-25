package p3000

import (
	"bufio"
	"github.com/go-co-op/gocron"
	log "github.com/pion/ion-log"
	"io"
	"mwp3000/api/config"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	isNeedReconnect = true
	SessId          string
	addr            string
	ss              = gocron.NewScheduler(time.UTC)
	conf            config.Config
)

func StartMessage(f config.Config) {
	conf = f
	log.Infof("Start Message!")

	isNeedReconnect = true
	for {
		if isNeedReconnect {
			log.Infof("\n--------------connecting to server---------------")
			log.Infof("startMessage() isNeedReconnect=%v", isNeedReconnect)

			connectServer()
		}

		time.Sleep(time.Second * 5)
	}
}

func connectServer() {
	//  通过配置文件获取目标主机的socket
	host := conf.Message.Host
	port := conf.Message.Port
	addr = host + ":" + port

	client, err := NewClient()
	if err != nil {
		if _, t := err.(*net.OpError); t {
			log.Errorf("Some problem connecting.")
		} else {
			log.Errorf("Unknown error: " + err.Error())
		}
	} else {
		isNeedReconnect = false

		//  注册消息处理函数
		client.OnMessage(func(data []byte) {
			//  暂时不处理收到的订阅消息，只打印
			log.Infof("%s", string(data))
		})

		client.OnEstablish(func() error {
			//  建立业务连接
			err, str := establish(client)
			if err != nil {
				return err
			}

			if len(str) > 0 {
				log.Infof("connection established! sessId: %s", str)
				SessId = str
			}
			return nil
		})

		//  注册心跳发送函数
		client.OnConnected(func() {
			//  开始发送心跳
			startHeartBeat(client)
		})

		//  注册重连函数
		client.OnDisconnected(func(err error) {
			isNeedReconnect = true
			log.Infof("Client disconnected. isNeedReconnect=%v, err=%v", isNeedReconnect, err)
		})

		go client.listen()
	}
}

//
// establish
//  @Description: 往tcp连接发送/establish的http报文
//
func establish(c *Client) (error, string) {
	//  构造http报文
	url := "http://" + addr + "/establish"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	req.Header.Add("Host", "localhost")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Connection", "Keep-Alive")
	req.Header.Add("Content-Length", "0")
	if err != nil {
		log.Errorf("Create /establish request failed! reason: %v", err)
		return err, ""
	}

	//  将http报文转成字节流
	byteReq, err := httputil.DumpRequest(req, false)
	if err != nil {
		log.Errorf("DumpRequest() Failed! Reason: %v", err)
		return err, ""
	}

	//log.Infof("%s", byteReq)
	err = c.SendBytes(byteReq)
	if err != nil {
		return err, ""
	}

	reader := bufio.NewReader(c.conn)
	buffer := make([]byte, 1024*10)
	var packet []byte
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			log.Infof("ReadBytes err: %v", err)
			return err, ""
		}

		if n > 0 {
			packet = append(packet, buffer[:n]...)
			//  判断报文是否结束
			if isMatch, _ := regexp.Match(`\{.*}`, packet); isMatch {
				return nil, getSessId(packet)
			}
		}
	}
}

//
// getSessId
//  @Description: 从响应报文中提取sessId
//  @param packet 响应报文
//  @return string sessId
//
func getSessId(packet []byte) string {
	message := string(packet)
	index := strings.Index(message, "key")
	startIndex := index + len(`key":"`)
	endIndex := len(message) - len(`"}`)
	return message[startIndex:endIndex]
}

//
// heartbeat
//  @Description: 往tcp连接发送/heartbeat的http报文
//
func heartbeat(c *Client) {
	//  构造http报文
	url := "http://" + addr + "/heartbeat"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	req.Header.Add("Host", "localhost")
	req.Header.Add("Connection", "Keep-Alive")
	req.Header.Add("Content-Length", "0")
	if err != nil {
		log.Errorf("Create /heartbeat request failed! reason: %v", err)
		return
	}

	//  将http报文转成字节流
	byteReq, err := httputil.DumpRequest(req, false)
	if err != nil {
		log.Errorf("DumpRequest() Failed! Reason: %v", err)
		return
	}

	//log.Infof("%s", byteReq)
	err = c.SendBytes(byteReq)
	if err != nil {
		//  对于心跳的错误不进行处理
		return
	}

	return
}

//
// startHeartBeat
//  @Description: 定时发送心跳
//
func startHeartBeat(c *Client) {
	//  心跳间隔通过配置文件配置
	ss.Every(30).Second().Do(func() {
		heartbeat(c)

		//  如果掉线重连，结束心跳线程
		if isNeedReconnect {
			ss.Stop()
			return
		}
	})

	ss.StartAsync()
}

//
// Subscribe
//  @Description: 订阅主题
//	@param topic 主题 CYGBase:RTDB-CYGBase:modify
//
func Subscribe(topic string) error {
	url := "http://" + addr + "/subscribe/" + topic + "/" + SessId
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		log.Errorf("create /subscribe request failed! reason: %v", err)
		return err
	}

	//  期望接收的body类型为任意类型
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Connection", "Keep-Alive")

	//  发送/subscribe报文
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("send /subscribe request failed! reson: %v", err)
		return err
	}

	io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()

	return nil
}

//
// Publish
//  @Description: 推送消息
//  @param data 消息体
//	@param topic 主题 CYGBase:RTDB-CYGBase:modify
//	报文格式：
//	{
//		"name": "",
//		"node": "",
//		"topic": "",
//		"sn": "",
//		"data": "[{\"key\":\"CYGDW:Hash:Device:1:161\",\"tvModifys\":[{\"t\":\"yx\",\"sv\":\"0\",\"dv\":\"1\",\"time\":\"\"}]}]"
//	}
//	转义字符不能省略
//
func Publish(data string, topic string) error {
	url := "http://" + addr + "/publish/" + topic + "/" + SessId

	var req *http.Request
	var err error
	if data != "" {
		req, err = http.NewRequest(http.MethodPost, url, strings.NewReader(data))
	} else {
		req, err = http.NewRequest(http.MethodPost, url, nil)
	}

	if err != nil {
		log.Errorf("create /subscribe request failed! reason: %v", err)
		return err
	}

	//  期望接收的body类型为任意类型
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Connection", "Keep-Alive")

	//  发送/subscribe报文
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("send /subscribe request failed! reson: %v", err)
		return err
	}

	io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()

	return nil
}
