package main

import (
	"MyPack/database2"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type result struct { //定义返回数据格式
	Code int
	Msg  string
	Data []string
}

func main() {
	database2.InitDB()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./view/static/"))))
	http.HandleFunc("/index", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/homepage", homepage)
	err := http.ListenAndServe("127.0.0.1:9300", nil)
	if err != nil {
		return
	} //设置监听的端口
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		write, _ := template.ParseFiles("view/login.html")
		err := write.Execute(w, nil)
		if err != nil {
			return
		}
	}
}

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	//strings := r.Form["email"]
	fmt.Println("homepage")
	//fmt.Println(strings)
	if r.Method == "GET" {
		write, _ := template.ParseFiles("view/homepage.html")
		err := write.Execute(w, nil)
		if err != nil {
			return
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	if r.Method == "POST" || r.Method == "GET" {
		email, _ := r.Form["emailLogin"]
		emailLogin := email[0]
		password, _ := r.Form["passwordLogin"]
		passwordLogin := password[0]
		realPassword := database2.QueryUserSingle("email", emailLogin, "password")
		if passwordLogin == realPassword {
			//http.Redirect(w, r, "/homepage", http.StatusTemporaryRedirect)
			arr := &result{
				200,
				"登陆成功",
				[]string{},
			}
			b, jsonErr := json.Marshal(arr) //json化结果集
			if jsonErr != nil {
				fmt.Println("encoding faild")
			} else {
				io.WriteString(w, string(b)) //返回结果
				fmt.Println(string(b))
			}
		} else {
			fmt.Println("p")
			_, err := w.Write([]byte("<script>alert('wrong')</script>"))
			if err != nil {
				return
			}
		}
	}
}
