// Copyright (c) 2015-2020, go_eddystone authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/google/uuid"
)

func TestConstructNSByTracatedHashFQDN(t *testing.T) {
	expect := []byte{228, 27, 42, 224, 106, 114, 155, 150, 228, 185}
	got := ConstructNSByTruncatedHashFQDN("kakaoenterprise.com")
	if !bytes.Equal(expect, got) {
		t.Errorf("expect %v got %v", expect, got)
	}
}

func TestConsturctNSByElidedUUID(t *testing.T) {
	expect, _ := hex.DecodeString("8b0ca750095477cb3e77")
	id, _ := uuid.Parse("8b0ca750-e7a7-4e14-bd99-095477cb3e77")
	got := ConstructNSByElidedUUID(id)
	if !bytes.Equal(expect, got) {
		t.Errorf("expect %v got %v", expect, got)
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
