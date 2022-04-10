package main

import (
	"lh/fabric"
	"lh/http"
)

func main() {
	//DISCOVERY_AS_LOCALHOST=TRUE  필수[env]
	fabric.Init_Wallet()
	fabric.ConTest()
	http.Run_HttpServer("", "")

}
