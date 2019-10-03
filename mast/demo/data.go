// 静态 HTML CSS LESS JS 是存储在 DB 中
// 软件的生命周期, 启动即激活, 恢复初始只需删库(跑路)
// SERVER 的增减也是动态的, 可热拆卸
// 管理身份可操作增加对象, 对象具有属性
// 操作方法却不可以增加..? 即 MOD
// 关于数据挂钩, 读取 主题 时, 附加的标签 MARK
// 重复动作使用模板(管道), 动态动作使用映射
// 初始的通道是重复动作, 供应

// 对数据的存储
// 对数据的 CURD

package main

import (
	"encoding/binary"
	"log"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	DB map[string]*leveldb.DB // databases
	CA map[string]*sync.Map   // cache
	CH map[string]chan int64  // chan
)

func init() {
	var err error

	DB = make(map[string]*leveldb.DB)
	CA = make(map[string]*sync.Map)
	CH = make(map[string]chan int64)

	// init databases
	dbList := []string{"account", "password", "salt", "token", "counter", "user", "project"}
	for _, name := range dbList {
		DB[name], err = leveldb.OpenFile("data/"+name, nil)
		if err != nil {
			panic("OpenFile error")
		}
	}

	// init cache
	//caList := []string{"token"}
	//for _, name := range caList {
	//	CA[name] = make(map[string])
	//}

	// init counter
	chList := []string{"user", "project"}
	for _, name := range chList {
		CH[name] = make(chan int64)
		go func(name string, c chan int64) {
			var err error
			var sum int64

			buf := make([]byte, 8)

			data, err := DB["counter"].Get([]byte(name), nil)
			if err != nil {
				binary.BigEndian.PutUint64(buf, uint64(sum))
				err = DB["counter"].Put([]byte(name), buf, nil)
				log.Println("计数器初始化", name)
				if err != nil {
					panic("DB Put error")
				}
			} else {
				sum = int64(binary.BigEndian.Uint64(data))
				log.Println("计数器", name, sum)
			}

			for {
				sum++
				c <- sum
				binary.BigEndian.PutUint64(buf, uint64(sum))
				err = DB["counter"].Put([]byte(name), buf, nil)
				if err != nil {
					panic("counter error")
				}
			}
		}(name, CH[name])
	}

	// init Object
	//obList := []string{"user", "project"}
	//for _, name := range obList {}
}
