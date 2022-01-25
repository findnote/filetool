package p3000

import (
	"bufio"
	log "github.com/pion/ion-log"
	"net"
	"regexp"
)

//
// Client
//  @Description: 与p90消息总线建立tcp长连接
//
type Client struct {
	conn           net.Conn
	onMessage      func(data []byte)
	onConnected    func()
	onDisconnected func(err error)
	establish      func() error
}

//
// NewClient
//  @Description: 新建一个tcp连接
//
func NewClient() (*Client, error) {
	log.Infof("NewClient(), Connecting to %s...\n", addr)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Errorf("Failed to connect to server[%v]", addr)
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

//
// OnMessage
//  @Description: 绑定消息处理函数
//  @receiver c
//
func (c *Client) OnMessage(handler func(data []byte)) {
	c.onMessage = handler
}

//
// OnConnected
//  @Description: 绑定心跳处理函数
//
func (c *Client) OnConnected(handler func()) {
	c.onConnected = handler
}

//
// OnDisconnected
//  @Description: 绑定连接重连函数
func (c *Client) OnDisconnected(handler func(err error)) {
	c.onDisconnected = handler
}

func (c *Client) OnEstablish(handler func() error) {
	c.establish = handler
}

//
// listen
//  @Description: 接收tcp数据和发送心跳
//  @receiver c
//
func (c *Client) listen() {
	//  建立业务连接
	err := c.establish()
	if err != nil {
		c.conn.Close()
		c.onDisconnected(err)
		return
	}

	topics["CYGBase:RTDB-CYGBase:modify"] = 1
	//  订阅设备变位
	for topic := range topics {
		err := Subscribe(topic, defaultMessageHandler)
		if err != nil {
			//  订阅失败，重新连接
			c.conn.Close()
			c.onDisconnected(err)
			return
		}
	}
	log.Infof("Subscribe success!!")

	//  开始心跳
	go c.onConnected()

	reader := bufio.NewReader(c.conn)
	buffer := make([]byte, 1024*10)
	for {
		var packet []byte
		n, err := reader.Read(buffer)
		if err != nil {
			c.conn.Close()
			log.Infof("ReadBytes err: %v", err)
			c.onDisconnected(err)
			return
		}

		log.Infof("read bytes....., n=%v", n)

		if n > 0 {
			packet = append(packet, buffer[:n]...)
			//  判断报文是否结束
			if isMatch, _ := regexp.Match(`\{.*}`, packet); isMatch {
				//  截取数据
				compile, _ := regexp.Compile(`\{.*}`)
				data := compile.Find(packet)

				go c.onMessage(data)
			}
		}

	}
}

// Send
//  @Description: 通过tcp连接发送消息
//  @param message 待发送的消息
//
func (c *Client) Send(message string) error {
	log.Infof("Send: message=%v", message)
	return c.SendBytes([]byte(message))
}

// Send bytes to client
func (c *Client) SendBytes(b []byte) error {
	// log.Infof("SendBytes: data=[% 02x]", b)

	_, err := c.conn.Write(b)
	if err != nil {
		c.conn.Close()
		log.Infof("c.conn.Write err: %v", err)
	}
	return err
}
