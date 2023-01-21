package util

import (
	"SeKai/internal/config"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"regexp"
)

// StandardizeSpaces 去除字符串中所有的空格和换行符/*
func StandardizeSpaces(s string) string {
	regex, _ := ReplaceStringByRegex(s, "\\s+", "")
	return regex
}

func ReplaceStringByRegex(str, rule, replace string) (string, error) {
	reg, err := regexp.Compile(rule)
	if reg == nil || err != nil {
		return "", errors.New("正则MustCompile错误:" + err.Error())
	}
	return reg.ReplaceAllString(str, replace), nil
}

func RandStr(strSize int) string {
	dictionary := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var bytes = make([]byte, strSize)
	_, _ = rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

func HashGoogleSecret(googleAuthSecret string) string {
	sha := sha256.New()
	sha.Write([]byte(googleAuthSecret + config.ApplicationConfig.GoogleAuth.Salt))
	// 只需要前16个字符
	hashedStringSecret := hex.EncodeToString(sha.Sum(nil))[0:16]
	return ReplaceAllNumber(hashedStringSecret)
}

func ReplaceAllNumber(s string) string {
	byteResult := []byte(s)
	for i := 0; i < len(s); i++ {
		if '0' <= byteResult[i] && byteResult[i] <= '9' {
			byteResult[i] = byteResult[i] - '0' + 'A'
		}
	}
	return string(byteResult)
}
