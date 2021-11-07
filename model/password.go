package model

import (
	"regexp"

	scrypt "github.com/elithrar/simple-scrypt"
)

var (
	numberReg      = regexp.MustCompile("\\d+")
	lowerLetterReg = regexp.MustCompile("[a-z]+")
	upperLetterReg = regexp.MustCompile("[A-Z]+")
)

func ValidPasswdStrength(str string) bool {
	if len([]rune(str)) < 8 {
		return false
	}
	result := numberReg.MatchString(str)
	if !result {
		return false
	}
	result = lowerLetterReg.MatchString(str)
	if !result {
		return false
	}
	return upperLetterReg.MatchString(str)
}

//EncryptedPassword 对密码进行加密
func EncryptedPassword(password string) (string, error) {
	hash, err := scrypt.GenerateFromPassword([]byte(password), scrypt.DefaultParams)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
