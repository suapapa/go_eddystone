// Copyright 2015, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

import (
	"encoding/hex"
	"fmt"
)

// Frame represent Eddystone frame
type Frame []byte

// NewUIDFrame makes Eddystone-UID frame
// https://github.com/google/eddystone/tree/master/eddystone-uid
func NewUIDFrame(namespace, instance string, txPwr int) (Frame, error) {
	n, err := hexStringToBytes(namespace, 10)
	if err != nil {
		return nil, err
	}

	i, err := hexStringToBytes(instance, 6)
	if err != nil {
		return nil, err
	}

	f := make(Frame, 2, 20)
	f[0] = byte(ftUID)
	f[1] = intToByte(txPwr)

	f = append(f, n...)
	f = append(f, i...)
	f = append(f, 0x00, 0x00)

	return f, nil
}

// NewURLFrame makes Eddystone-URL frame
// https://github.com/google/eddystone/tree/master/eddystone-url
func NewURLFrame(url string, txPwr int) (Frame, error) {
	p, u, err := encodeURL(url)
	if err != nil {
		return nil, err
	}

	f := make(Frame, 3, 21)
	f[0] = byte(ftURL)
	f[1] = intToByte(txPwr)
	f[2] = p

	f = append(f, u...)

	return f, nil
}

// NewTLMFrame makes Eddystone-TLM frame
// https://github.com/google/eddystone/tree/master/eddystone-tlm
func NewTLMFrame(batt uint16, temp float32, advCnt, secCnt uint32) (Frame, error) {
	f := make(Frame, 2, 14)
	f[0] = byte(ftTLM)
	f[1] = 0x00 // TLM version

	// TODO: check min mix for each items

	f = append(f, uint16ToBytes(batt)...)
	f = append(f, uint16ToBytes(float32ToFix(temp))...)
	f = append(f, uint32ToBytes(advCnt)...)
	f = append(f, uint32ToBytes(secCnt)...)

	return f, nil
}

func (f Frame) String() string {
	t := frameType(f[0])

	switch t {
	case ftUID:
		return fmt.Sprintf("%s[Namespace:0x%s Instance:0x%s TxPwr:%ddBm]",
			t,
			hex.EncodeToString(f[2:2+10]),
			hex.EncodeToString(f[12:12+6]),
			byteToInt(f[1]),
		)
	case ftURL:
		url, err := decodeURL(f[2], f[3:])
		if err != nil {
			url = "invaild url frame: " + err.Error()
		}

		return fmt.Sprintf("%s[Url:%s TxPwr:%ddBm]",
			t,
			url,
			byteToInt(f[1]),
		)
	case ftTLM:
		return fmt.Sprintf("%s[batt:%d temp:%f, advCnt:%d secCnt:%d]",
			t,
			bytesToUint16(f[2:2+2]),
			fixTofloat32(bytesToUint16(f[4:4+2])),
			bytesToUint32(f[7:7+4]),
			bytesToUint32(f[11:7+4]),
		)
	}

	return t.String()
}

// Parse convert []byte to eddystone.Frame
func Parse(s []byte) (Frame, error) {
	// TODO:
	return nil, errNotImplemented
}
