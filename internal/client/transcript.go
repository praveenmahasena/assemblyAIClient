package client

import (
	"fmt"
	"net"
)

func GetTransScript(fb []byte, scriptCh chan<- string, errorCh chan<- error) {
	defer close(scriptCh)
	defer close(errorCh)

	con, conErr := net.Dial("tcp", ":42069")


	if conErr != nil {
		errorCh <- fmt.Errorf("error during dialing to tcp server %v", conErr)
		return
	}
	defer con.Close()

	if _, fbErr := con.Write(fb); fbErr != nil {
		errorCh <- fmt.Errorf("error during sending buffer to tcp server %v", fbErr)
		return
	}
}
