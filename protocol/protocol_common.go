package protocol

import (
	"bytes"
	"strings"
)

const (
	PROTOCOL_START_FLAG string = "$"
	PROTOCOL_END_FLAG   string = "\r\n"
	PROTOCOL_SEP        string = ":"
	PROTOCOL_MIN_LENGTH int    = 26

	PROTOCOL_LOGRT int = 1

	PROTOCOL_ILLEGAL   int = 254
	PROTOCOL_HALF_PACK int = 255

	PROTOCOL_REQ_LOGIN string = "TLOGIN"
)

var PROTOCOL = map[string]int{
	"PLOGRT": PROTOCOL_LOGRT,
}

func write_header(protocol_id string, imei string) string {
	cmd := PROTOCOL_START_FLAG + PROTOCOL_SEP
	cmd += protocol_id
	cmd += PROTOCOL_SEP
	cmd += imei
	cmd += PROTOCOL_SEP

	return cmd
}

func CheckProtocol(buffer *bytes.Buffer) (int, []string) {
	cmd := PROTOCOL_ILLEGAL
	var values []string
	bufferlen := buffer.Len()
	if bufferlen == 0 {
		return PROTOCOL_ILLEGAL, nil
	}
	p := string(buffer.Bytes())
	if string(p[0]) != PROTOCOL_START_FLAG {
		buffer.ReadByte()
		cmd, values = CheckProtocol(buffer)
	} else if bufferlen >= PROTOCOL_MIN_LENGTH {
		end_idx := strings.Index(p, PROTOCOL_END_FLAG)
		if end_idx == -1 {
			return PROTOCOL_HALF_PACK, nil
		} else {
			buf, _ := buffer.ReadString('\n')
			values = strings.Split(buf, PROTOCOL_SEP)
			return PROTOCOL[values[1]], values
		}
	} else {
		return PROTOCOL_HALF_PACK, nil
	}

	return cmd, values
}
