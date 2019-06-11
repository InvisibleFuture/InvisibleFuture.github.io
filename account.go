package main

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	ACCOUNT_ID_DB *leveldb.DB
	ACCOUNT_PW_DB *leveldb.DB
)

type Account struct {
	Account  string
	Password string
}

func init() {
	var err error
	ACCOUNT_PW_DB, err = leveldb.OpenFile("../data/account_pw", nil)
	if err != nil {
		panic("ACCOUNT_PW_DB INIT ERROR")
	}
	ACCOUNT_ID_DB, err = leveldb.OpenFile("../data/account_id", nil)
	if err != nil {
		panic("ACCOUNT_ID_DB INIT ERROR")
	}
	ACCOUNT_PW_DB.Put([]byte("Last"), []byte("dedeff"), nil)
	ACCOUNT_ID_DB.Put([]byte("Last"), []byte("233"), nil)
}

func (a *Account)GetPassword() ([]byte, error) {
	return ACCOUNT_PW_DB.Get([]byte(a.Account), nil)
}
func (a *Account)GetId() ([]byte, error) {
	return ACCOUNT_ID_DB.Get([]byte(a.Account), nil)
}



/*
func (a *Account)CK() (string, bool) {
	password, err := ACCOUNT_PW_DB.Get([]byte(a.Account), nil)
	// err 与 判断条件不是同一级别的错误, 不应同时判断
	if err != nil || password != a.Password {
		return "", false
	}
	id, err := ACCOUNT_ID_DB.Get([]byte(a.Account), nil)
	if err != nil {
		return "", false
	}
	return id, true
}
*/
