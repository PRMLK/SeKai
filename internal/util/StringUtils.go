package util

import (
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
