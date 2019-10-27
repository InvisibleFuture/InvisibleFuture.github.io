package main

import (
	"log"
	"strconv"
)

type Account []byte

func (a Account)Create(password []byte) ([]byte, bool) {
	id, err := ACCOUNT_ID_DB.Get(a, nil)
	if err == nil {
		return id, false
	}

	err = ACCOUNT_PW_DB.Put(a, password, nil)
	if err != nil { panic(err) }

	autoid := <-AUTOID_USER_CH
	id = []byte(strconv.FormatInt(autoid, 10))
	err = ACCOUNT_ID_DB.Put(a, id, nil)
	if err != nil { panic(err) }
	return id, true
}

func (a Account)Signin(password string) ([]byte, bool) {
	var err error
	var pw, id []byte

	pw, err = ACCOUNT_PW_DB.Get(a, nil)
	if err != nil || password != string(pw[:]) {
		log.Println(string(pw[:]))
		log.Println(string(password))
		return id, false
	}

	id, err = ACCOUNT_ID_DB.Get(a, nil)
	if err != nil {
		panic("已匹配账号却未找到ID")
	}

	return id, true
}