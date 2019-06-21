package main

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	PROJECT_NAME_DB *leveldb.DB
	PROJECT_MARK_DB *leveldb.DB
	PROJECT_LIST_DB *leveldb.DB // 项目列表 所有
	// 为项目打TAG?
	// 如何计算项目关注度与活跃度两项或更多项指标?
	// 主页显示什么, 分P如何显示,分P的矩阵关系表
	// 
)

type Project struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Time    int64    `json:"time"`
	Item    string   `json:"item"`
	Task    []string `json:"task"`
	Mark    []string `json:"mark"`
	Master  []string `json:"master"`
	Partner []string `json:"partner"`
}

func init() {
	var err error
	PROJECT_NAME_DB, err = leveldb.OpenFile("../data/peoject_name", nil)
	if err != nil {
		panic("PROJECT_NAME_DB INIT ERROR")
	}
	PROJECT_MARK_DB, err = leveldb.OpenFile("../data/peoject_mark", nil)
	if err != nil {
		panic("PROJECT_NAME_DB INIT ERROR")
	}
}

// 底层操作, 逻辑选择
// 删除一个project, project下有 task item mark repy, 关联 user
// 删除一个task, task下有repy, 关联 user
// 即删除每个对象都操作关联的下级, 因为project并不是独立档, 而是由其他档拼接而成

//func (p Project)Delete() error {
	// 包含的
	//Mark project的存在, 如何得知新增了哪些呢? 如何新增呢..
	//delete("project", "id")
	//get project_mark // ["32","33","34"]
	// 判断是否删除时的权限表呢,,  也是一种内在关联,, 当角色要删除一个对象, 不必读取对象本身, 而去权限表验证, 前端角色又是怎么知道自己有没有权限呢
	// 关联的
	//USER_PROJECT_DB.Delete([]byte(p.Id), nil)
//}

func (p *Project)GetName()([]byte, error){
	return PROJECT_NAME_DB.Get([]byte(p.Id), nil)
}
func (p *Project)GetMark()([]byte, error){
	return PROJECT_MARK_DB.Get([]byte(p.Id), nil)
}

func (p *Project)SetName(){}
func (p *Project)SetMark(){}

func (p *Project)Delete() error {
	id := []byte(p.Id)
	if err := PROJECT_NAME_DB.Delete(id, nil); err != nil {
		return err
	}
	if err := PROJECT_MARK_DB.Delete(id, nil); err != nil {
		return err
	}
	//if err := PROJECT_ID_DB.Delete(id, nil); err != nil {
	//	return err
	//}
	return nil
	// 删评论 由于已经验证了对主贴的权限, 不必管跟贴权限
}
