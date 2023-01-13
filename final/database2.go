package database2

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id             int
	name           string
	email          string
	password       string
	address        string
	birthday       string
	introduction   string
	photo          string
	sex            string
	tag            string
	visit_his_num  int
	search_his_num int
	article_num    int
}

type visit struct {
	id      int `stbl:"id, PRIMARY_KEY, AUTO_INCREMENT"`
	post_id int `stbl:"post_id"`
}
type search struct {
	id      int    `stbl:"id, PRIMARY_KEY, AUTO_INCREMENT"`
	content string `stbl:"content"`
}
type my_post struct {
	id      int `stbl:"id, PRIMARY_KEY, AUTO_INCREMENT"`
	post_id int `stbl:"post_id"`
}
type answer struct {
	id         int
	answererid int
	questionid int
	answerid   int
	content    string
	likes      int
	view       int
	collects   int
}

func randint(min int, max int) int {
	return min + rand.Intn(max-min)
}
func Int2String(x int)(ret string){
	return strconv.Itoa(x)
}
func String2Int(s string)(ret int)  {
	ret, err := strconv.Atoi(s)
	check_err(err)
	return ret
}
func int64_to_int(theid int64) (id int) {
	id, _ = strconv.Atoi(strconv.FormatInt(theid, 10))
	return
}
func check_err(err any) (flag bool) {
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func QueryUser(keyword string, value any, query []string) (ret []string) {
	ret = make([]string, len(query))
	for idx, q := range query {
		sqlstr := "select " + q + " from user where " + keyword + "=?"
		var x string
		err := db.QueryRow(sqlstr, value).Scan(&x)
		if err != nil {
			print(err)
			return
		}
		ret[idx] = x
	}
	return ret
}
func QueryUserSingle(keyword string, value any, query string) (ret string) {
	sqlstr := "select " + query + " from user where " + keyword + "=?"
	err := db.QueryRow(sqlstr, value).Scan(&ret)
	flag := check_err(err)
	if !flag {
		return ""
	}
	return ret
}
func NewUser(info map[string]string) (id int) {
	id = -1
	sqlstr := "insert into user (name) values (?)"
	ret, err := db.Exec(sqlstr, "username")
	if err != nil {
		fmt.Println(err)
		return
	}
	theid, err := ret.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	id, _ = strconv.Atoi(strconv.FormatInt(theid, 10))
	fmt.Println("get id success")
	var flag = ModifyUser(id, info)
	if !flag {
		return -1
	}
	//searchsql := "create table " + strconv.Itoa(id) + "search (id int(20) primary key auto_increment, " +
	//	"searchcontent varchar(50))"
	//search, err := db.Prepare(searchsql)
	//if err != nil {
	//	fmt.Println(err)
	//	return -1
	//}
	//search.Exec()
	//
	//NewTable(id, "questioncollect", "questionid")
	//NewTable(id, "answercollect", "answerid")
	//NewTable(id, "postvisit", "postid")
	//NewTable(id, "post", "postid")
	//NewTable(id, "friend", "friendid")
	//NewTable(id, "postcollect", "postid")
	//NewTable(id, "question", "questionid")
	//NewTable(id, "answer", "answerid")
	//NewTable(id, "answervisit", "answerid")
	//NewTable(id, "questionvisit", "questionid")
	return
}
func NewTable(id int, table_name string, label string) {
	sql := "create table " + strconv.Itoa(id) + table_name + " (id int(20) primary key auto_increment, " +
		label + " int(20))"
	table, err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	table.Exec()
	return
}
func IfExist(info string, keyword string, tablename string)(ret bool){
	sqlstr := "select id from " + tablename + " where " + keyword + " =?"
	r, err := db.Query(sqlstr, info)
	check_err(err)
	defer  r.Close()
	return r.Next()
}
func InitDB() (err error) {
	dsn := "root:haohao0626@tcp(127.0.0.1:3305)/project"
	// 不会校验账号密码是否正确
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func ModifyUser(id any, info map[string]string) (ret bool) {
	sqlstr1 := "update user set "
	sqlstr2 := "=? where id=?"
	for key, value := range info {
		_, err := db.Exec(sqlstr1+key+sqlstr2, value, id)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	return true
}

func QueryOtherInfo(id string, key string, table_name string) (ret []string) { // 可以查询friend，collect，post，search，visit
	ret = make([]string, 0)
	sqlstr := "select " + key + " from " + id + table_name
	ids, _ := db.Query(sqlstr)
	var tmp string
	for ids.Next() {
		ids.Scan(&tmp)
		ret = append(ret, tmp)
	}
	return ret
}

func AddPost(info map[string]string) (flag bool) { // map中可以包含poster id，title，配图，必须包含正文。键的名字是posterid，title，picture，content
	sqlstr := "insert into posts (content, likes,collects,view) values (?,?,?,?)"
	ret, err := db.Exec(sqlstr, "", 0, 0, 0)
	flag = check_err(err)
	if !flag {
		return
	}
	theid, err := ret.LastInsertId()
	flag = check_err(err)
	if !flag {
		return
	}
	id := int64_to_int(theid)
	for key, value := range info {
		sqlstr = "update posts set " + key + "=? where id=?"
		_, err = db.Exec(sqlstr, value, id)
		flag = check_err(err)
		if !flag {
			return
		}
	}
	posterid, ok := info["posterid"]
	if ok {
		sqlstr = "insert into " + posterid + "post (postid) value (?)"
		_, err = db.Exec(sqlstr, id)
		flag = check_err(err)
		if !flag {
			return
		}
		article_num := QueryUserSingle("id", posterid, "article_num")
		article_num1, _ := strconv.Atoi(article_num)
		article_num1 += 1
		sqlstr := "update user set article_num=? where id=?"
		_, err = db.Exec(sqlstr, article_num1, posterid)
		flag = check_err(err)
		if !flag {
			return
		}
	}
	return true
}

func AddQuestion(info map[string]string) (flag bool) {
	sqlstr := "insert into questions (content, likes, collects, view) values (?,?,?,?)"
	ret, _ := db.Exec(sqlstr, "", 0, 0, 0)
	theid, _ := ret.LastInsertId()
	id := int64_to_int(theid)
	for key, value := range info {
		sqlstr = "update questions set " + key + "=? where id=?"
		_, err := db.Exec(sqlstr, value, id)
		flag = check_err(err)
		if !flag {
			return
		}
	}
	questionerid, ok := info["questionerid"]
	if ok {
		sqlstr = "insert into " + questionerid + "question (questionid) value (?)"
		_, err := db.Exec(sqlstr, id)
		flag = check_err(err)
		if !flag {
			return
		}
		article_num := QueryUserSingle("id", questionerid, "question_num")
		article_num1, _ := strconv.Atoi(article_num)
		article_num1 += 1
		sqlstr := "update user set question_num=? where id=?"
		_, err = db.Exec(sqlstr, article_num1, questionerid)
		flag = check_err(err)
		if !flag {
			return
		}
	}
	return true
}

func AddAnswer(info map[string]string) (flag bool) {
	sqlstr := "insert into answers (content, likes, collects, view) values (?,?,?,?)"
	ret, _ := db.Exec(sqlstr, "", 0, 0, 0)
	theid, _ := ret.LastInsertId()
	id := int64_to_int(theid)
	for key, value := range info {
		sqlstr = "update answers set " + key + "=? where id=?"
		_, err := db.Exec(sqlstr, value, id)
		flag = check_err(err)
		if !flag {
			return
		}
	}
	answererid, ok := info["answererid"]
	if ok {
		sqlstr = "insert into " + answererid + "answer (answerid) value (?)"
		_, err := db.Exec(sqlstr, id)
		flag = check_err(err)
		if !flag {
			return
		}
		article_num := QueryUserSingle("id", answererid, "answer_num")
		article_num1, _ := strconv.Atoi(article_num)
		article_num1 += 1
		sqlstr := "update user set answer_num=? where id=?"
		_, err = db.Exec(sqlstr, article_num1, answererid)
		flag = check_err(err)
		if !flag {
			return
		}
	}
	return true
}

func QueryAnswersToQuestion(questionid string) (ret []answer) {
	ret = make([]answer, 0)
	sqlstr := "select * from answers where questionid=" + questionid
	rows, _ := db.Query(sqlstr)
	var ans answer
	for rows.Next() {
		_ = rows.Scan(&ans.id, &ans.answererid, &ans.questionid, &ans.answerid,
			&ans.content, &ans.likes, &ans.view, &ans.collects)
		ret = append(ret, ans)
	}
	return ret
}

func QueryAnswersToAnswer(answerid string) (ret []answer) {
	ret = make([]answer, 0)
	sqlstr := "select * from answers where answerid=" + answerid
	rows, _ := db.Query(sqlstr)
	var ans answer
	for rows.Next() {
		_ = rows.Scan(&ans.id, &ans.answererid, &ans.questionid, &ans.answerid,
			&ans.content, &ans.likes, &ans.view, &ans.collects)
		ret = append(ret, ans)
	}
	return ret
}
func AddCookie(email string, cookie string) {
	sqlstr := "insert into cookie (email, randomcode) values (?, ?)"
	db.Exec(sqlstr, email, cookie)
}
func FindCookie(cookie string) (ret string){
	sqlstr := "select email from cookie where randomcode=?"
	db.QueryRow(sqlstr, cookie).Scan(&ret)
	return ret
}
func ChangeCookie(cookie string, email string){
	sqlstr := "update cookie set email=? where randomcode=?"
	_, err := db.Exec(sqlstr, email, cookie)
	if err != nil{
		fmt.Println(err)
	}
}
func DeleteAllCookie(){
	sqlstr := "truncate table cookie"
	_, err := db.Exec(sqlstr)
	if err != nil{
		fmt.Println(err )
	}
}
func QueryPicture(id string) (pic string){
	sqlstr := "select pic from picture where id=?"
	db.QueryRow(sqlstr, id).Scan(&pic)
	return pic
}
func QueryPersonalPost(id string)(ids []string){
	ids = make([]string, 0)
	sqlstr := "select postid from " + id + "post"
	rows, _ := db.Query(sqlstr)
	var pid string
	for rows.Next(){
		err := rows.Scan(&pid)
		check_err(err)
		ids = append(ids, pid)
	}
	return ids
}
func QueryPost(pid string)(info map[string]string){
	info = make(map[string]string)
	sqlstr := "select title, content, html from posts where id=?"
	var title, content, html string
	err := db.QueryRow(sqlstr, pid).Scan(&title, &content, &html)
	check_err(err)
	info["title"] = title
	info["content"] = content
	info["html"] = html
	return info
}
func AddTreatment(input []string){
	for {
		if (len(input) < 4){
			input = append(input, "")
		}else{
			break
		}
	}
	sqlstr := "insert into treatment (name, introduction, uses, sideeffect) values (?, ?, ?, ?)"
	db.Exec(sqlstr, input[0], input[1], input[2], input[3])
}

func AddCondition(input []string){
	for {
		if (len(input) < 4){
			input = append(input, "")
		}else{
			break
		}
	}
	sqlstr := "insert into condition222 (name, introduction, symptom, treatment) values (?, ?, ?, ?)"
	_, err := db.Exec(sqlstr, input[0], input[1], input[2], input[3])
	if err != nil{
		fmt.Println(err)
	}
}
func QueryCondition(id string)(info []string){
	sqlstr := "select * from condition222 where id=?"
	var ret []string
	ret = make([]string, 4)
	err := db.QueryRow(sqlstr, id).Scan(&id, &ret[0], &ret[1], &ret[2], &ret[3])
	check_err(err)
	return ret
}
func CreateUsers() {
	csvfile, _ := os.Open("users_5.csv")
	defer csvfile.Close()
	r := csv.NewReader(csvfile)
	var email = [8]string{"@qq.com", "@163.com", "@sjtu.edu.cn", "@126.com", "@yahoo.cn", "@sina.com",
		"@gmail.com", "@souhu.com"}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		info := make(map[string]string)
		info["name"] = record[0]
		info["tag"] = record[1]
		info["email"] = strconv.Itoa(randint(100000000, 1000000000)) + email[randint(0, 8)]
		info["password"] = strconv.Itoa(randint(100000, 1000000000))
		NewUser(info)
	}

}

func GetTreatmentText()(dict map[string]string){
	dict = make(map[string]string)
	sqlstr := "select * from treatment"
	rows, _ := db.Query(sqlstr)
	var id, name, introduction, uses, sideeffect string
	for rows.Next(){
		_ = rows.Scan(&id, &name, &introduction, &uses, &sideeffect)
		dict[id] = name + introduction + uses + sideeffect
	}
	return dict
}
func GetQuestionText()(dict map[string]string){
	dict = make(map[string]string)
	sqlstr := "select id, title, content from questions"
	rows, _ := db.Query(sqlstr)
	var id, title, content string
	for rows.Next(){
		_ = rows.Scan(&id, &title, &content)
		if title == "''"{
			title = ""
		}
		dict[id] = title + content
	}
	return dict
}
func GetBlogText()(dict map[string]string){
	dict = make(map[string]string)
	sqlstr := "select id, title, content from posts"
	rows, _ := db.Query(sqlstr)
	var id, title, content string
	for rows.Next(){
		_ = rows.Scan(&id, &title, &content)
		if title == "''"{
			title = ""
		}
		dict[id] = title + content
	}
	return dict
}
func GetUserText()(dict map[string]string){
	dict = make(map[string]string)
	sqlstr := "select id, name from user"
	rows, _ := db.Query(sqlstr)
	var id, name string
	for rows.Next(){
		_ = rows.Scan(&id, &name)
		dict[id] = name
	}
	return dict
}
func GetConditionText()(dict map[string]string){
	dict = make(map[string]string)
	sqlstr := "select * from condition222"
	rows, _ := db.Query(sqlstr)
	var id, name, introdution, symptom, treatment string
	for rows.Next(){
		err := rows.Scan(&id, &name, &introdution, &symptom, &treatment)
		check_err(err)
		dict[id] = name + introdution + symptom + treatment
	}
	return dict
}
func QueryQuestion(id string)(ret map[string]string){
	ret = make(map[string]string)
	sqlstr := "select title, content, view, likes, collects, html from questions where id=?"
	var title, content, view, likes, collects, html string
	err := db.QueryRow(sqlstr, id).Scan(&title, &content, &view, &likes, &collects, &html)
	check_err(err)
	ret["title"] = title
	ret["content"] = content
	ret["view"] = view
	ret["likes"] = likes
	ret["collects"] = collects
	ret["html"] = html
	return ret
}
func QueryBlog(id string)(ret map[string]string){
	ret = make(map[string]string)
	sqlstr := "select title, content, view, likes, collects, html from posts where id=?"
	var title, content, view, likes, collects, html string
	err := db.QueryRow(sqlstr, id).Scan(&title, &content, &view, &likes, &collects, &html)
	check_err(err)
	ret["title"] = title
	ret["content"] = content
	ret["view"] = view
	ret["likes"] = likes
	ret["collects"] = collects
	ret["html"] = html
	return ret
}
func QueryTreatment(id string)(ret []string){
	if db == nil{
		fmt.Println("error")
		return
	}
	sqlstr := "select name, introduction, uses, sideeffect from treatment where id=?"
	var name string
	var introduction string
	var uses string
	var sideeffects string
	err := db.QueryRow(sqlstr, id).Scan(&name, &introduction, &uses, &sideeffects)
	if (err != nil){
		fmt.Println(err)
	}
	ret = []string{name, introduction, uses, sideeffects}
	return ret
}

func DeleteTable(tablename string){
	sqlstr := "drop table if exists " + tablename
	_, err := db.Exec(sqlstr)
	if err != nil{
		fmt.Println(err)
	}
}
func DeleteUser(id string){
	sqlstr := "delete from user where id = ?"
	_, err := db.Exec(sqlstr, id)
	check_err(err)
	DeleteTable(id+"questioncollect")
	DeleteTable(id+"answercollect")
	DeleteTable(id+"postvisit")
	DeleteTable(id+"post")
	DeleteTable(id+"friend")
	DeleteTable(id+"postcollect")
	DeleteTable(id+"question")
	DeleteTable(id+"answer")
	DeleteTable(id+"answervisit")
	DeleteTable(id+"questionvisit")
	DeleteTable(id+"search")
}
func main() {

	err := InitDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	FindCookie("5577006791947779410")
	//info := make(map[string]string)
	//info["name"] = "admin"
	//info["email"] = "525802191@qq.com"
	//info["password"] = "password"
	//info["address"] = "Dongchuan Road Number 800"
	//id := NewUser(info)
	//log.Print(id)

	////test AddPost
	//info := map[string]string{"content": "你干嘛。。。哼哼。。。哎呦。。。"}
	//flag := AddPost(info)
	//if !flag {
	//	fmt.Println("err")
	//} else {
	//	fmt.Println("success")
	//}
	//
	////test query_history
	//friends := QueryOtherInfo("8", "friendid", "friend")
	//for _, value := range friends {
	//	fmt.Println(value)
	//}
	//return
	//
	////test QueryUser
	//query := []string{"id", "name", "email", "password", "address", "birthday", "introduction",
	//	"photo", "sex", "visit_history_num", "search_history_num", "article_num"}
	//ret := QueryUser("name", "xie", query)
	//for _, value := range ret {
	//	fmt.Println(value)
	//}
	//
	////test new_user
	//info := []string{"liu", "234@123.com", "123456", "adress", "birthday", "unknown", "unknown", "unknown", "0", "0", "0"}
	//id := new_user(info)
	//fmt.Println(id)
	//
	////test modify_user
	//m := map[string]string{
	//	"name": "zhangsan",
	//	"sex":  "femal",
	//}
	//fmt.Print(ModifyUser(8, m))
	//
	////test AddQuestion
	//info := map[string]string{"content": "你干嘛。。。哼哼。。。哎呦。。。", "answererid": "9"}
	//flag := AddAnswer(info)
	//if !flag {
	//	fmt.Println("err")
	//} else {
	//	fmt.Println("success")
	//}
	//
	////test query_answer_to_question
	//ret := QueryAnswersToAnswer("-1")
	//for ans := range ret {
	//	fmt.Println(ret[ans].content)
	//}
}
