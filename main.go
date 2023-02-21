package main

import (
	. "golang_server/context"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	/*
		if len(os.Args) < 3 {
			panic("len(os.Args) < 3")
		}

		config_server_type := os.Args[1:2]
		config_path := os.Args[2:3]
	*/
	server_context := GetServerContext()
	server_context.Init(20000)
	server_context.Listen()

}
