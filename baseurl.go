package baseurl

import (
	"net/http"
	"strings"
)

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
