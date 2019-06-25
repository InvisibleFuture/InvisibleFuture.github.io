package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"strconv"
	"strings"
	"io/ioutil"
	"os/signal"
	"syscall"
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

var (
	ITEM_DB *leveldb.DB
	INDEX []byte
)
func init(){
	file, err := os.Open("index.html")
	if err != nil {
		log.Fatal(err)
	}
	INDEX, err = ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
}
func main(){
	fmt.Println("srart!")
	http.HandleFunc("/", index)
	http.HandleFunc("/api", api)
	//http.HandleFunc("/user", user)
	http.HandleFunc("/project", project)
	http.HandleFunc("/project_create", project_create)
	http.HandleFunc("/signin", sign_in)
	http.HandleFunc("/signup", sign_up)
	http.HandleFunc("/signrw", sign_rw)

	s := &http.Server{
		Addr:           ":80",
		Handler:        http.DefaultServeMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Println(s.ListenAndServe())
		log.Println("server shutdown")
	}()

	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	// Stop the service gracefully.
	log.Println(s.Shutdown(nil))

	// Wait gorotine print shutdown message
	//time.Sleep(time.Second * 1)
	// 结束所有数据库连接
	log.Println("done.")
}
func index(w http.ResponseWriter, r *http.Request){
	file, err := os.Open("index.html")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(doc)
	//w.Write(INDEX)
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

type userinfo struct{
	Id   string
	Name string
}

/*
func user(w http.ResponseWriter, r *http.Request) {
	var u User

	// 未登录的在前端不会连接到这里, cookie提供是必需的

	id, err := r.Cookie("id")
	if err != nil {
		return false
	}
	token, err := r.Cookie("token")
	if err != nil {
		return false
	}

	var u User
	u.Id = id.Value
	u.Token = token.Value
	if ok := u.Authentication(); !ok {
		return false
	}
	// 已经确定登录身份
	// 登录状态

	// USER ID CK
	// DATA IS
	// EK
	// RES
	ids := []string{"233","234","235"}
	var list []userinfo
	for _, id := range ids {
		name, _ := u.GetUser(id)
		user := userinfo{Id:id,Name: string(name)}
		list = append(list, user)
	}
	Echo(w, Msg{
		Code: "200",
		Info: "OK",
		User: list,
	})
}
*/

func sign_in(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

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
	account := Account{
		Account: r.FormValue("account"),
		Password: r.FormValue("password"),
	}

	fmt.Println(r.FormValue("account"))
	if r.Method != "POST" {
		Reply(w, 403, "POST 方式提交必须")
		return
	}

	if l := len(account.Account); l < 4 || l > 32 {
		Reply(w, 403, "账号格式不符")
		return
	}

	if l := len(account.Password); l != 32 {
		Reply(w, 403, "密钥格式不符")
		return
	}

	_, err := account.GetId()
	if err == nil {
		Reply(w, 403, "账号已占用")
		return
	}

	id := <-AUTOID_USER_CH
	user := User{ Id: strconv.FormatInt(id, 10), Token: randSeq(32) }

	account.Create(user.Id)
	user.SetToken()

	cookieK := http.Cookie{Name:"id",Value:user.Id,Path:"/", MaxAge:86400}
	cookieV := http.Cookie{Name:"token",Value:user.Token,Path:"/",MaxAge:86400}
	http.SetCookie(w, &cookieK)
	http.SetCookie(w, &cookieV)
	w.Write([]byte("sign in ok !"))
}

	// 邮件并不是安全的, 使用自己的安全验证方式
	// 特征码
	//signature := "voiajcfiojcoicjaioj39203920"
	// 左利手右利手 甜粽党咸粽党
	// 注册时间 注册地区 语言偏好 曾用name
	//verification := "45648"
	//pw = a.GetPassword()

	//检查名字占用
	//USER_NAME_DB, 高效检索库? 首字母 多字母
	// 生成一个ID, item task作为高频读写项目是否必要带id
	// 如不对项目进行单独的读取, 则不必要带id
	// 角色与项目是必要使用id的, 然后标签分类也要
	//id := NewId()
	//USER_NAME_DB.Put([]byte(id))

func sign_rw(w http.ResponseWriter, r *http.Request) {
	var u User
	if ok := Identity(r, &u); !ok {
		cookieA := http.Cookie{Name:"id", Path:"/", MaxAge:-1}
		http.SetCookie(w, &cookieA)
		cookieB := http.Cookie{Name:"token", Path:"/", MaxAge:-1}
		http.SetCookie(w, &cookieB)
		Reply(w, 401, "续约失败, Token 已过期, 请重新登录")
		return
	}
	fmt.Println(u)
	Reply(w, 200,"续约成功, Token已更新")
	/* 组成续约的像
	 * 契约最小时间, 契约最大时间
	 * 再生 token 写入token
	 */
}

func Reply(w http.ResponseWriter, code int, info string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(info))
}

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 验证身份
// 操作目标 > 权限验证
// 回执信息
func Echo(w http.ResponseWriter, data interface{}) {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func Identity(r *http.Request, u *User) bool {
	id, err := r.Cookie("id")
	if err != nil {
		return false
	}
	token, err := r.Cookie("token")
	if err != nil {
		return false
	}
	u.Id = id.Value
	u.Token = token.Value
	if ok := u.Authentication(); !ok {
		return false
	}
	return true
}
type Msg struct {
	Code string     `json:"code"`
	Info string     `json:"info"`
	User []userinfo `json:"user"`
}

func project(w http.ResponseWriter, r *http.Request) {
	var u User
	if ok := Identity(r, &u); !ok {
		Echo(w, Msg{Code:"401", Info:"请求要求用户的身份认证"})
		return
	}

	var p Project
	p.Id = "23432" //r.FormValue("p")
	mark, err := p.GetMark()
	if err != nil {
		Echo(w, Msg{Code:"404", Info:"不存在的项目"})
		return
	}
	p.Mark = strings.Fields(string(mark[:]))
	Echo(w, Project{
		Id:      p.Id,
		Name:    "is project new",
		Time:    0,
		Item:    "19",
		Task:    []string{},
		Mark:    []string{"434433","32321","32323"},
		Master:  []string{"a64d5a","d46a4d","d46w2a2"},
		Partner: []string{"d46ad46a4d5","d4w6a4d65"},
	})
}

func Chaos(w http.ResponseWriter, r *http.Request) {
	var u User
	if ok := Identity(r, &u); !ok {
		Echo(w, Msg{Code:"401", Info:"请登录"})
		return
	}
}

func project_create(w http.ResponseWriter, r *http.Request) {
	var u User
	if ok := Identity(r, &u); !ok {
		Echo(w, Msg{Code:"401", Info:"请求要求用户的身份认证"})
		return
	}
	token := r.FormValue("token")
	if token != u.Token {
		Echo(w, Msg{Code:"500", Info:"非法请求, Post必须附上Token"})
		return
	}
	var p Project
	p.Id = "233" //生成序列数
	p.Name = r.FormValue("name")
	p.Time = time.Now().Unix()
	p.Master = []string{u.Id}
	fmt.Println(p)

	// 验证输入的数据有效
	// 创建 PID, 建立P
	// 追加 PID 到 USER表
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
func userss(w http.ResponseWriter, r *http.Request){
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

