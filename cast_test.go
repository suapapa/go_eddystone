// Copyright 2015, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

import "testing"

var testData = map[float32]uint16{
	0.0:  0x0000,
	1.0:  0x0100,
	1.5:  0x0180,
	1.75: 0x01c0,
	// 1.00396: 0x0101, // FAIL on TestFixToFloat32. got:1.003906
	-1.0:  0xff00,
	-1.5:  0xfe80,
	-2:    0xfe00,
	-127:  0x8100,
	-0.5:  0xff80,
	-0.25: 0xffc0,
	0.5:   0x0080,
	-128:  0x8000,
	127:   0x7f00,
	2.25:  0x0240,
	-2.25: 0xfdc0,
}

func TestFloat32ToFix(t *testing.T) {
	for k, v := range testData {
		got := float32ToFix(k)
		if got != v {
			t.Errorf("%f wanted: 0x%04x, got:0x%04x", k, v, got)
		}
	}
}

func TestFixToFloat32(t *testing.T) {
	for k, v := range testData {
		got := fixTofloat32(v)
		if got != k {
			t.Errorf("0x%04x wanted: %f, got:%f", v, k, got)
		}
	}
}

func TestUintToBytes(t *testing.T) {
	v := uint16ToBytes(uint16(0x0102))
	if v[0] != 0x01 || v[1] != 0x02 {
		t.Errorf("failed to convert uint16 to byte")
	}

	v = uint32ToBytes(uint32(0x01020304))
	if v[0] != 0x01 || v[1] != 0x02 || v[2] != 0x03 || v[3] != 0x04 {
		t.Errorf("failed to convert uint16 to byte")
	}
}

func TestIntToByte(t *testing.T) {
	if intToByte(-20) != 0xec {
		t.Errorf("failed to convert int to byte")
	}

	if byteToInt(0xec) != -20 {
		t.Errorf("failed to convert byte to int")
	}
}
