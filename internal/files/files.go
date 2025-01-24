package files

import (
	"fmt"
	"os"
)


func ReadPath()(string,error){
	if len(os.Args)<2{
		//:TODO
		return "",fmt.Errorf("error during reading path from os args possibaly since the path was not given")
	}
	return os.Args[1],nil
}

func ReadFile(path string)([]byte,error){
	buff,buffErr:=os.ReadFile(path)
	if buffErr!=nil{
		return nil,fmt.Errorf("error during reading file %v",buffErr)
	}
	return buff,nil
}
