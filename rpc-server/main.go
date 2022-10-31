package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct{}

type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *string) error {
	// Fill reply pointer to send the data back;
	*reply = "time.Now().Unix()"
	return nil
}

func main() {
	timeserver := new(TimeServer)
	rpc.Register(timeserver)
	rpc.HandleHTTP()

	// Listen for requests on port 1234
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	http.Serve(listener, nil)
}
