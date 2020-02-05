package services

import (
	"fmt"
	"strings"
)

var magicTable = map[string]string{
	"\xff\xd8\xff"		:   "jpeg",
	"\x89PNG\r\n\x1a\n"	: 	"png",
	"GIF87a"			:   "gif",
	"GIF89a"			:   "gif",
}

func CheckValidity(file []byte) bool {
	fileStr := []byte(file)
	for magic, _ := range magicTable {
		if strings.HasPrefix(string(fileStr), magic) {
			fmt.Println("invalid")
			return true
		}
	}
	fmt.Println("valid")
	return false
}

