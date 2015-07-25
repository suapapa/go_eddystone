// Copyright 2015, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

import (
	"errors"
	"strings"
)

var urlSchemePrefix = map[string]byte{
	"http://www.":  0x00,
	"https://www.": 0x01,
	"http://":      0x02,
	"https://":     0x03,
}

var urlEncoding = map[string]byte{
	".com/":  0x00,
	".org/":  0x01,
	".edu/":  0x02,
	".net/":  0x03,
	".info/": 0x04,
	".biz/":  0x05,
	".gov/":  0x06,
	".com":   0x07,
	".org":   0x08,
	".edu":   0x09,
	".net":   0x0a,
	".info":  0x0b,
	".biz":   0x0c,
	".gov":   0x0d,
}

func encodeURL(u string) (byte, []byte, error) {
	prefix := byte(0x02)
	for k, v := range urlSchemePrefix {
		if strings.HasPrefix(u, k) {
			prefix = v
			u = u[len(k):]
			break
		}
	}

	for k, v := range urlEncoding {
		u = strings.Replace(u, k, string(v), -1)
	}

	if len(u) > 17 {
		return 0x00, nil, errors.New("url too long")
	}

	return prefix, []byte(u), nil
}
