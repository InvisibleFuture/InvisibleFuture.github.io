package mian

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	ACCOUNT_PW_DB     *leveldb.DB
	ACCOUNT_ID_DB     *leveldb.DB

	USER_NAME_DB      *leveldb.DB

	TAG_NAME_DB       *leveldb.DB

	PROJECT_NAME_DB   *leveldb.DB
	PROJECT_MASTER_DB *leveldb.DB
)

func init() {

}
