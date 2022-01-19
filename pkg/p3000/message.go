package p3000

import (
	"fmt"
	log "github.com/pion/ion-log"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

// p3000消息推送
var (
	httpClient *http.Client
	once       sync.Once
	wg         sync.WaitGroup
)

func CreateHTTPClient() {
	//  指定本地端口
	localPort := net.TCPAddr{Port: 9999}
	// 使用单例创建client
	once.Do(func() {
		httpClient = &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   3 * time.Second,
					KeepAlive: 30 * time.Second, //  每30s发送一次心跳来保持长连接
					LocalAddr: &localPort,
				}).DialContext,
				MaxIdleConnsPerHost: 100, // 对每个host的最大连接数量(MaxIdleConnsPerHost<=MaxIdleConns)
				IdleConnTimeout:     0,   // 多长时间未使用自动关闭连接，0表示没有限制
			},
		}
	})
}

//func main() {
//	CreateHTTPClient()
//
//	// 测试/establish
//	err := Establish(httpClient)
//	if err != nil {
//		for err != nil {
//			err = Establish(httpClient)
//		}
//	}
//
//	//  开始心跳
//	wg.Add(1)
//	go StartHeartBeat()
//	wg.Wait()
//}

//
// Establish
//  @Description: 建立长连接
//
func Establish(client *http.Client) error {
	log.Infof("sending /establish message...")
	req, err := http.NewRequest(http.MethodPost, "http://10.7.3.197:30002/establish", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "Keep-Alive")
	if err != nil {
		log.Errorf("create /establish request failed! reason: %v", err)
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Errorf("send /establish request failed! reason: %v", err)
		return err
	}

	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))

	res.Body.Close()

	log.Infof("send /establish message success!")
	return nil
}

//
// HeartBeat
//  @Description: 发送心跳
//
func HeartBeat(client *http.Client) {
	log.Infof("sending heartbeat...")
	req, err := http.NewRequest(http.MethodPost, "http://10.7.3.197:30002/heartbeat", nil)
	req.Header.Set("Connection", "Keep-Alive")
	if err != nil {
		log.Errorf("create /heartbeat failed! reason: %v", err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		log.Errorf("send /heartbeat failed! reason: %v", err)
		return
	}

	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))

	res.Body.Close()
	log.Infof("send heartbeat success!")
}

//
// StartHeartBeat
//  @Description: 开始定时发送心跳
//
func StartHeartBeat() {
	log.Infof("start heartbeat!")
	defer wg.Done()

	//  定时发送心跳
	for {
		HeartBeat(httpClient)
		time.Sleep(time.Second * 3)
	}
}

//
// Push
//  @Description: 本地模拟消息中心下发推送消息
//
func Push(client *http.Client) error {
	req, err := http.NewRequest(http.MethodGet, "http://10.7.3.245:9999/push", nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Errorf("new request failed! reason: %v", err)
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		log.Errorf("request failed! reason: %v", err)
		return err
	}

	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))

	res.Body.Close()
	return nil
}

//
// Subscribe
//  @Description: 订阅主题
//  @param topic 要订阅的主题
//  @param sessId 长连接标识
//
func Subscribe(topic string, sessId string) error {
	url := fmt.Sprintf("http://10.7.3.197:30002/%s/%s", topic, sessId)
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

	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))

	resp.Body.Close()
	return nil
}
