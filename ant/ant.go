package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

func main() {
	current, _ := os.Getwd()
	fmt.Println(current)
	folder := substr(current, 0, strings.LastIndex(current, "/")) + "/files/"
	ListFolder(folder)
	path := "JSON-data.json"
	content := GetTextFileContent(folder, path)
	ruleIds := DistinctRuleId(content)
	fmt.Println("Origin Size: ", len(ruleIds))
	ruleIds = removeRepeatElement(ruleIds)
	fmt.Println("Distinct Size: ", len(ruleIds))
	fmt.Println(ruleIds)
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func ListFolder(name string) {

	fileSystem := os.DirFS(name)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)
		return nil
	})
}

func GetTextFileContent(path string, fileName string) (content string) {
	file := os.DirFS(path)
	contentByte, error := fs.ReadFile(file, fileName)
	if error == nil {
		content := string(contentByte)
		// fmt.Println(content)
		return content
	}
	log.Fatal(error)
	return ""
}

type RuleStruct struct {
	ruleId     string
	codeIssuer string
}

func DistinctRuleId(jsonContent string) (ruleIds []string) {
	// p := []RuleStruct{}
	var p []interface{}
	error := json.Unmarshal([]byte(jsonContent), &p)
	if error != nil {
		log.Fatal(error)
	}
	_ruleIds := make([]string, len(p))
	fmt.Println(len(p))
	for index, v := range p {
		// fmt.Println(v)
		m := v.(map[string]interface{})
		for k, vv := range m {
			if k == "ruleId" {
				// fmt.Println(vv)
				result, ok := vv.(string)
				if ok {
					_ruleIds[index] = result
				} else {
					log.Fatal(ok)
				}
			}
		}
	}
	return _ruleIds
}

func removeRepeatElement[T comparable](list []T) []T {
	// 创建一个临时map用来存储数组元素
	temp := make(map[T]struct{})
	index := 0
	// 将元素放入map中
	for _, v := range list {
		temp[v] = struct{}{}
	}
	tempList := make([]T, len(temp))
	for key := range temp {
		tempList[index] = key
		index++
	}
	return tempList
}
