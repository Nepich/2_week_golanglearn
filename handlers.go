package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

const URL = "https://api.waifu.pics/"

var (
	categories      = []string{"sfw", "nsfw"}
	subcategories = []string{"waifu", "neko", "shinobu", "megumin", "bully", "cuddle"}
)

type Request struct {
	URL         string
	Category    string
	Subcategory string
}

func NewRequest() Request {
	rand.Seed(time.Now().UnixNano())
	choice := rand.Intn(len(subcategories))
	return Request{
		URL:      URL,
		Category: categories[0],
		Subcategory: subcategories[choice],
	}
}

func GetPicURLHandler(w http.ResponseWriter, r *http.Request) {
	request := NewRequest()

	resp, err := http.Get(fmt.Sprintf("%s%s%s", request.URL, request.Category, request.Subcategory))
	if err != nil {
		fmt.Println(err)
	}

	pic, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	w.Write(pic)
}

func GetPicListURLHandler(w http.ResponseWriter, r *http.Request) {
	request := NewRequest()
	fmt.Printf("%smany/%s/%s\n", request.URL, request.Category, request.Subcategory)

	req := []byte(`{}`)

	resp, err := http.Post(fmt.Sprintf("%smany/%s/%s", request.URL, request.Category, request.Subcategory), "application/json", bytes.NewBuffer(req))
	if err != nil {
		fmt.Println(err)
	}
	pic, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	w.Write(pic)
}
