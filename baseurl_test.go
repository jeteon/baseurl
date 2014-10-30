package baseurl

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var testHTTPHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, BaseUrl(r)) })

func TestHTTP(t *testing.T) {
	ts := httptest.NewServer(testHTTPHandler)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != ts.URL {
		t.Fatalf(`failed to get BaseUrl:
	expect:	%v
	got:	%v`, ts.URL, string(data))
	}
}

func TestHTTPS(t *testing.T) {
	ts := httptest.NewTLSServer(testHTTPHandler)
	defer ts.Close()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{Transport: tr}

	res, err := client.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != ts.URL {
		t.Fatalf(`failed to get BaseUrl:
	expect:	%v
	got:	%v`, ts.URL, string(data))
	}

	if !strings.HasPrefix(string(data), `https://`) {
		t.Fatal(`this url is not https://`)
	}
}
