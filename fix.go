// Copyright 2015, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eddystone

// fixed point representation : https://courses.cit.cornell.edu/ee476/Math/

func float32ToFix(a float32) uint16 {
	return uint16(a * 256)
}

func fixTofloat32(a uint16) float32 {
	if a&0x8000 == 0 {
		return float32(a) / 256.0
	}
	return -(float32(^a) + 1) / 256.0
}
