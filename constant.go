// Copyright 2015, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

// Eddystone Service UUID
const SvcUUID = 0xFEAA

// Eddystone Service UUID in Little Endian format
var SvcUUIDBytes = []byte{0xAA, 0xFE}

// FrameType for Eddystone frames
type frameType byte

func (ft frameType) String() string {
	switch ft {
	case ftUID:
		return "Eddystone-UID"
	case ftURL:
		return "Eddystone-URL"
	case ftTLM:
		return "Eddystone-TLM"
	}
	return "Invaild Frame"
}

// Eddystone frame types
const (
	ftUID frameType = 0x00
	ftURL           = 0x10
	ftTLM           = 0x20
)
