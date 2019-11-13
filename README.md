TestEd
======

Simple testing utilities for Golang.

Status: In-Development

```
package mypackage

import "testing"
import "kilobit.ca/go/tested/assert"

// Sanity Check
//
func TestMyPackageTest(t *testing.T) {
	assert.Expect(t, true, true)
}

```

Wait... another testing library?

Yes!  Another testing library.  Intended to be small, simple and
expand to satisfy new requirements as they emerge.  Hopefully it, or
something in the [See Other](#see-other) section will be useful for
your project.

Features
--------

Right now, tested couldn't be simpler while being any use at all.

Current Features:

- A simple Assert library to reduce test writing effort.
- Implemented the Expect assertion.

Future Features:
- More assert functions.
- Customize and direct output.
- Whatever seems useful while writing tests.

See Other
---------

There are lot's of great testing libraries that may better suite your
needs.

- [Testify](https://github.com/stretchr/testify)
- [Awesome Go #Testing](https://github.com/avelino/awesome-go#testing)

Installation
------------

```
$ go get 'kilobit.ca/go/tested'
```

Building
--------

```
$ cd tested
$ go test -v
$ go build
```

Contribute
----------

Contributions and collaborations are welcome!

Please submit a pull request with any bug fixes or feature requests
that you have. All submissions imply consent to use / distribute under
the terms of the LICENSE.

Support
-------

Submit tickets through [github](https://github.com/kilobit/tested).

License
-------

See LICENSE.

--  
Created: Nov 13, 2019  
By: Christian Saunders <cps@kilobit.ca>  
Copyright 2019 Kilobit Labs Inc.  
  
  
