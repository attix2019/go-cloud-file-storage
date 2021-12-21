package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r* http.Request){
	if r.Method =="GET" {
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil{
			fmt.Printf(err.Error())
		}
		w.Write(data)
	}else if r.Method =="POST"{
		//
	}
}