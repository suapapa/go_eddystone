// Copyright (c) 2015-2020, go_eddystone authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

import (
	"errors"
	"strings"
)

var urlSchemePrefix = []string{
	"http://www.",
	"https://www.",
	"http://",
	"https://",
}

var urlEncoding = []string{
	".com/",
	".org/",
	".edu/",
	".net/",
	".info/",
	".biz/",
	".gov/",
	".com",
	".org",
	".edu",
	".net",
	".info",
	".biz",
	".gov",
}

func encodeURL(u string) (byte, []byte, error) {
	prefix := byte(0x02) // http://
	for i, v := range urlSchemePrefix {
		if strings.HasPrefix(u, v) {
			prefix = byte(i)
			u = u[len(v):]
			break
		}
	}

	for i, v := range urlEncoding {
		u = strings.Replace(u, v, string(byte(i)), -1)
	}

	if len(u) > 17 {
		return 0x00, nil, errors.New("url too long")
	}

	return prefix, []byte(u), nil
}

func decodeURL(prefix byte, encodedURL []byte) (string, error) {
	if int(prefix) >= len(urlSchemePrefix) {
		return "", errors.New("invaild prefix")
	}

	s := urlSchemePrefix[prefix]

	for _, b := range encodedURL {
		switch {
		case 0x00 <= b && b <= 0x13:
			s += urlEncoding[b]
		case 0x0e <= b && b <= 0x20:
			fallthrough
		case 0x7f <= b && b <= 0xff:
			return "", errors.New("invalid byte")
		default:
			s += string(b)
		}
	}

	return s, nil
}
