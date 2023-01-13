package main

import (
	"bytes"
	"database2"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Comment struct {
	Name       string
	TimeBefore string
	IsAuthor   bool
	Content    string
	Portal     string
	IsSelf     bool
}
type CommentPair struct {
	MainLevel Comment
	Reply     []Comment
}
type CommentData struct {
	Comments     []CommentPair
	Photo        template.URL
	Name         string
	Disease      string
	Birthday     string
	Diseaseage   string
	Labels       string
	Occupation   string
	Introduction string
	Email        string
}

func main() {
	database2.InitDB()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./view/static/"))))
	http.HandleFunc("/Question", question)
	http.HandleFunc("/BlogWrite", blogWrite)
	http.HandleFunc("/Query", query)
	http.HandleFunc("/FAQ", FAQ)
	http.HandleFunc("/CommentSave", CommentSave)
	http.HandleFunc("/save", save)
	http.HandleFunc("/FollowPost", followPost)
	err := http.ListenAndServe("127.0.0.1:9700", nil)
	if err != nil {
		return
	} //设置监听的端口
}

func question(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		write, _ := template.ParseFiles("view/question.html")
		err := write.Execute(w, nil)
		if err != nil {
			return
		}
	}
}

func blogWrite(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		write, _ := template.ParseFiles("view/blogWrite.html")
		err := write.Execute(w, nil)
		if err != nil {
			return
		}
	}
}
func query(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		write, _ := template.ParseFiles("view/query.html")
		err := write.Execute(w, nil)
		if err != nil {
			return
		}
	}
}
func getCookie(w http.ResponseWriter, r *http.Request) (cookie string) {
	cookie = ""
	cookies := strings.Split(r.Header.Get("Cookie"), "; ")
	for _, value := range cookies {
		if strings.Contains(value, "randomcode") {
			cookie = strings.Split(value, "=")[1]
		}
	}
	return cookie
}
func FAQ(w http.ResponseWriter, r *http.Request) {
	cookie := getCookie(w, r)
	email := database2.FindCookie(cookie)
	fmt.Println(email)
	information := database2.QueryUser("email", email, []string{"name", "birthday", "introduction",
		"photo", "tag", "occupation", "disease", "diseaseage", "email"})
	if r.Method == "GET" {
		comment := Comment{
			Name:       "self",
			TimeBefore: "20 Minus",
			IsAuthor:   true,
			Content:    "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Velit omnis animi et iure laudantium vitae, praesentium optio, sapiente distinctio illo?",
			Portal:     database2.QueryPicture(information[3]),
			IsSelf:     true,
		}
		comment2 := Comment{
			Name:       "Guest",
			TimeBefore: "10 Minus",
			IsAuthor:   false,
			Content:    "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Velit omnis animi et iure laudantium vitae, praesentium optio, sapiente distinctio illo?",
			Portal:     "static/FAQ/picture/re.jpg",
			IsSelf:     false,
		}
		var reply []Comment
		reply = append(reply, comment2, comment)
		pair := CommentPair{comment, reply}
		pair2 := CommentPair{comment2, reply}
		pairs := []CommentPair{pair, pair2}
		var data CommentData
		data.Comments = pairs
		data.Name = information[0]
		data.Birthday = information[1]
		data.Introduction = information[2]
		data.Photo = template.URL(database2.QueryPicture(information[3]))
		data.Labels = information[4]
		data.Occupation = information[5]
		data.Disease = information[6]
		data.Diseaseage = information[7]
		data.Email = information[8]
		write, _ := template.ParseFiles("view/FAQ.html")
		err := write.Execute(w, data)
		if err != nil {
			return
		}
	}
}

func CommentSave(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" || r.Method == "POST" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		r.ParseForm()
		comment := r.Form["Comment"][0]
		fmt.Println(comment)
	}
}
func save(w http.ResponseWriter, r *http.Request) {
	fmt.Println("1")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	cookie := getCookie(w, r)
	email := database2.FindCookie(cookie)
	id := database2.QueryUserSingle("email", email, "id")
	if r.Method == "POST" {
		editor, _ := r.Form["editor"]
		editor1 := editor[0]
		Len := bytes.Count([]byte(editor1), nil) - 1
		editor1 = editor1[39 : Len-7]
		loc1 := strings.Index(editor1, "<h1>")
		loc2 := strings.Index(editor1, "</h1>")
		header1 := editor1[loc1+4 : loc2]
		fmt.Println(header1)
		loc3 := strings.Index(editor1, "<p>")
		loc4 := strings.Index(editor1, "</p>")

		content1 := editor1[loc3+3 : loc4]
		fmt.Println(content1)
		//var content1, header1 string
		//if len(content) == 0 {
		//	content1 = ""
		//} else {
		//	content1 = content[0]
		//}
		//if len(header) == 0 {
		//	header1 = ""
		//} else {
		//	header1 = header[0]
		//}
		fmt.Println(content1, id)
		database2.AddPost(map[string]string{"posterid": id, "title": header1, "content": content1, "html": editor1})
		fmt.Println(editor1, content1, header1)
	}
}

func followPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	if r.Method == "POST" || r.Method == "GET" {
		Type, _ := r.Form["type"]
		follower, _ := r.Form["follower"]
		followee, _ := r.Form["followee"]
		fmt.Println(Type, follower, followee)
	}
}
