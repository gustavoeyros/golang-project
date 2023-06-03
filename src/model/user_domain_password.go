package model

import (
	"crypto/md5"
	"encoding/hex"
)

func (ud *userDomain) EncryptPassowrd() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
