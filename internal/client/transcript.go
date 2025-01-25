package client

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
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

	binary.Write(con,binary.LittleEndian,int64(len(fb)))
	io.CopyN(con,bytes.NewReader(fb),int64(len(fb)))


	r := struct {
		Status int    `json:"status"`
		Data   string `json:"data"`
	}{}
	if err := json.NewDecoder(con).Decode(&r); err != nil {
		errorCh <- fmt.Errorf("error during decoding data %v", err)
		return
	}

	scriptCh <- r.Data
}
