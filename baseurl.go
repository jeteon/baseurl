package baseurl

import (
	"net/http"
	"strings"
)

// This func is returned base url from `net/http#Request`.
//
// If requested url is `http://localhost:8080/path/to/file`,
// this func is returned `http://localhost:8080`.
//
// And If `net/http#Request` object is defined `TLS` property,
// This func is returned `https://localhost:8080`.
func BaseUrl(r *http.Request) string {
	var proto string
	if r.TLS != nil {
		proto = `https`
	} else {
		proto = `http`
	}

	host := r.Host
	return strings.Join([]string{proto, `://`, host}, ``)
}
