// Copyright (c) 2015-2020, go_eddystone authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

import (
	"fmt"
)

// SvcUUID is Eddystone service UUID
const SvcUUID = 0xFEAA

var (
	// SvcUUIDBytes is Eddystone service UUID in Little Endian format
	SvcUUIDBytes = []byte{0xAA, 0xFE}

	// ErrInvalidFrame can be returned from Make*Frame()
	ErrInvalidFrame = fmt.Errorf("invalid frame")
	// ErrInvalidData can be returend if given data for Make*Fame() is invalid
	ErrInvalidData = fmt.Errorf("invalid data")
)

// Header for Eddystone frames
type Header byte

func (hdr Header) String() string {
	switch hdr {
	case UID:
		return "UID"
	case URL:
		return "URL"
	case TLM:
		return "TLM"
	case EID:
		return "EID"
	}
	return "Invaild Frame"
}

// Eddystone frame types
const (
	// UID means UID frame
	UID Header = 0x00
	// URL means URL frame
	URL = 0x10
	// TLM means TLM frame
	TLM = 0x20
	// EID means EID frame
	EID = 0x30
	// Unknown means it may not Eddystone frame
	Unknown = 0xff
)
