// Copyright (c) 2015-2020, go_eddystone authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

// Frame represent Eddystone frame
type Frame []byte

// MakeUIDFrame makes Eddystone-UID frame
// https://github.com/google/eddystone/tree/master/eddystone-uid
func MakeUIDFrame(namespace, instance string, txPwr int) (Frame, error) {
	n, err := hexStringToBytes(namespace, 10)
	if err != nil {
		return nil, err
	}

	i, err := hexStringToBytes(instance, 6)
	if err != nil {
		return nil, err
	}

	f := make(Frame, 2, 20)
	f[0] = byte(UID)
	f[1] = intToByte(txPwr)

	f = append(f, n...)
	f = append(f, i...)
	f = append(f, 0x00, 0x00)

	return f, nil
}

// MakeURLFrame makes Eddystone-URL frame
// https://github.com/google/eddystone/tree/master/eddystone-url
func MakeURLFrame(url string, txPwr int) (Frame, error) {
	p, u, err := encodeURL(url)
	if err != nil {
		return nil, err
	}

	f := make(Frame, 3, 21)
	f[0] = byte(URL)
	f[1] = intToByte(txPwr)
	f[2] = p

	f = append(f, u...)

	return f, nil
}

// MakeTLMFrame makes Eddystone-TLM frame
// https://github.com/google/eddystone/tree/master/eddystone-tlm
func MakeTLMFrame(batt uint16, temp float32, advCnt, secCnt uint32) (Frame, error) {
	f := make(Frame, 14)
	f[0] = byte(TLM)
	f[1] = 0x00 // TLM version

	// TODO: check min mix for each items
	binary.BigEndian.PutUint16(f[2:2+2], batt)
	binary.BigEndian.PutUint16(f[4:4+2], float32ToFix(temp))
	binary.BigEndian.PutUint32(f[6:6+4], advCnt)
	binary.BigEndian.PutUint32(f[10:10+4], secCnt)

	return f, nil
}

// MakeEIDFrame makes Eddystone-EID frame
// https://github.com/google/eddystone/tree/master/eddystone-eid
func MakeEIDFrame(eid []byte, txPwr int) (Frame, error) {
	if len(eid) != 8 {
		return nil, ErrInvalidData
	}
	f := []byte{byte(EID)}
	f = append(f, intToByte(txPwr))
	return append(f, eid...), nil
}

func (f Frame) String() string {
	t := Header(f[0])

	switch t {
	case UID:
		return fmt.Sprintf("%s[Namespace:0x%s Instance:0x%s TxPwr:%ddBm]",
			t,
			hex.EncodeToString(f[2:2+10]),
			hex.EncodeToString(f[12:12+6]),
			byteToInt(f[1]),
		)
	case URL:
		url, err := decodeURL(f[2], f[3:])
		if err != nil {
			panic(err)
		}
		return fmt.Sprintf("%s[Url:%s TxPwr:%ddBm]",
			t,
			url,
			byteToInt(f[1]),
		)
	case TLM:
		return fmt.Sprintf("%s[batt:%d temp:%f, advCnt:%d secCnt:%d]",
			t,
			binary.BigEndian.Uint16(f[2:2+2]),
			fixTofloat32(binary.BigEndian.Uint16(f[4:4+2])),
			binary.BigEndian.Uint32(f[6:6+4]),
			binary.BigEndian.Uint32(f[10:10+4]),
		)
	case EID:
		return fmt.Sprintf("%s[EphemetalIdentifier:0x%s TxPwr:%ddBm]",
			t,
			hex.EncodeToString(f[2:2+8]),
			byteToInt(f[1]),
		)
	}

	return t.String()
}
