package httpserver

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

//RunServer is default listen with network tcp
// address exaple 127.0.0.1:8080
// recive two parameter address string and handler http.Handler
func RunServer(address string, handler http.Handler) error {
	listenServe, err := net.Listen("tcp", address)

	if err != nil {
		return err
	}

	log.Println("Server is running")
	fmt.Println("Awesome server is Runing ", address)

	if err = http.Serve(listenServe, handler); err != nil {
		return err
	}

	return nil
}
