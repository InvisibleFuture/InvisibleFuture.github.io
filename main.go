package main
import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
)

func main(){
	fmt.Println("srart!")
	http.HandleFunc("/api", api)
	http.HandleFunc("/user", user)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndSere: ", err)
	}
}

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
}

func apis(w http.ResponseWeiter, r *http.Request) {
	//mo
}

func api(w http.ResponseWriter, r *http.Request) {
	// all api > user
	var u User
	u.Id = r.FormValue("id")
	u.Token = r.FormValue("token")
	//if ok := u.Authentication(); !ok {
	//	return //token not..  token time..
	//}

	profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	aim := r.FormValue("aim")
	ids := []string{"4232","323"}
	switch aim {
	case "delete": u.Delete(ids)
	default: fmt.Println("Did not provide the necessary parameters")
	}
}

/*
	# 项目
	多方案 > 可推进方案 > 最优方案
	热数据 在线的user
*/
