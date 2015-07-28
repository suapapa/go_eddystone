// Copyright 2015, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

import (
	"bytes"
	"testing"
)

func TestURLFrame(t *testing.T) {
	urlGoogle := "http://google.com/"
	f, err := NewURLFrame(urlGoogle, -20)
	if err != nil {
		panic(err)
	}

	expect := []byte{
		0x10, // FtURL
		intToByte(-20),
		0x02, // URL Scheme Prefix: http://
		0x07, // Length
		'g',  // 'g'
		'o',  // 'o'
		'o',  // 'o'
		'g',  // 'g'
		'l',  // 'l'
		'e',  // 'e'
		0x00, // Eddystone-URL HTTP URL encoding: .com/
	}

	if !bytes.Equal([]byte(f), expect) {
		t.Errorf("expect: %v, got:%v", []byte(f), expect)
	}
}
