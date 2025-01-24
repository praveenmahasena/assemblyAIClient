package internal

import (
	"fmt"

	"github.com/praveenmahasena/client/internal/client"
	"github.com/praveenmahasena/client/internal/files"
)

func Start() error {
	path, pathErr := files.ReadPath()
	if pathErr != nil {
		return pathErr
	}
	fileBytes, byteErr := files.ReadFile(path)
	if byteErr != nil {
		return byteErr
	}

	scriptCh := make(chan string)
	errorCh := make(chan error)

	go client.GetTransScript(fileBytes, scriptCh, errorCh)

	select {
	case tranScript := <-scriptCh:
		fmt.Println(tranScript)
	case err := <-errorCh:
		return err
	}

	return nil
}
