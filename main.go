package main

import (
	"net/http"
	"io/ioutil"
)

func main () {
	http.ListenAndServe(":8080", http.FileServer(http.Dir("./")))
	//http.HandleFunc("/", home)
	//http.HandleFunc("/main.js", js)
	//http.HandleFunc("/main.css", css)
	//http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		body, _ := ioutil.ReadFile("index.html")
		w.Write(body)
		return
	}
	if r.Method == "POST" {
		// 一般没有
	}
}

func css(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadFile("main.css")
	w.Write(body)
}
func js(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadFile("main.js")
	w.Write(body)
}