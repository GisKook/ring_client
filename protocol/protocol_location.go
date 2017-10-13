package protocol

import (
	"time"
)

type LocationPacket struct {
	Imei       string
	Longtitude string
	Latitude   string
}

func (p *LocationPacket) Serialize() []byte {
	v := write_header(PROTOCOL_REQ_LOCATION, p.Imei)
	v += PROTOCOL_SEP
	v += "00"
	v += PROTOCOL_SEP
	v += "100"
	v += PROTOCOL_SEP
	v += time.Now().Format("060102-150405")
	v += PROTOCOL_SEP
	v += "0"
	v += PROTOCOL_SEP
	v += "0"
	v += PROTOCOL_SEP
	v += "128.1234,38.1234,10"
	v += PROTOCOL_END_FLAG

	return []byte(v)
}
