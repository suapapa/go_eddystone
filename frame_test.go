// Copyright (c) 2015-2020, go_eddystone authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

import (
	"bytes"
	"testing"
)

func TestMakeUIDFrame(t *testing.T) {
	f, err := MakeUIDFrame("0102030405060708090a", "123456", -30)
	if err != nil {
		panic(err)
	}

	expect := []byte{
		0, // hdrUID
		intToByte(-30),
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a,
		0x00, 0x00, 0x00, 0x12, 0x34, 0x56,
		0, 0,
	}

	if !bytes.Equal([]byte(f), expect) {
		t.Errorf("expect: %v, got:%v", expect, []byte(f))
	}
}

func TestMakeURLFrame(t *testing.T) {
	urlGoogle := "http://google.com/"
	f, err := MakeURLFrame(urlGoogle, -20)
	if err != nil {
		panic(err)
	}

	expect := []byte{
		0x10, // hdrURL
		intToByte(-20),
		0x02, // URL Scheme Prefix: http://
		'g',  // 'g'
		'o',  // 'o'
		'o',  // 'o'
		'g',  // 'g'
		'l',  // 'l'
		'e',  // 'e'
		0x00, // Eddystone-URL HTTP URL encoding: .com/
	}

	if !bytes.Equal([]byte(f), expect) {
		t.Errorf("expect: %v, got:%v", expect, []byte(f))
	}
}

func TestMakeTLMFrame(t *testing.T) {
	f, err := MakeTLMFrame(50, 28, 100, 100)
	if err != nil {
		panic(err)
	}

	expect := []byte{
		0x20, // hdrTLM
		0x00,
		0x00,
		0x32,
		0x1C,
		0x00,
		0x00,
		0x00,
		0x00,
		0x64,
		0x00,
		0x00,
		0x00,
		0x64,
	}

	if !bytes.Equal([]byte(f), expect) {
		t.Errorf("expect: %v, got:%v", expect, []byte(f))
	}
}
