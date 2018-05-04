package server

// you must cd [project dir]/bytemurmur.com/server
// and exec command : go-bindata  -o=components/asset.go -pkg=asset static/...

import (
	"testing"
	"time"
	"net/http"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
	"github.com/pkg/errors"
	"bytemurmur.com/server/components/router"
	"fmt"
)

func TestServer(t *testing.T) {
	c := make(chan bool, 1)
	NewHTTPServer()
	<-c
}

func TestNewHTTPServer(t *testing.T) {
	c := make(chan bool, 1)
	srv, err := NewHTTPServerWithoutServer(c)
	assert.NoError(t, err)
	assert.NotNil(t, srv)
	select {
	case b := <- c:
		assert.True(t, b)
	case <-time.After(1 * time.Second):
		panic(errors.New("Time out"))
	}

	resp, err := http.Get("http://localhost:4621/static/js/test.js")
	assert.NoError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, string(body), "var val = \"Hello World!\";")

}

type testHandler struct {}

func (d *testHandler)Get(ctx *router.Context) {
	w, _ := ctx.Response, ctx.Request
	param, exist := ctx.Params["id"]
	if exist {
		fmt.Fprintf(w, "Test Pass! " + param)
	} else {
		fmt.Fprintf(w, "Test Pass!")
	}
}

func TestHandler(t *testing.T) {
	router.Register("/test", &testHandler{})
	router.Register("/test/:id", &testHandler{})
	c := make(chan bool, 1)
	srv, err := NewHTTPServerWithoutServer(c)
	assert.NoError(t, err)
	assert.NotNil(t, srv)
	select {
	case b := <- c:
		assert.True(t, b)
	case <-time.After(1 * time.Second):
		panic(errors.New("Time out"))
	}

	resp, err := http.Get("http://localhost:4621/test/1")
	assert.NoError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, string(body), "Test Pass! 1")

	resp, err = http.Get("http://localhost:4621/test")
	assert.NoError(t, err)

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, string(body), "Test Pass!")
}