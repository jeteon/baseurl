package baseurl

import (
	"code.google.com/p/go.net/context"
	"net/http"
	"strings"
)

func BaseUrl(r *http.Request) string {
	proto := r.URL.Scheme
	host := r.Host
	return strings.Join([]string{proto, `://`, host}, ``)
}

const contextKey = `http.request.baseurl`

func NewContext(parent context.Context, r *http.Request) context.Context {
	return context.WithValue(parent, contextKey, BaseUrl(r))
}

func FromContext(parent context.Context) (string, bool) {
	baseUrl, ok := parent.Value(contextKey).(string)
	return baseUrl, ok
}
