package main

import "github.com/christopher-alden/responsi/server"

func main(){

	//untuk running

	// bikin server baru
	srv := server.NewServer("localhost:1234")


	// suruh server buat run
	srv.Start()
}