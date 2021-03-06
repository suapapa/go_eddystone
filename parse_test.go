// Copyright (c) 2015-2020, go_eddystone authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseHeader(t *testing.T) {
	frames := []byte{byte(UID)}
	hdr := ParseHeader(frames)
	if hdr != UID {
		t.Fatal("Expected type is UID")
	}

	frames = []byte{byte(TLM)}
	hdr = ParseHeader(frames)
	if hdr != TLM {
		t.Fatal("Expected type is TLM")
	}

	frames = []byte{byte(URL)}
	hdr = ParseHeader(frames)
	if hdr != URL {
		t.Fatal("Expected type is URL")
	}

}

func TestParseUIDFrame(t *testing.T) {
	namespace := "AAAAAAAAAABBBBBBBBBB"
	instance := "123456123456"
	txPwr := 99
	f, err := MakeUIDFrame(namespace, instance, txPwr)
	if err != nil {
		t.Fatal(err)
	}

	namespace1, instance1, txPower1 := ParseUIDFrame(f)

	if namespace != strings.ToUpper(namespace1) {
		t.Fatal("namespace mismatch", namespace, namespace1)
	}
	if instance != instance1 {
		t.Fatal("instance mismatch", instance, instance1)
	}
	if txPwr != txPower1 {
		t.Fatal("txPower mismatch", txPwr, txPower1)
	}

}

func TestParseTLMFrame(t *testing.T) {
	var batt uint16 = 10
	var temp float32 = 30.1
	var advCnt uint32 = 100
	var secCnt uint32 = 200
	f, err := MakeTLMFrame(batt, temp, advCnt, secCnt)
	if err != nil {
		t.Fatal(err)
	}
	batt1, temp1, advCnt1, secCnt1 := ParseTLMFrame(f)
	if batt != batt1 {
		t.Fatal("batt mismatch", batt, batt1)
	}

	tempr := fmt.Sprintf("%.2f", temp)
	tempr1 := fmt.Sprintf("%.2f", temp1)
	if tempr != tempr1 {
		t.Fatal("temp mismatch", temp, temp1)
	}
	if advCnt != advCnt1 {
		t.Fatal("advCnt mismatch", advCnt, advCnt1)
	}
	if secCnt != secCnt1 {
		t.Fatal("secCnt mismatch", secCnt, secCnt1)
	}

}

func TestParseURLFrame(t *testing.T) {
	url := "https://example.com"
	txPwr := 99

	f, err := MakeURLFrame(url, txPwr)
	if err != nil {
		t.Fatal(err)
	}

	url1, txPwr1, err := ParseURLFrame(f)
	if err != nil {
		t.Fatal(err)
	}

	if url != url1 {
		t.Fatal("url mismatch", url, url1)
	}

	if txPwr != txPwr1 {
		t.Fatal("txPwr mismatch", txPwr, txPwr1)
	}

}
