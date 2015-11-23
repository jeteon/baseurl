package baseurl

import (
	"net/http"
	"strings"
)

// This function returns the base url from `net/http#Request`.
//
// If the request URL is `http://localhost:8080/path/to/file`,
// this function returns `http://localhost:8080`.
//
// If the `net/http#Request` object defines a `TLS` property,
// this function returns `https://localhost:8080`.
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
