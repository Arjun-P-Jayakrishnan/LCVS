package handlers

import (
	//"bufio"
	//"io"
	//"fmt"
	"os"
)

func handleErr(err error){
	if(err!=nil){
		panic(err)
	}
}

func ReadFromFileAsBytes(path string) []byte {
	data,err:=os.ReadFile(path)
	handleErr(err)

	return data
}