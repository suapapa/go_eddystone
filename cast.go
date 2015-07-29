// Copyright 2015, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

import (
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

func uint16ToBytes(a uint16) []byte {
	v := make([]byte, 2)
	v[0] = byte(a >> 8)
	v[1] = byte(a)
	return v
}

func uint32ToBytes(a uint32) []byte {
	v := make([]byte, 4)
	v[0] = byte(a >> 24)
	v[1] = byte(a >> 16)
	v[2] = byte(a >> 8)
	v[3] = byte(a)
	return v
}

func bytesToUint16(a []byte) (v uint16) {
	if len(a) != 2 {
		panic("invalid input")
	}

	v = uint16(a[0])<<8 | uint16(a[1])
	return
}

func bytesToUint32(a []byte) (v uint32) {
	if len(a) != 4 {
		panic("invalid input")
	}
	v = uint32(a[0])<<24 | uint32(a[1])<<16 | uint32(a[2])<<8 | uint32(a[3])
	return
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
