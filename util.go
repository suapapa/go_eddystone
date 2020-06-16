// Copyright (c) 2015-2020, go_eddystone authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

import (
	"crypto/aes"
	"encoding/binary"
	"fmt"
)

// ComputeEIDValue returns 8 bytes EID value
// https://github.com/google/eddystone/blob/master/eddystone-eid/eid-computation.md
func ComputeEIDValue(identityKey []byte, ts uint32, k byte) (eid []byte, err error) {
	if len(identityKey) != 16 {
		return nil, fmt.Errorf("identityKey should be 16byte length for AES-128")
	}
	if 0 > k || k > 15 {
		return nil, fmt.Errorf("k should be between 0 and 15")
	}

	ts = ts & makeTsMask(k)

	tsBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(tsBytes, ts)

	// computing the temporary key
	tmpKeySrc := make([]byte, 11)             // padding
	tmpKeySrc = append(tmpKeySrc, 0xFF)       // salt
	tmpKeySrc = append(tmpKeySrc, 0x00, 0x00) // padding
	tmpKeySrc = append(tmpKeySrc, tsBytes[0:0+2]...)

	tmpBlk, err := aes.NewCipher(identityKey)
	if err != nil {
		return nil, err
	}

	tmpKey := make([]byte, 16)
	tmpBlk.Encrypt(tmpKey, tmpKeySrc)

	// computing the EID value
	eidSrc := make([]byte, 11) // padding
	eidSrc = append(eidSrc, k) // rotation period exponent
	eidSrc = append(eidSrc, tsBytes...)

	eidBlk, err := aes.NewCipher(tmpKey)
	if err != nil {
		return nil, err
	}
	eid = make([]byte, 16)
	eidBlk.Encrypt(eid, eidSrc)

	return eid[:8], nil
}

func makeTsMask(k byte) uint32 {
	var tsMask uint32
	for i := byte(0); i < k; i++ {
		tsMask <<= 1
		tsMask |= 1
	}
	return ^tsMask
}
