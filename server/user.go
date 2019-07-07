package main

type User []byte

func (u User)Create(name, item []byte) {
	// 初始化一个角色
	USER_NAME_DB.Put(u, name, nil)
	USER_TIME_DB.Put(u, name, nil)

	USER_ITEM_DB.Put(u, name, nil)
	USER_MARK_DB.put(u, name, nil)
}

func (u User)Rename(name []byte) {
	USER_NAME_DB.Put(u, name, nil)
}



