// Package `baseurl` is a utility for get base url from *http.Request.
//
// This library is supported to both http and https,
// and this library is very simple and easy of use.
//
// For Example:
//
//   package main
//   import (
//       "net/http"
//       "io"
//
//       "github.com/chnlr/baseurl"
//   )
//
//   func handler(w http.ResponseWriter, r *http.Request) {
//       io.WriteString(baseurl.BaseUrl(r))
//   }
package baseurl
