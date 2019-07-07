package main

type Account []byte

func (account Account)Create(password []byte) {
	var err error
	err = ACCOUNT_PW_DB.Put(account, password, nil)
	if err != nil { panic(err) }
	err = ACCOUNT_ID_DB.Put(account, id, nil)
	if err != nil { panic(err) }
}

/**
type Account struct {
	Account  string
	Password string
}

func (a *Account)Create(id string) {
	var err error
	err = ACCOUNT_PW_DB.Put([]byte(a.Account), []byte(a.Password), nil)
	if err != nil { panic("ACCOUNT_PW_DB CREATE ERROR") }
	err = ACCOUNT_ID_DB.Put([]byte(a.Account), []byte(id), nil)
	if err != nil { panic("ACCOUNT_ID_DB CREATE ERROR") }
}

func (a Account)GetPassword() ([]byte, error) {
	return ACCOUNT_PW_DB.Get([]byte(a.Account), nil)
}
func (a Account)GetId() ([]byte, error) {
	return ACCOUNT_ID_DB.Get([]byte(a.Account), nil)
}
**/

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
