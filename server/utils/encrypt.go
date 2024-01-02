package utils

import (
	"Demo/global"
	"fmt"
	"golang.org/x/crypto/argon2"
)

func EncryptPassWord(passWord string) string {
	hashedPassword := argon2.IDKey([]byte(passWord), []byte(global.GLO_SALT), 1, 64*1024, 4, 32)
	return fmt.Sprintf("%x", hashedPassword)
}
