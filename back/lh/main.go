package main

import (
	"lh/fabric"
	"lh/http"
)

func main() {

	fabric.ConTest()
	http.Run_HttpServer("", "")

}
