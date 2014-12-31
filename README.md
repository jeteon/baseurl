baseurl
=======

[![Build Status](https://travis-ci.org/chnlr/baseurl.svg?branch=master)](https://travis-ci.org/chnlr/baseurl) [![GoDoc](http://godoc.org/github.com/chnlr/baseurl?status.svg)](http://godoc.org/github.com/chnlr/baseurl)


A golang utility for getting base url from `*http.Request`


Install
-------

```bash
$ go get github.com/chnlr/baseurl
```

Example Code
------------

```go
package main

import (
    "net/http"
    "io"

    "github.com/chnlr/baseurl"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // If this handler is listening on `http://localhost:8080/foo/bar`,
    // `baseurl.BaseUrl` is returned `http://localhost:8080`
    io.WriteString(baseurl.BaseUrl(r))
}
```

Documentation
-------------

Please See <http://godoc.org/github.com/chnlr/baseurl> about more information of this library.

Author
------

Naoki OKAMURA a.k.a nyarla <nyarla@thotep.net>

License
-------

MIT
