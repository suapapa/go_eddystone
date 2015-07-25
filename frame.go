// Copyright 2015, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

import (
	"encoding/hex"
	"fmt"
)

// Frame represent Eddystone frame
type Frame []byte

// NewUIDFrame makes Eddystone-UID frame
// https://github.com/google/eddystone/tree/master/eddystone-uid
func NewUIDFrame(namespace, instance []byte, txPwr int) (Frame, error) {
	f := make(Frame, 20)
	f[0] = byte(FtUID)
	f[1] = byte(txPwr & 0xFF)
	copy(f[2:], namespace[:10+1])
	copy(f[12:], instance[:6+1])
	return f, nil
}

// NewURLFrame makes Eddystone-URL frame
// https://github.com/google/eddystone/tree/master/eddystone-url
func NewURLFrame(url string, txPwr int) (Frame, error) {
	p, u, err := encodeURL(url)
	if err != nil {
		return nil, err
	}

	f := make(Frame, 4)
	f[0] = byte(FtURL)
	f[1] = byte(txPwr & 0xFF)
	f[2] = p
	f[3] = byte(len(u))

	f = append(f, u...)

	return f, nil
}

// NewTLMFrame makes Eddystone-TLM frame
// https://github.com/google/eddystone/tree/master/eddystone-tlm
func NewTLMFrame(url string) Frame {
	panic(errNotImplemented)
}

func (f Frame) String() string {
	t, p := FrameType(f[0]), int(f[1])

	if p&0x80 != 0 {
		p = ^p + 1
	}

	switch t {
	case FtUID:
		return fmt.Sprintf("%s[Namespace:0x%s Instance:0x%s TxPwr:%ddBm]",
			t,
			hex.EncodeToString(f[2:2+10+1]),
			hex.EncodeToString(f[12:12+6+1]),
			p,
		)
	case FtURL:
		url, err := decodeURL(f[2], f[4:f[3]+1])
		if err != nil {
			url = "invaild url frame: " + err.Error()
		}

		return fmt.Sprintf("%s[Url:0x%s TxPwr:%ddBm]",
			t,
			url,
			p,
		)
	case FtTLM:
		panic(errNotImplemented)
	}

	return t.String()
}
