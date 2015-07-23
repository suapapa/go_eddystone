// Copyright 2015, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

import (
	"encoding/hex"
	"errors"
)

func hexDecode(s string, max int) ([]byte, error) {
	r, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}

	if len(s) > max {
		return nil, errors.New("too long data")
	}

	return r, nil
}
