package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./view/static/"))))
	http.HandleFunc("/save", save)
	http.HandleFunc("/release", release)
	err := http.ListenAndServe("127.0.0.1:9300", nil)
	if err != nil {
		return
	} //设置监听的端口
}

func save(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	if r.Method == "POST" {
		editor, _ := r.Form["editor"]
		editor_ := editor[0]
		Len := bytes.Count([]byte(editor_), nil) - 1
		editor_ = editor_[39 : Len-7]
		fmt.Println(editor_)
	}
}

func release(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	if r.Method == "POST" {
		editor, _ := r.Form["editor"]
		editor_ := editor[0]
		Len := bytes.Count([]byte(editor_), nil) - 1
		editor_ = editor_[39 : Len-7]
		fmt.Println(editor_)
	}
}
