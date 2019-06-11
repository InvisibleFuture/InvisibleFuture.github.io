package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
)

type Profile struct {
	Name    string
	Hobbies []string
}

type item struct {
	Name    string
	Time    string
	Master  []string
	Partner []string
	Task    []string
}

type task struct {
	Name string
	Time string
	say []string // 关联数据在内存中, 不入库
}

type say struct {
	Master  string
	Content string // Attach file in reply
	Attach  []string // a7a873eyhq78cyhq3rkjrh3, Echo Ban nesting
	Time    string
}

type attach struct {
	Name string
	Path string
	Size string
	Hash string
	Time string
}

type Project struct {
	Name    string   `json:"name"`
	Time    string   `json:"time"`
	Item    string   `json:"item"`
	Mark    []string `json:"mark"`
	Master  []string `json:"master"`
	Partner []string `json:"partner"`
}

var (
	ITEM_DB *leveldb.DB
)

func main(){
	fmt.Println("srart!")
	http.HandleFunc("/api", api)
	http.HandleFunc("/user", user)
	http.HandleFunc("/project", project)
	http.HandleFunc("/login", sign_in)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndSere: ", err)
	}
}

func mark(w http.ResponseWriter, r *http.Request){
	//mark不是组件自带功能, 作为插入件存在
	//a是否mark了此文档, 取决于此文档被谁mark过的动态数据
	//此文档都是被谁mark过, 直接附加于主文档之上
	//a mark过哪些文档, a的附属mark档
	//故 双向链表在此处有两个
	//组件的mark功能引入自基础件"库"中, 
	// 基础件 db 库
	// 读取此组件的mark条目, mark不会触发time更新
	// mark单独读写不必关联大幅数据解析
	// 组件读取功能是不解析的
	// 所有组件有必要分离
}

func project_item(w http.ResponseWriter, r *http.Request){
	//p1 := item3232
	//p2 := item3333
	//data := []item{}
}

func sign_in(w http.ResponseWriter, r *http.Request) {
	// post !
	var a Account
	a.Account = "Last" //r.FormValue("account")
	a.Password = "dedeff" //r.FormValue("password")

	password, err := a.GetPassword()
	//目标键值不存在时会返回err, 作为用户输入数据不应中断主进程
	//panic("SIGN_IN GET_PASSWORD ERROR")
	if err != nil || a.Password != string(password[:]) {
		cookieK := http.Cookie{Name:"id", Path: "/", MaxAge: -1}
		cookieV := http.Cookie{Name:"token", Path:"/", MaxAge: -1}
		http.SetCookie(w, &cookieK)
		http.SetCookie(w, &cookieV)
		w.Write([]byte("sign in err !"))
		return
	}
	id, err := a.GetId()
	if err != nil {
		panic("SIGN_IN GET_ID ERROR 存在账户却不存在角色")
	}

	var u User
	u.Id = string(id[:])
	u.Token = randSeq(32)
	u.SetToken()

	cookieK := http.Cookie{Name:"id",Value:u.Id,Path:"/", MaxAge:86400}
	cookieV := http.Cookie{Name:"token",Value:u.Token,Path:"/",MaxAge:86400}
	http.SetCookie(w, &cookieK)
	http.SetCookie(w, &cookieV)
	w.Write([]byte("sign in ok !"))
}
func sign_up(w http.ResponseWriter, r *http.Request) {
	//account := "Last"
	//password := "dedeff"
	//mail := "last@gmail.com"

	//name := "Last"

	//verification := "45648"
	//pw = a.GetPassword()
}

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func project(w http.ResponseWriter, r *http.Request) {
	var u User
	//通过url参数或post数据获取身份认证信息
	//u.Id = r.FormValue("id")
	//u.Token = r.FormValue("token")
	//通过cookies获取身份认证信息// 禁止get方式操作数据操作只允许post 且附上token码
	id, err := r.Cookie("id")
	if err != nil {
		return // 未登录
	}
	token, err := r.Cookie("token")
	if err != nil {
		return // 未登录
	}
	u.Id = id.Value
	u.Token = token.Value
	fmt.Println(u.Id)
	fmt.Println(u.Token)
	if ok := u.Authentication(); !ok {
		data := Profile{"Alex", []string{"snowboarding", "programming"}}
		js, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return
	}
	data := Project{
		"is project new",
		"1984/06/12",
		"19",
		[]string{"434433","32321","32323"},
		[]string{"a64d5a","d46a4d","d46w2a2"},
		[]string{"d46ad46a4d5","d4w6a4d65"},
	}
	//profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	//js, err := json.Marshal(data)
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func Item_Tasks(w http.ResponseWriter, r *http.Request) {
	// xieru
	// 读取身份
	// 读取数据
	// 验证权限
	// 入库数据
	// 返回信息
}
func api(w http.ResponseWriter, r *http.Request) {
	// 取得身份, 由于普通访客是不允许交互的 不再验证
	var u User
	u.Id = r.FormValue("id")
	u.Token = r.FormValue("token")
	//if ok := u.Authentication(); !ok {
	//	return //token not..  token time.. 验证失败或未登录都不允许继续
	//}

	// 读取目标数据
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 验证权限
	// 执行操作
	// 返回数据
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func user(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	var u User
	// 任何交互操作都需要提供身份证明与目标目的
	// 这个页面是 http 交互层, 必须是 post 方法?
	u.Id = r.FormValue("id")
	u.Token = r.FormValue("token")
	fmt.Println("u.Token", r.FormValue("token"))
	fmt.Println(u)
	//aim := r.FormValue("aim")
	//ids := []string{"4232","323"}
	//switch aim {
	//	case "delete": u.Delete(ids)
	//	default: fmt.Println("Did not provide the necessary parameters")
	//}
}

