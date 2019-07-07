package main

type File []byte

func (f File)Create(name, size, types, time, path, md5 []byte) {
	FILE_NAME_DB.Put(f, name, nil)
	FILE_PATH_DB.Put(f, path, nil)
	// 文件并不频繁修改, 数据可以合并
}
func (f File)Download() {
	FILE_PATH_DB.Put()
	订购者
}
func (f File)Delete() {
}
