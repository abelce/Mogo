/*
 * File: utils.go
 * Project: Mogo
 * File Created: Thursday, 16th May 2019 3:24:54 pm
 * Author: zxtang (1061225829@qq.com)
 * -----
 * Last Modified: Thursday, 16th May 2019 3:24:54 pm
 * Modified By: zxtang (1061225829@qq.com>)
 * -----
 * Copyright 2017 - 2019 Your Company, Your Company
 */
package mogo

import (
	"log"
	"os"
	"regexp"
)

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func trimHTML(str string) string {
	if str == "" {
		return str
	}
	re, _ := regexp.Compile(`<[\s\S]+?(>|$)`)
	newstr := re.ReplaceAllString(str, "")
	return newstr
}

func subStr(str string, start, end int) string {
	if start < 0 {
		log.Panic("start position is wrong!")
	}
	if end > len(str) {
		log.Panic("end positon is wrong!")
	}
	if start > end {
		log.Panic("wrong position!")
	}

	rs := []rune(str)
	return string(rs[start:end])
}
