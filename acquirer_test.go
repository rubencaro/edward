package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rubencaro/edward/lib/cnf"
	"github.com/rubencaro/edward/lib/tst"
)

func TestAcquiresOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		tst.Eq(t, "/some/path", req.URL.String())
		rw.Write([]byte(`OK`))
	}))
	defer server.Close()

	c := &cnf.Config{}
	c.HTTPClient = server.Client()
	c.ImgURL = server.URL + "/some/path"

	acq := &DefaultAcquirer{}
	body, err := acq.Acquire(c)

	tst.Ok(t, err)
	tst.Eq(t, []byte(`OK`), body)
}

func TestAcquiresNotOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		tst.Eq(t, "/some/path", req.URL.String())
		rw.WriteHeader(403)
	}))
	defer server.Close()

	c := &cnf.Config{}
	c.HTTPClient = server.Client()
	c.ImgURL = server.URL + "/some/path"

	acq := &DefaultAcquirer{}
	_, err := acq.Acquire(c)

	tst.Assert(t, err != nil, "There should be an error")
}
