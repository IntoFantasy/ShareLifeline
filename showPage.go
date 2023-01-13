package main

import (
	"database2"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func main() {
	database2.InitDB()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./view/static/"))))
	http.HandleFunc("/", showPage)
	http.HandleFunc("/question", showquestion)
	http.HandleFunc("/blog", showblog)
	err := http.ListenAndServe("127.0.0.1:9400", nil)
	if err != nil {
		return
	}
}

type condition struct {
	Name         string
	Introduction string
	Symptom      string
	Treatment    string
}

func showPage(w http.ResponseWriter, r *http.Request) {
	info := strings.Split(r.URL.String(), "/")
	if info[1] == "condition" {
		id := info[2]
		ret := database2.QueryCondition(id)
		for _, val := range ret {
			fmt.Println(val)
		}
		var send condition
		send.Name = ret[0]
		send.Introduction = ret[1]
		send.Symptom = ret[2]
		send.Treatment = ret[3]
		if r.Method == "GET" {
			write, _ := template.ParseFiles("view/condition.html")
			err := write.Execute(w, send)
			if err != nil {
				return
			}
		}
	} else if info[1] == "treatment" {
		id := info[2]
		ret := database2.QueryTreatment(id)
		var send condition
		send.Name = ret[0]
		send.Introduction = ret[1]
		send.Symptom = ret[2]
		send.Treatment = ret[3]
		if r.Method == "GET" {
			write, _ := template.ParseFiles("view/treatment.html")
			err := write.Execute(w, send)
			if err != nil {
				return
			}
		}
	}
}

type question struct {
	Content string
	Title   string
	Html    string
}

func showquestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	id := r.Form["idx"]
	ret := database2.QueryQuestion(id[0])
	var resp question
	resp.Content = ret["content"]
	resp.Title = ret["title"]
	if r.Method == "GET" {
		write, _ := template.ParseFiles("view/question.html")
		err := write.Execute(w, resp)
		if err != nil {
			return
		}
	}
}
func showblog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	id := r.Form["idx"]
	ret := database2.QueryBlog(id[0])
	var resp question
	resp.Content = ret["content"]
	resp.Title = ret["title"]
	if r.Method == "GET" {
		write, _ := template.ParseFiles("view/blog.html")
		err := write.Execute(w, resp)
		if err != nil {
			return
		}
	}
}
