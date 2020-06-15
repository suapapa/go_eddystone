// Copyright (c) 2015-2020, go_eddystone authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

import (
	"encoding/binary"
	"encoding/hex"
)

// ParseHeader returns type of Eddystone frame
func ParseHeader(frames []byte) Header {
	return Header(frames[0])
}

// ParseUIDFrame returns contents of UID frame
func ParseUIDFrame(f []byte) (ns, instance string, txPower int) {
	return hex.EncodeToString(f[2 : 2+10]),
		hex.EncodeToString(f[12 : 12+6]),
		byteToInt(f[1])
}

// ParseURLFrame returns contents of URL frame
func ParseURLFrame(f []byte) (url string, txPower int, err error) {
	txPower = byteToInt(f[1])
	url, err = decodeURL(f[2], f[3:])
	if err != nil {
		return url, txPower, err
	}
	return url, txPower, nil
}

// ParseTLMFrame returns contents of TLM frame
func ParseTLMFrame(f []byte) (batt uint16, temp float32, advCnt uint32, secCnt uint32) {
	return binary.BigEndian.Uint16(f[2 : 2+2]),
		fixTofloat32(binary.BigEndian.Uint16(f[4 : 4+2])),
		binary.BigEndian.Uint32(f[6 : 6+4]),
		binary.BigEndian.Uint32(f[10 : 10+4])
}
