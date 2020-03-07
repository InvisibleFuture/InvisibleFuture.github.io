package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func project(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Reply(w, 200, "Get访问某个项目")
		return
	}

	// 取得所需数据并验证格式正确性
	uid := r.FormValue("uid")
	token := r.FormValue("token")
	if uid == "" || token == "" {
		Reply(w, 403, "非法参数")
		return
	}

	// 验证身份
	val, ok := WEB_TOKEN_MAP.Load(uid)
	if !ok || val != token {
		Reply(w, 401, "验证失败 请重新登录")
		return
	}

	data := make(map[string]string)
	data["info"] = "asdadas"
	Echo(w, data)
}
func main() {
	http.HandleFunc("/", index)

	//http.HandleFunc("/user", user)
	http.HandleFunc("/home", home)
	//http.HandleFunc("/project", project)
	http.HandleFunc("/project_create", project_create)
	http.HandleFunc("/signin", sign_in)
	http.HandleFunc("/signup", sign_up)
	http.HandleFunc("/signrw", sign_rw)
	http.HandleFunc("/plist", list_project)
	http.HandleFunc("/delete", Delete)

	http.HandleFunc("/project", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			Reply(w, 200, "Get访问某个项目")
			return
		}
		// 取得所需数据并验证格式正确性
		uid := r.FormValue("uid")
		token := r.FormValue("token")
		log.Println(uid, token)
		//if uid == "" || token == "" {
		//	Reply(w, 403, "非法参数")
		//	return
		//}

		// 认证身份
		val, ok := WEB_TOKEN_MAP.Load(uid)
		if !ok || val != token {
			Reply(w, 401, "验证失败 请重新登录")
			return
		}
		// 操作目标 / 操作权限
		// 选择操作目标后才知道需要哪些数据
		switch action {
		case "create":
			id := r.FormValue("id")
			// 读取确认权限? 不是web层的
			user.CreateProject(id)
		case "delete":
		case "rewrite":
		default:
		}
		
		// 回执
	})

	s := &http.Server{
		Addr:           ":8080",
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

func list_project(w http.ResponseWriter, r *http.Request) {
	list := List{
		Id: r.FormValue("id"),
	}
	if list.Id == "" {
		Reply(w, 404, "非法输入")
		return
	}
	data, err := list.GetProjects()
	if err != nil {
		Reply(w, 404, "页面不存在")
		return
	}
	log.Println(data)
	Echo(w, data)
}
func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Reply(w, 404, "目标资源不存在")
		return
	}
	file, err := os.Open("./html/index.html")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(doc)
}

type HomeData struct {
	Projects []ListProject
}

func home(w http.ResponseWriter, r *http.Request) {
	// 输入并验证数据 执行 反馈

	l := List{Id: "1"}
	p, err := l.GetProjects()
	if err != nil {
	}

	Echo(w, HomeData{Projects: p})
}
func Delete(w http.ResponseWriter, r *http.Request) {
	// 所有 WEB 操作都是映射到角色对象上
	if r.Method != "Post" {
		Reply(w, 403, "必须使用POST方法提交命令")
		return
	}

	//取得数据
	id := r.FormValue("id")
	target := r.FormValue("target")
	uid := r.FormValue("uid")
	token := r.FormValue("token")

	if uid == "" || token == "" || target == "" || id == "" {
		Reply(w, 403, "非法参数")
		return
	}

	/**
	var o Object
	switch target {
	case "project": o = Project{ Id: id }
	case "user":    o = User{ Id: id }
	default:        return // 非法参数拒绝执行
	}

	//o := Object(object)
	if ok := o.Delete(); !ok {
		Reply(w, 200, "拒绝")
		return
	}
	**/

	Reply(w, 403, "成功")
}

// 注册, 找回, 登录, 续约, 退出
func sign_in(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		Reply(w, 403, "POST 方式登录必须")
		return
	}

	account := r.FormValue("account")
	password := r.FormValue("password")

	if l := len(account); l < 4 || l > 32 {
		Reply(w, 403, "账号格式不符")
		return
	}

	if l := len(password); l != 32 {
		Reply(w, 403, "未加密的密钥")
		return
	}

	id, ok := Account([]byte(account)).Signin(password)
	if !ok {
		Reply(w, 403, "帐号密码不匹配")
		return
	}

	token := rwtoken(string(id))

	Reply(w, 200, token) //+id

	/**
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
	**/
}
func sign_up(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		Reply(w, 403, "POST 方式提交必须")
		return
	}

	account := r.FormValue("account")
	password := r.FormValue("password")

	if l := len(account); l < 4 || l > 32 {
		Reply(w, 403, "账号格式不符")
		return
	}

	if l := len(password); l != 32 {
		Reply(w, 403, "未加密的密钥")
		return
	}

	id_byte, ok := Account([]byte(account)).Create([]byte(password))
	if !ok {
		Reply(w, 403, "账号已占用")
		return
	}

	id := string(id_byte)
	token := rwtoken(id)

	m := make(map[string]string)
	m["id"] = id
	m["token"] = token
	log.Println(m)
	Echo(w, m)
}
func sign_rw(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		Reply(w, 403, "POST 方式提交必须")
		return
	}

	// 读取数据
	id := r.FormValue("id")
	token := r.FormValue("token")
	if id == "" || token == "" {
		Reply(w, 401, "未登录")
		return
	}

	// 验证身份
	val, ok := WEB_TOKEN_MAP.Load(id)
	if !ok || val != token {
		Reply(w, 401, "验证失败 请重新登录")
		return
	}

	// 重置 token
	//u := User([]byte(id))
	//token = u.SetToken()
	token = rwtoken(id)

	// 回执
	m := make(map[string]interface{})
	m["id"] = id
	m["token"] = token
	Echo(w, m)

	//var u User
	//if ok := Identity(r, &u); !ok {
	//	cookieA := http.Cookie{Name:"id", Path:"/", MaxAge:-1}
	//	http.SetCookie(w, &cookieA)
	//	cookieB := http.Cookie{Name:"token", Path:"/", MaxAge:-1}
	//	http.SetCookie(w, &cookieB)
	//	Reply(w, 401, "续约失败, Token 已过期, 请重新登录")
	//	return
	//}
	//fmt.Println(u)
	//Reply(w, 200,"续约成功, Token已更新")
}

func Reply(w http.ResponseWriter, code int, info string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write([]byte(info))
}
func Echo(w http.ResponseWriter, data interface{}) {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}

func project_create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method != "POST" {
		Reply(w, 403, "POST 方式提交必须")
		return
	}

	uid := r.FormValue("uid")
	token := r.FormValue("token")

	val, ok := WEB_TOKEN_MAP.Load(uid)
	if !ok || val != token {
		Reply(w, 401, "验证失败 请重新登录")
		return
	}

	/**
	name := r.FormValue("name")
	if name == "" {
		Reply(w, 304, "name 参数不符长度")
	}
	var p Project
	id := <-AUTOID_PROJECT_CH
	p.Id = strconv.FormatInt(id, 10)
	p.Name = r.FormValue("name")
	p.Time = time.Now().Unix()
	p.Master = []string{uid}
	fmt.Println(p)

	p.Create(name, p.Id)
	u.ProjectPush(p.Id)
	**/

	Reply(w, 200, "返回结果 pid")
	// 追加 PID 到 USER表
}
