package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", handler )

	err := http.ListenAndServe(":8080", nil)
	if err!= nil{
		fmt.Println("error! ", err.Error())
	}

	fmt.Println("Program has ended it work")
}

func handler(w http.ResponseWriter, r *http.Request){
	str:= "Hello, DevOps!"
	b:= []byte(str)
	_,err:= w.Write(b)
	if err!= nil{
		fmt.Println("error: ", err.Error())
	} else{
		fmt.Println("success!")
	}
}