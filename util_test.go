// Copyright (c) 2015-2020, go_eddystone authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

import (
	"bytes"
	"testing"
)

func TestConstructNSByTracatedHash(t *testing.T) {
	expect := []byte{228, 27, 42, 224, 106, 114, 155, 150, 228, 185}
	got := ConstructNSByTruncatedHash("kakaoenterprise.com")
	if !bytes.Equal(expect, got) {
		t.Errorf("expect %v got %v", expect, got)
	}
}

func TestConsturctNSByElidedUUID(t *testing.T) {
	a := ConstructNSByElidedUUID()
	b := ConstructNSByElidedUUID()
	if bytes.Equal(a, b) {
		t.Errorf("it sholudnt be same")
	}
}

func TestMakeTsMask(t *testing.T) {
	expect := uint32(0xFFFFFFFF)
	got := makeTsMask(0)
	if expect != got {
		t.Errorf("expect %v got %v", expect, got)
	}
	expect = uint32(0xFFFFFFF0)
	got = makeTsMask(4)
	if expect != got {
		t.Errorf("expect %v got %v", expect, got)
	}
	expect = uint32(0xFFFF8000)
	got = makeTsMask(15)
	if expect != got {
		t.Errorf("expect %v got %v", expect, got)
	}
}
