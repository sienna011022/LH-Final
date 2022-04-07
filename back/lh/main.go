package main

import (
	"lh/http"
)

func main() {

	/*var (
		drivername string
		db_url     string
	)
	go get_af_data(drivername, db_url)*/
	//http.StartWoker(3)
	http.Run_HttpServer("", "")

}
