package main

import (
	"fmt"
	"net/http"
)
import "filestore-server/handler"

func main(){
	http.HandleFunc("/file/upload", handler.UploadHandler )
	http.HandleFunc("/file/upload/success", handler.UploadSuccessHandler)

	err:= http.ListenAndServe(":8080", nil)
	if err!= nil{
		fmt.Print(err.Error())
	}
}