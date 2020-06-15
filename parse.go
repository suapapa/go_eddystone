package eddystone

import (
	"encoding/binary"
	"encoding/hex"
)

type EddystoneType string

const (
	EddystoneTypeUnknown EddystoneType = ""
	EddystoneTypeURL                   = "url"
	EddystoneTypeTLM                   = "tlm"
	EddystoneTypeUID                   = "uid"
)

func ParseType(frames []byte) EddystoneType {
	switch frameType(frames[0]) {
	case ftUID:
		return EddystoneTypeUID
	case ftTLM:
		return EddystoneTypeTLM
	case ftURL:
		return EddystoneTypeURL
	}
	return EddystoneTypeUnknown
}

func ParseUIDFrame(f []byte) (ns, instance string, txPower int) {
	return hex.EncodeToString(f[2 : 2+10]),
		hex.EncodeToString(f[12 : 12+6]),
		byteToInt(f[1])
}

func ParseURLFrame(f []byte) (url string, txPower int, err error) {
	txPower = byteToInt(f[1])
	url, err = decodeURL(f[2], f[3:])
	if err != nil {
		return url, txPower, err
	}
	return url, txPower, nil
}

func ParseTLMFrame(f []byte) (batt uint16, temp float32, advCnt uint32, secCnt uint32) {
	return binary.BigEndian.Uint16(f[2 : 2+2]),
		fixTofloat32(binary.BigEndian.Uint16(f[4 : 4+2])),
		binary.BigEndian.Uint32(f[6 : 6+4]),
		binary.BigEndian.Uint32(f[10 : 10+4])
}
