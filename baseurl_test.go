package baseurl

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBaseURL(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, BaseUrl(r))
	}
	req, err := http.NewRequest("GET", "https://example.com/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler(w, req)
	str := w.Body.String()
	if str != `https://example.com` {
		t.Fatalf(`failed to get BaseUrl: expect: '%s', got: '%s'`, `https://example.com`, str)
	}
}
