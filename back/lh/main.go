package main

import (
	"lh/fabric"
	"lh/http"
	"os"
)

func main() {
	os.Setenv("testenv", "TRUE")
	//DISCOVERY_AS_LOCALHOST=TRUE  필수[env]
	os.Setenv("DISCOVERY_AS_LOCALHOST", "TRUE")
	fabric.Init_Wallet()
	fabric.ConTest()
	http.Run_HttpServer("", "")

}
