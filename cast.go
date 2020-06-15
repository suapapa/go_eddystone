// Copyright 2015-2020, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
)

// fixed point representation : https://courses.cit.cornell.edu/ee476/Math/
func float32ToFix(a float32) uint16 {
	return uint16(a * 256)
}

func fixTofloat32(a uint16) float32 {
	if a&0x8000 == 0 {
		return float32(a) / 256.0
	}
	return -(float32(^a) + 1) / 256.0
}

func bytesToUint16(a []byte) (v uint16) {
	return binary.BigEndian.Uint16(a)
}

func bytesToUint32(a []byte) (v uint32) {
	return binary.BigEndian.Uint32(a)
}

func intToByte(a int) byte {
	return byte(a & 0xff)
}

func byteToInt(a byte) (v int) {
	v = int(a)
	if v&0x80 != 0 {
		v = -((^v + 1) & 0xff)
	}
	return
}

func hexStringToBytes(s string, size int) ([]byte, error) {
	r, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}

	if len(r) > size {
		return nil, errors.New("too long data")
	}

	if len(r) < size {
		return append(make([]byte, size-len(r), size), r...), nil
	}

	return r, nil
}
