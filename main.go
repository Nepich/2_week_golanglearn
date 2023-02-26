package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8000"

func main() {
	http.HandleFunc("/get_pic", GetPicURLHandler)
	http.HandleFunc("/get_pic_list", GetPicListURLHandler)

	fmt.Printf("server started at http://127.0.0.1%s\n", port)

	if err := http.ListenAndServe(port, nil); err!=nil{
		log.Fatal(err)
	}
}