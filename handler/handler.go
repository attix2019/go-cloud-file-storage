package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func UploadSuccessHandler(w http.ResponseWriter, r* http.Request){
	io.WriteString(w, "upload succeeded")
}

func UploadHandler(w http.ResponseWriter, r* http.Request){
	if r.Method =="GET" {
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil{
			fmt.Printf("error occured when reading index.html : %s\n" + err.Error())
		}
		w.Write(data)
	}else if r.Method =="POST"{

		receivedFile, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("error occured during uploading file: %s\n", err.Error())
		}
		defer receivedFile.Close()

		localFilePath := filepath.Join("tmp")
		os.MkdirAll(localFilePath, os.ModePerm)
		localFile, err := os.OpenFile(filepath.Join(localFilePath, handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil{
			fmt.Printf("error occured during opening file: %s\n" + err.Error() )
		}
		defer localFile.Close()

		_, err = io.Copy(localFile, receivedFile)
		if err != nil{
			fmt.Printf("error occured during saving file : %s\n" + err.Error())
		}

		http.Redirect(w,r, "/file/upload/success", http.StatusFound)
	}
}