package main

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	IMAGE_NAME_DB *leveldb.DB
	IMAGE_PATH_DB *levledb.DB
)

type Image []byte

func (i Image)Create(name, item []byte) {
	// 图像还包括缩略图
	// 图像不需要时间?
	// 若是绘画站则还需要签名和时间
	IMAGE_NAME_DB.Put(f, name, nil)
	IMAGE_PATH_DB.Put(f, name, nil)
}

func (i Image)Delete() {
	IMAGE_NAME_DB.Delete(i, nil)
	IMAGE_PATH_DB.Delete(i, nil)
}

func init() {
	var err error
	IMAGE_NAME_DB, err = leveldb.OpenFile("../data/image_name", nil)
	if err != nil { panic(err) }
	IMAGE_PATH_DB, err = leveldb.OpenFile("../data/image_name", nil)
	if err != nil { panic(err) }
}


