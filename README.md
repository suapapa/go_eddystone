# eddystone : Golang package for support Eddystone beacon
[![GoDoc](https://godoc.org/github.com/suapapa/go_eddystone?status.svg)](https://godoc.org/github.com/suapapa/go_eddystone)
[![Build Status](https://travis-ci.org/suapapa/go_eddystone.png?branch=master)](https://travis-ci.org/suapapa/go_eddystone)

[Eddystone](https://github.com/google/eddystone) is an open beacon format from Google.

This package covers [Eddystone-UID][1], [Eddystone-URL][2], [Eddystone-TLM][3] from [Eddystone Protocol Specification][0].

    $ go get github.com/suapapa/go_eddystone


## Example
Checkout `example/beacon.go` for using this package with [gatt][4] to make a Eddystone beacon in Golang.

## Authors

* Homin Lee &lt;homin.lee@suapapa.net&gt;
* Luca Capra &lt;luca.capra@gmail.com&gt;

## Copyright & License

Copyright (c) 2015-2020, go_eddystone authors.
All rights reserved.
Use of this source code is governed by a BSD-style license that can be
found in the LICENSE file.

[0]: https://github.com/google/eddystone/blob/master/protocol-specification.md
[1]: https://github.com/google/eddystone/tree/master/eddystone-uid
[2]: https://github.com/google/eddystone/tree/master/eddystone-url
[3]: https://github.com/google/eddystone/tree/master/eddystone-tlm
[4]: https://github.com/paypal/gatt
