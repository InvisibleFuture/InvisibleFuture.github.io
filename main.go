package main
import (
	"fmt"
	"net/http"
	"log"
)

func main(){
	fmt.Println("srart!")

	http.HandleFunc("/user", user)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndSere: ", err)
	}
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
