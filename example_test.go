// Copyright (c) 2015-2020, go_eddystone authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone_test

import (
	"fmt"

	eddystone "github.com/suapapa/go_eddystone"
)

func ExampleMakeUIDFrame() {
	f, _ := eddystone.MakeUIDFrame("0102030405060708090a", "0b0c0d0e0f10", -30)
	fmt.Println(f)
	fmt.Println([]byte(f))
	// Output: UID[Namespace:0x0102030405060708090a Instance:0x0b0c0d0e0f10 TxPwr:-30dBm]
	// [0 226 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 0 0]
}

func ExampleMakeURLFrame() {
	f, _ := eddystone.MakeURLFrame("http://github.com", -30)
	fmt.Println(f)
	fmt.Println([]byte(f))
	// Output: URL[Url:http://github.com TxPwr:-30dBm]
	// [16 226 2 103 105 116 104 117 98 7]
}

func ExampleMakeTLMFrame() {
	f, _ := eddystone.MakeTLMFrame(3300, 23.5, 1, 2)
	fmt.Println(f)
	fmt.Println([]byte(f))
	// Output: TLM[batt:3300 temp:23.500000, advCnt:1 secCnt:2]
	// [32 0 12 228 23 128 0 0 0 1 0 0 0 2]
}

func ExampleMakeEIDFrame() {
	eid, _ := eddystone.ComputingEIDValue([]byte("0123456789abcdef"), 0x12345678, 8)
	f, _ := eddystone.MakeEIDFrame(eid, -30)
	fmt.Println(f)
	fmt.Println([]byte(f))
	// EID[EphemetalIdentifier:0xb846453b3c9e7b59 TxPwr:-30dBm]
	// [48 226 184 70 69 59 60 158 123 89]
}
