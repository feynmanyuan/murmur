package server

// you must cd [project dir]/bytemurmur.com/server
// and exec command : go-bindata  -o=components/asset.go -pkg=asset static/...

import "testing"

func TestNewHTTPServer(t *testing.T) {
	c := make(chan bool, 1)
	NewHTTPServer()
	<-c
}