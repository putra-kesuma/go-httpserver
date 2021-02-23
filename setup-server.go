package httpserver

import (
	"net"
	"net/http"
)

//RunServer is default listen with network tcp
// address exaple 127.0.0.1:8080
func RunServer(address string) error {
	listenServe, err := net.Listen("tcp", address)

	if err != nil {
		return err
	}

	if err = http.Serve(listenServe, nil); err != nil {
		return err
	}

	return nil
}
