package handler

import (
	"filestore-server/meta"
	"filestore-server/util"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
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

		receivedFile, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("error occured during uploading file: %s\n", err.Error())
		}
		defer receivedFile.Close()

		localFilePath := filepath.Join("tmp")
		os.MkdirAll(localFilePath, os.ModePerm)
		fileDestination := filepath.Join(localFilePath, head.Filename)
		localFile, err := os.OpenFile(fileDestination, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil{
			fmt.Printf("error occured during opening file: %s\n" + err.Error() )
		}
		defer localFile.Close()

		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: fileDestination,
			UploadAt: time.Now().Format("2022-05-13 17:07:07"),
		}
		fileMeta.FileSize, err = io.Copy(localFile, receivedFile)
		if err != nil{
			fmt.Printf("error occured during saving file : %s\n" + err.Error())
		}
		localFile.Seek(0,0)
		fileMeta.FileSha1 = util.FileSha1(localFile)
		meta.UpdateFileMeta(fileMeta)

		http.Redirect(w,r, "/file/upload/success", http.StatusFound)
	}
}