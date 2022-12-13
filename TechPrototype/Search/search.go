package main

import (
	"MyPack/database2"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	database2.InitDB()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./view/static/"))))
	http.HandleFunc("/searchResult", SearchResult)
	http.HandleFunc("/search", Search)
	err := http.ListenAndServe("127.0.0.1:9300", nil)
	if err != nil {
		return
	} //设置监听的端口
}

func SearchResult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	strings := r.Form["searchContent"][0]
	fmt.Println(strings)
	if r.Method == "GET" {
		write, _ := template.ParseFiles("view/searchResult.html")
		err := write.Execute(w, nil)
		if err != nil {
			return
		}
	}
}

func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		write, _ := template.ParseFiles("view/search.html")
		err := write.Execute(w, nil)
		if err != nil {
			return
		}
	}
}
