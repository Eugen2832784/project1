package main

import (
	"net/http"
)

func main(){
	http.HandleFunc("/", handler )
}

func handler(w http.ResponseWriter, r *http.Request){
	str:= "Hello, DevOps!"
	b:= []byte(str)
	_,err:= w.Write(b)
	if err!= nil{
		
	}
}