// Copyright 2015, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

// Eddystone Service UUID
const SvcUUID = 0xFEAA

// Eddystone Service UUID in Little Endian format
var SvcUUIDBytes = []byte{0xAA, 0xFE}

// FrameType for Eddystone frames
type FrameType byte

func (ft FrameType) String() string {
	switch ft {
	case FtUID:
		return "Eddystone-UID"
	case FtURL:
		return "Eddystone-URL"
	case FtTLM:
		return "Eddystone-TLM"
	}
	return "Invaild Frame"
}

// Eddystone frame types
const (
	FtUID FrameType = 0x00
	FtURL           = 0x10
	FtTLM           = 0x20
)
