package main

import "math/rand"

type Object struct {
	Name string
	ID   string
}

func (o Object) Master() []string {
	//value, err := DB[o.Name].Get([]byte(o.ID), nil)
	//if err != nil {
	//	return nil, err
	//}
	return []string{"323", "12", "666"}
}

type User struct {
	ID    string
	Token string
}

func (u User) Create(o Object, data string) error {
	// 检查 o.Name 是否存在, 从server层传递过来必然要存在 所以不必在这层检查
	// 创建应当此处生成 ID, 并返回给外部
	// ID 对象存在哪些绝对赋值项? 理论并没有
	return DB[o.Name].Put([]byte(o.ID), []byte(data), nil)
}

// 当需要创建物品时, 如主贴
type (
	Project struct {
		ID     string
		Master string
		Time   int64
	}
	Reply struct {
		ID     string
		Master string
		Time   int64
	}
)

// id
// time
// 有多个数据集的对象没有time
// 数据集分别修改?

// project
// json{name, info, time, master}
// project{name,}

// Delete Object
func (u User) Delete(o Object) error {
	var err error
	list := o.Master()
	for _, id := range list {
		// is []byte !? no..
		if id == o.ID {
			return DB[o.Name].Delete([]byte(o.ID), nil)
		}
	}
	return err
}

// Rewrite user 只是无权修改其他角色的, 并非没有这个选项
func (u User) Rewrite(o Object, data string) error {
	//if u.ID != o.ID {
	//	return nil
	//}
	return DB[o.Name].Put([]byte(o.ID), []byte(data), nil)
}

// Master : The owner of the user is himself
func (u User) Master() string {
	return u.ID
}

// Signout is user
func (u *User) Signout() bool {
	token, ok := CA["token"].Load(u.ID)
	if !ok || token != u.Token {
		return false
	}
	CA["token"].Delete(u.ID)
	return true
}

// Signin is user
func (u *User) Signin(account, password string) bool {
	// 检测不准为空, 另验证码
	if account == "" || password == "" {
		return false
	}
	// 验证码 code
	// 读取 account
	idByte, err := DB["account"].Get([]byte(account), nil)
	if err != nil {
		return false
	}
	// 分割 data 为 pw 和 salt 和 id?
	// 暂时不作分割, 等待更新为更安全的密钥存储方式
	pwByte, err := DB["password"].Get(idByte, nil)
	if err != nil {
		panic("db get user salt error")
	}
	//saltByte, err := DB["salt"].Get(idByte, nil)
	//if err != nil {
	//	panic("db get user salt error")
	//}
	if password != string(pwByte[:]) {
		return false
	}

	id := string(idByte[:])

	//if password != string(pw[:]) {
	//	return "err is pw", nil
	//}

	// 设置 token, 通过了密码验证
	letters := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	token := string(b)
	CA["token"].Store(id, token)

	u.ID = id
	u.Token = token
	// 返回 token true
	return true
}

// 生成伪随机字符串, 用于token
func randSeq(n int) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
