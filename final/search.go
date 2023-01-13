package main

import (
	"database2"
	"fmt"
	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
	"html/template"
	"net/http"
)

var (
	// searcher is coroutine safe
	searcher          = riot.Engine{}
	searcher_question = riot.Engine{}
	searcher_blog     = riot.Engine{}
	searcher_user     = riot.Engine{}
)

func init_searcher() {
	searcher.Init(types.EngineOpts{
		// Using:             4,
		NotUseGse: true,
	})
	defer searcher.Close()
	//_ = database2.InitDB()
	dict := database2.GetTreatmentText()
	for idx, text := range dict {
		searcher.Index(idx, types.DocData{Content: text, Labels: []string{"Treatment"}})
	}
	dict_condition := database2.GetConditionText()
	for idx, text := range dict_condition {
		tmp := database2.String2Int(idx) + 1e8
		searcher.Index(database2.Int2String(tmp), types.DocData{Content: text, Labels: []string{"Condition"}})
	}
	searcher.Flush()
}

func init_searcher_question() {
	searcher_question.Init(types.EngineOpts{
		// Using:             4,
		NotUseGse: true,
	})
	defer searcher_question.Close()
	//_ = database2.InitDB()
	dict := database2.GetQuestionText()
	for idx, text := range dict {
		searcher_question.Index(idx, types.DocData{Content: text})
	}
	searcher_question.Flush()
}
func init_searcher_blog() {
	searcher_blog.Init(types.EngineOpts{
		// Using:             4,
		NotUseGse: true,
	})
	defer searcher_blog.Close()
	//_ = database2.InitDB()
	dict := database2.GetBlogText()
	for idx, text := range dict {
		searcher_blog.Index(idx, types.DocData{Content: text})
	}
	searcher_blog.Flush()
}
func init_searcher_user() {
	searcher_user.Init(types.EngineOpts{NotUseGse: true})
	defer searcher_user.Close()
	dict := database2.GetUserText()
	for idx, text := range dict {
		searcher_user.Index(idx, types.DocData{Content: text})
	}
	searcher_user.Flush()
}
func main() {
	database2.InitDB()
	init_searcher()
	init_searcher_question()
	init_searcher_blog()
	init_searcher_user()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./view/static/"))))
	http.HandleFunc("/searchResult", SearchResult)
	http.HandleFunc("/searchquestion", SearchQuestion)
	http.HandleFunc("/searchblog", SearchBlog)
	http.HandleFunc("/searchuser", SearchUser)
	err := http.ListenAndServe("127.0.0.1:9500", nil)
	if err != nil {
		return
	} //设置监听的端口
}

type knowledge struct {
	Idx          int
	Kind         string //treatment or condition
	Name         string
	Introduction string
	Uses         string //condition 应为 symptom
	Sideeffect   string // condition 应为 treatment
}
type response struct {
	Knowledge []knowledge
}

func SearchResult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	//for k, v := range r.Form {
	//	fmt.Println(k, v)
	//}
	tmp := r.Form["key"]
	strings := "Diet"
	if len(tmp) != 0 {
		strings = r.Form["key"][0]
	}
	if strings == "" {
		strings = "Diet"
	}
	output := searcher.SearchDoc(types.SearchReq{Text: strings})
	var resp response
	for _, doc := range output.Docs {
		var know knowledge
		id := database2.String2Int(doc.DocId)
		var ret []string
		if id > 1e8 {
			id -= 1e8
			ret = database2.QueryCondition(database2.Int2String(id))
			know.Kind = "condition"
		} else {
			ret = database2.QueryTreatment(database2.Int2String(id))
			know.Kind = "treatment"
		}
		know.Idx = id
		know.Name = ret[0]
		know.Introduction = ret[1]
		know.Uses = ret[2]
		know.Sideeffect = ret[3]
		resp.Knowledge = append(resp.Knowledge, know)
	}
	//fmt.Println(resp.Knowledge[0].Name)
	if r.Method == "GET" {
		fmt.Println("123")
		write, _ := template.ParseFiles("view/searchResult.html")
		err := write.Execute(w, resp)

		if err != nil {
			return
		}
	}
}

type question struct {
	Id      string
	Title   string
	Content string
}
type question_resp struct {
	Question []question
}

func SearchQuestion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("1")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	tmp := r.Form["key"]
	strings := "Diet"
	if len(tmp) != 0 {
		strings = r.Form["key"][0]
	}
	if strings == "" {
		strings = "Diet"
	}
	output := searcher_question.SearchDoc(types.SearchReq{Text: strings})
	var resp question_resp
	for _, doc := range output.Docs {
		var ques question
		ques.Id = doc.DocId
		ret := database2.QueryQuestion(ques.Id)
		ques.Title = ret["title"]
		ques.Content = ret["content"]
		resp.Question = append(resp.Question, ques)
	}
	if r.Method == "GET" {
		fmt.Println("123")
		write, _ := template.ParseFiles("view/searchQuestion.html")
		err := write.Execute(w, resp)

		if err != nil {
			return
		}
	}
}
func SearchBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	tmp := r.Form["key"]
	strings := "Diet"
	if len(tmp) != 0 {
		strings = r.Form["key"][0]
	}
	if strings == "" {
		strings = "Diet"
	}

	output := searcher_blog.SearchDoc(types.SearchReq{Text: strings})
	var resp question_resp
	for _, doc := range output.Docs {
		var ques question
		ques.Id = doc.DocId
		ret := database2.QueryBlog(ques.Id)
		ques.Title = ret["title"]
		ques.Content = ret["content"]
		fmt.Println(ques.Title, ques.Content)
		resp.Question = append(resp.Question, ques)
	}
	if r.Method == "GET" {
		fmt.Println("123")
		write, _ := template.ParseFiles("view/searchBlog.html")
		err := write.Execute(w, resp)
		if err != nil {
			return
		}
	}
}

type user struct {
	Name  string
	Email string
	Id    string
}
type user_resp struct {
	Users []user
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	tmp := r.Form["key"]
	strings := "Diet"
	if len(tmp) != 0 {
		strings = r.Form["key"][0]
	}
	if strings == "" {
		strings = "Diet"
	}
	output := searcher_user.SearchDoc(types.SearchReq{Text: strings})
	var resp user_resp
	for _, doc := range output.Docs {
		var user1 user
		user1.Id = doc.DocId
		user1.Email = database2.QueryUserSingle("id", user1.Id, "email")
		user1.Name = database2.QueryUserSingle("id", user1.Id, "name")
		fmt.Println(user1.Name)
		resp.Users = append(resp.Users, user1)
	}
	if r.Method == "GET" {
		fmt.Println("123")
		write, _ := template.ParseFiles("view/searchUser.html")
		err := write.Execute(w, resp)
		if err != nil {
			return
		}
	}
}
