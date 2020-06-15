// Copyright (c) 2015-2020, go_eddystone authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

import "errors"

// SvcUUID is Eddystone service UUID
const SvcUUID = 0xFEAA

// SvcUUIDBytes is Eddystone service UUID in Little Endian format
var SvcUUIDBytes = []byte{0xAA, 0xFE}

// ErrInvalidFrame can be returned from Make*Frame()
var ErrInvalidFrame = errors.New("invalid frame")

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
