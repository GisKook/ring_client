package conn

import (
	"bytes"
	"github.com/giskook/ring_client/base"
	"github.com/giskook/ring_client/conf"
	"github.com/giskook/ring_client/protocol"
	"log"
	"net"
	"time"
)

var ConnSuccess uint8 = 0
var ConnUnauth uint8 = 1

type Conn struct {
	conn        *net.TCPConn
	config      *conf.Configuration
	recv_buffer *bytes.Buffer
	ticker      *time.Ticker
	closeChan   chan struct{}
	terminal    *base.Terminal
	read_more   bool
	status      uint8
}

func NewConn(imei string, config *conf.Configuration) *Conn {
	return &Conn{
		recv_buffer: bytes.NewBuffer([]byte{}),
		config:      config,
		ticker:      time.NewTicker(time.Duration(config.Client.HeartInterval) * time.Second),
		closeChan:   make(chan struct{}),
		read_more:   true,
		terminal: &base.Terminal{
			Imei:       imei,
			DeviceType: "iii",
			Protocol:   "vvv",
		},
		status: ConnUnauth,
	}
}

func (c *Conn) Close() {
	c.ticker.Stop()
	c.recv_buffer.Reset()
	c.conn.Close()
	close(c.closeChan)
}

func (c *Conn) stable_connect() {
reconn:
	if c.conn != nil {
		c.conn.Close()
		c.recv_buffer.Reset()
	}
	tcpaddr, err := net.ResolveTCPAddr("tcp", c.config.Server.Addr)

	c.conn, err = net.DialTCP("tcp", nil, tcpaddr)
	if err != nil {
		log.Println(err.Error())
		goto reconn
	}
}

func (c *Conn) Do2() {
	for {
		select {
		case <-c.closeChan:
			return
		case <-c.ticker.C:
			if c.status == ConnSuccess {

			}
		}
	}
}

func (c *Conn) Do() {
	defer func() {
		go c.Do()
	}()
	c.stable_connect()
	login := &protocol.LoginPacket{
		Imei:       c.terminal.Imei,
		DeviceType: c.terminal.DeviceType,
		Protocol:   c.terminal.Protocol,
	}

	c.send(login.Serialize())

	for {
		buffer := make([]byte, 1024)
		buf_len, err := c.conn.Read(buffer)
		if err != nil {
			log.Println(err)
			return
		}
		c.recv_buffer.Write(buffer[0:buf_len])
		log.Printf("<IN> %x\n", buffer[0:buf_len])
		c.read_more = true
		event_handler_server_msg_common(c)
	}
}

func (c *Conn) send(data []byte) {
	log.Printf("<OUT> %x %s\n", c, string(data))
	c.conn.Write(data)
}
