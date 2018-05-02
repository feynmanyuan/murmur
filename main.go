package main

import (
	"bytemurmur.com/server"
	"bytemurmur.com/ui"
	"os"
)

func main() {
	server.NewHTTPServer()
	ui.NewMainFrame(os.Args)
}