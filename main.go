package main

import (
	"fmt"
	"net/http"
)
import "filestore-server/handler"

func main(){
	http.HandleFunc("/file/upload", handler.UploadHandler )
	http.HandleFunc("/file/upload/success", handler.UploadSuccessHandler)
	http.HandleFunc("/file/query", handler.QueryFileBySha1Handler)
	http.HandleFunc("/file/batch", handler.QueryFileInBatch)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/delete", handler.DeleteHandler)
	http.HandleFunc("/file/rename", handler.RenameFile)

	err:= http.ListenAndServe(":8080", nil)
	if err!= nil{
		fmt.Print(err.Error())
	}
}