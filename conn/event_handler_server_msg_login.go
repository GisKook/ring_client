package conn

import (
	"log"
)

func event_handler_server_msg_login(c *Conn, values []string) {
	log.Println("event_handler_server_msg_login")
	if values[4] == "1" {
		c.status = ConnSuccess
	}
}
