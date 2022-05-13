package handler

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"filestore-server/meta"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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
		localFile, err := os.OpenFile(fileDestination, os.O_RDWR |os.O_CREATE, 0666)
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

		hash := sha1.New()
		if _, err := io.Copy(hash, localFile); err != nil{
			fmt.Println(err)
		}
		sum := hash.Sum(nil)
		fileMeta.FileSha1 = hex.EncodeToString(sum)
		println("新增文件的sha1值为:", fileMeta.FileSha1)
		meta.UpdateFileMeta(fileMeta)

		http.Redirect(w,r, "/file/upload/success", http.StatusFound)
	}
}

// windows下certutil -hashfile 文件名 SHA1
func QueryFileBySha1Handler(w http.ResponseWriter, r* http.Request) {
	r.ParseForm()

	key := r.Form["filehash"][0]
	println("查询请求里的hash:",key)
	fileMeta := meta.GetFileMeta(key)
	data, err := json.Marshal(fileMeta)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func QueryFileInBatch(w  http.ResponseWriter, r * http.Request){
	r.ParseForm()
	limit, _ := strconv.Atoi(r.Form["limit"][0])
	data, _ := json.Marshal(meta.GetFileMetas(limit))
	w.Write(data)
}