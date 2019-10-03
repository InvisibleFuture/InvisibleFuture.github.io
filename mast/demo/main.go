package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("start w~")
	http.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {
		num := <-CH["user"]
		log.Println(num)
	})
	http.HandleFunc("/project", func(w http.ResponseWriter, r *http.Request) {
		// get list44684
		// get p46465
		// post repyt :p6466,repnew
		//
	})
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		var data interface{}

		// 登录需要的信息, 若是多次操作需要输入验证码
		//[]string{"account", "password", "code"}

		// 注册需要的信息, 限制注册没有意义
		list := []string{"account", "password", "code"}
		for _, key := range list {
			value := r.FormValue(key)
			log.Println(value)
		}

		var u User

		exe := r.FormValue("exe")
		switch exe {
		case "signin":
			account := r.FormValue("account")
			password := r.FormValue("password")

			if u.Signin(account, password) {
				data = struct{ Code, Uid, Token string }{Uid: u.ID, Token: u.Token}
			} else {
				data = struct{ Code, Info string }{Code: "403", Info: "用户名不存在或密码不匹配"}
			}
		case "sigout":
			u.ID = r.FormValue("uid")
			u.Token = r.FormValue("token")
			if u.Signout() {
				data = struct{ Code, Info string }{Code: "200", Info: "ok"}
			} else {
				data = struct{ Code, Info string }{Code: "403", Info: "err"}
			}
		//case "rename":
		case "create":
			//account := r.FormValue("account")
			//password := r.FormValue("password")
			//if u.Create(account, password) {
			//	data = struct{ Code, Uid, Token string }{Uid: u.ID, Token: u.Token}
			//} else {
			//	data = struct{ Code, Info string }{Code: "403", Info: "用户名不存在或密码不匹配"}
			//}
		//	account := r.FormValue("account")
		//	password := r.FormValue("password")
		//	log.Println(account, password)
		// 检测不准为空, 另验证码
		// 检测是否被占用
		//var o Object
		//var u User
		//err := u.Create(Object{Name: "UserAccount", ID: newid})
		//err := u.Create(Object{Name: "UserPassword", ID: newid})
		default:
			// 此处直接拒绝非法指令 并非404
		}

		js, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			panic("echo json error")
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte(js))
	})

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
	time.Sleep(time.Second * 1)

	// Store the token in the database

	// End all database connections
	for name := range DB {
		DB[name].Close()
	}

	log.Println("done.")
}

// 数据, 前端 是底层抽象
// 上层抽象调动底层抽象, 而非一个列表直接运转
// 核心 控制器
// 生 方法? soul
// 物 生的物态 user
// 物 属性数据
// 包含重新编译前端文件吗? 前端是载入到内存的, 所以必然, 之后监听文件更改?
// 已经初始化, 则开始载入基本状态数据
// 状态预备完毕, 开启服务端口
// 流量涌入, 周期状态调整
// 插件是一种理论抽象 而非外部增补, 插件基于基因库, 无需代码

// 生成伪随机字符串, 用于token
//func randSeq(n int) string {
//	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
//	b := make([]rune, n)
//	for i := range b {
//		b[i] = letters[rand.Intn(len(letters))]
//	}
//	return string(b)
//}
