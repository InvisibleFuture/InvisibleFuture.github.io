package main

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type Mark struct {
	Object string
	Id     string
	User   string
}

var(
	PROJECT_MARK *leveldb.DB
)

func init() {
	// 初始化所有项? 若是关停某项, 此处依旧会初始化
	// 每项都是独立的, 不使用数据传导
	// 应当在同一页初始化全部数据
	MARK_PROJECT_USER_DB, err := leveldb.OpenFile("../data/mark_project_user")
	if err != nil {
		panic("MARK_PROJECT_USER_DB INIT ERROR")
	}
	defer MARK_PROJECT_USER_DB.Close()
	MRRK_USER_PROJECT_DB, err := leveldb.OpenFile("../data/mark_user_project")
	if err != nil {
		panic("MARK_USER_PROJECT_DB INIT ERROR")
	}
	defer MARK_USER_PROJECT_DB.Close()
}

func (m *Mark)on() {
	// 查询类型 id
	// 确认是否存在
	// 双向标记
	// 回执
}



