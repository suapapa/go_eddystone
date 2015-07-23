// Copyright 2015, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package eddystone provides tools for making eddystone frame
package eddystone

// UIDFrame is Eddystone-UID https://github.com/google/eddystone/tree/master/eddystone-uid
type UIDFrame struct {
	namespace string
	id        string
}

// URLFrame is Eddystone-URL https://github.com/google/eddystone/tree/master/eddystone-url
type URLFrame struct {
	url string
}

// TLMFrame is https://github.com/google/eddystone/tree/master/eddystone-tlm
type TLMFrame struct {
	version uint8
	vbatt   int
	temp    int
	advCnt  int
	SecCnt  int
}
