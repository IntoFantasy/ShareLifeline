package main

import (
	"database2"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type result struct { //定义返回数据格式
	Code int
	Msg  string
	Data []string
}

var dict map[string]string

func setCookie(w http.ResponseWriter, r *http.Request) (ret string) {
	// 定义两个cookie
	c1 := http.Cookie{
		Name:  "first_cookie",
		Value: "Go Programming",
	}
	randomcode := strconv.Itoa(rand.Int())
	fmt.Print(randomcode)
	c2 := http.Cookie{
		Name:     "randomcode",
		Value:    randomcode,
		HttpOnly: true,
	}
	// 设置Set-Cookie字段
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
	return randomcode
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
func main() {
	rand.Seed(time.Now().UnixNano())
	database2.InitDB()
	database2.DeleteAllCookie()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./view/static/"))))
	http.HandleFunc("/index", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/homepage", homepage)
	http.HandleFunc("/change", show_change)
	http.HandleFunc("/modify", change)
	err := http.ListenAndServe("127.0.0.1:9300", nil)
	if err != nil {
		return
	} //设置监听的端口
}
func show_change(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		write, _ := template.ParseFiles("view/change.html")
		err := write.Execute(w, nil)
		if err != nil {
			return
		}
	}

}
func GetInfo(r *http.Request, info_name string) (info string) {
	x, _ := r.Form[info_name]
	info = x[0]
	return info
}
func change(w http.ResponseWriter, r *http.Request) {
	cookie := getCookie(w, r)
	email1 := database2.FindCookie(cookie)
	//email1 := "525802191@qq.com"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	fmt.Println("change")
	if r.Method == "POST" {
		//name := GetInfo(r, "name")
		//password := GetInfo(r, "password")
		//email := GetInfo(r, "email")
		//address := GetInfo(r, "address")
		//birthday := GetInfo(r, "birthday")
		//introduction := GetInfo(r, "introduction")
		//phone := GetInfo(r, "phone")
		//sex := GetInfo(r, "sex")
		id := database2.QueryUserSingle("email", email1, "id")
		var info1 map[string]string
		info1 = make(map[string]string)
		key := []string{"name", "email", "password", "occupation", "birthday", "introduction", "disease", "sex"}
		var tmp string
		for _, val := range key {
			tmp = GetInfo(r, val)
			if tmp != "" {
				info1[val] = tmp
			}
		}
		//info1 := map[string]string{
		//	"name":         name,
		//	"email":        email,
		//	"password":     password,
		//	"address":      address,
		//	"birthday":     birthday,
		//	"introduction": introduction,
		//	"phone":        phone,
		//	"sex":          sex,
		//}
		database2.ModifyUser(id, info1)
		arr := &result{
			200,
			"change success",
			[]string{},
		}
		http.Redirect(w, r, "/homepage", http.StatusFound)
		b, jsonErr := json.Marshal(arr) //json化结果集
		if jsonErr != nil {
			fmt.Println("encoding faild")
		} else {
			io.WriteString(w, string(b)) //返回结果
			fmt.Println(string(b))
		}
	}
}

type question struct {
	Id      string
	Title   string
	Content string
	Html    string
}

type info struct {
	Name         string
	Sex          string
	Birthday     string
	Address      string
	Age          string
	Email        string
	Introduction string
	Photo        template.URL
	Labels       string
	Occupation   string
	Disease      string
	Diseaseage   string
	Question     []question
}

func homepage(w http.ResponseWriter, r *http.Request) {
	cookie := getCookie(w, r)
	fmt.Println("homepage")
	email := database2.FindCookie(cookie)
	information := database2.QueryUser("email", email, []string{"name", "sex", "address", "birthday", "introduction",
		"photo", "phone", "tag", "occupation", "disease", "diseaseage", "id"})
	var user info
	user.Name = information[0]
	user.Sex = information[1]
	user.Address = information[2]
	user.Birthday = information[3]
	user.Introduction = information[4]
	user.Photo = template.URL(database2.QueryPicture(information[5]))
	user.Age = information[6]
	user.Email = email
	user.Labels = information[7]
	user.Occupation = information[8]
	user.Disease = information[9]
	user.Diseaseage = information[10]
	id := information[11]
	pids := database2.QueryPersonalPost(id)
	for _, val := range pids {
		var ques question
		m := database2.QueryPost(val)
		ques.Title = m["title"]
		ques.Content = m["content"]
		ques.Id = val
		ques.Html = m["html"]
		user.Question = append(user.Question, ques)
	}
	//user.Labels = strings.Split(information[7], "|")
	//if len(user.Labels) == 1 && user.Labels[0] == "" {
	//	user.Labels[0] = "No Tags Yet"
	//}
	if r.Method == "GET" {
		write, _ := template.ParseFiles("view/people.html")
		err := write.Execute(w, user)
		if err != nil {
			return
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie := getCookie(w, r)
	if cookie == "" {
		setCookie(w, r)
	}
	if r.Method == "GET" {
		write, _ := template.ParseFiles("view/login.html")
		err := write.Execute(w, nil)
		if err != nil {
			return
		}
	}

}
func login(w http.ResponseWriter, r *http.Request) {
	cookie := getCookie(w, r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	if r.Method == "POST" {
		name, _ := r.Form["nameSignUp"]
		email, _ := r.Form["emailSignUp"]
		password := r.Form["passwordSignUp"]
		email_, _ := r.Form["emailLogin"]
		password_, _ := r.Form["passwordLogin"]
		nameSignUp := name[0]
		emailSignUp := email[0]
		passwordSignUp := password[0]
		if (nameSignUp != "" || emailSignUp != "" || passwordSignUp != "") && (len(email_) != 0 || len(password_) != 0) {
			if len(name) != 0 {
				fmt.Println("name:", name[0])
			}
			arr := &result{
				400,
				"fail, please don't fill in both sign up and login",
				[]string{},
			}
			b, jsonErr := json.Marshal(arr) //json化结果集
			if jsonErr != nil {
				fmt.Println("encoding faild")
			} else {
				io.WriteString(w, string(b)) //返回结果
				fmt.Println(string(b))
			}
		} else if nameSignUp == "" {
			emailLogin := email_[0]
			passwordLogin := password_[0]
			passwordRight := database2.QueryUserSingle("email", emailLogin, "password")
			if passwordLogin == passwordRight {
				fmt.Println("ok")
				email1 := database2.FindCookie(cookie)
				fmt.Println(email1, cookie, emailLogin)
				if email1 == "" {
					database2.AddCookie(emailLogin, cookie)
				} else {
					if email1 != emailLogin {
						database2.ChangeCookie(cookie, emailLogin)
					}
				}
				arr := &result{
					200,
					"login success",
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
		} else {
			if database2.IfExist(emailSignUp, "email", "user") {
				arr := &result{
					300,
					"sign up fail, because the email has been signed up",
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
				info := map[string]string{
					"name":     nameSignUp,
					"password": passwordSignUp,
					"email":    emailSignUp,
				}
				_ = database2.NewUser(info)
				cookie := getCookie(w, r)
				fmt.Println(cookie)
				if database2.IfExist(cookie, "randomcode", "cookie") {
					database2.ChangeCookie(cookie, emailSignUp)
				} else {
					database2.AddCookie(emailSignUp, cookie)
				}
				arr := &result{
					200,
					"sign up success",
					[]string{},
				}
				http.Redirect(w, r, "http://127.0.0.1:9500/searchResult?key=", http.StatusFound)
				b, jsonErr := json.Marshal(arr) //json化结果集
				if jsonErr != nil {
					fmt.Println("encoding faild")
				} else {
					io.WriteString(w, string(b)) //返回结果
					fmt.Println(string(b))
				}
			}
		}
	}
}
