package conn

import (
	"github.com/giskook/ring_client/protocol"
)

func event_handler_server_msg_common(conn *Conn) {
	for conn.read_more {
		cmdid, values := protocol.CheckProtocol(conn.recv_buffer)

		switch cmdid {
		case protocol.PROTOCOL_LOGRT:
			event_handler_server_msg_login(conn, values)
			conn.read_more = true
		case protocol.PROTOCOL_ILLEGAL:
			conn.read_more = false
		case protocol.PROTOCOL_HALF_PACK:
			conn.read_more = false
		}
	}
}
