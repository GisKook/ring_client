package protocol

type LoginPacket struct {
	Imei       string
	DeviceType string
	Protocol   string
}

func (p *LoginPacket) Serialize() []byte {
	v := write_header(PROTOCOL_REQ_LOGIN, p.Imei)
	v += p.DeviceType
	v += PROTOCOL_SEP
	v += p.Protocol
	v += PROTOCOL_END_FLAG

	return []byte(v)
}
