package conn

import (
	"log"
	"strings"
)

func event_handler_server_msg_login(c *Conn, values []string) {
	log.Println("event_handler_server_msg_login")
	log.Println(values)
	if strings.TrimSpace(values[4]) == "1" {
		c.status = ConnSuccess
		log.Println("success")
	}
}
