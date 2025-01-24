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
	// req,reqErr:=http.NewRequest(http.MethodPost,"https://api.assemblyai.com/v2/upload",bytes.NewReader(fileBytes))
	// req.Header.Add("Authorization", "5a69766da1df4837a6d6a225dda9b555")
	// req.Header.Add("Content-Type" ,"application/octet-stream")
	//
	// if reqErr!=nil{
	// 	return reqErr
	// }
	// res,resErr:=http.DefaultClient.Do(req)
	// if resErr!=nil{
	// 	return resErr
	// }
	// defer res.Body.Close()
	// data,dataErr:=io.ReadAll(res.Body)
	// if dataErr!=nil{
	// 	return dataErr
	// }
	// fmt.Println(string(data))
	// return nil
	//
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
