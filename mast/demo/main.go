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
	// 角色在登录后才允许发起 ws 连接, 否则只能临时会话
	http.HandleFunc("/project", func(w http.ResponseWriter, r *http.Request) {
	})
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		var data interface{}

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
		case "create":

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

