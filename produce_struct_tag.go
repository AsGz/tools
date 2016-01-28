// produce struct tag code for object
//
// Example:
//
// type MyStruct struct {
// 	Name      string
// 	MaxHeight int
// }
//
// var s MyStruct
//
// fmt.Prinln( tool.ProduceStructTag(s, "json"))
//
// type MyStruct struct {
// 	Name      string `json:"name"`
// 	MaxHeight int    `json:"max_height"`
// }
//
package tools

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func getTagName(currName, tag string) (newName string) {
	first := true
	for _, r := range currName {
		if unicode.IsUpper(r) {
			if first {
				newName = fmt.Sprintf("%s%s", newName, strings.ToLower(string(r)))
				first = false
			} else {
				newName = fmt.Sprintf("%s_%s", newName, strings.ToLower(string(r)))
			}
		} else {
			newName = fmt.Sprintf("%s%s", newName, string(r))
		}
	}
	newName = fmt.Sprintf("`%s:\"%s\"`", tag, newName)
	return
}

func ProduceStructTag(obj interface{}, tag string) string {
	var newDefineCode string
	s := reflect.ValueOf(obj)
	newDefineCode = fmt.Sprintf("type %s struct {\n", s.Type().String())
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		n := s.Type().Field(i).Name
		newDefineCode = fmt.Sprintf("%s\t%s\t%s\t\t%s\n",
			newDefineCode,
			n,
			f.Type(),
			getTagName(n, tag))
	}
	newDefineCode = fmt.Sprintf("%s}\n", newDefineCode)
	return newDefineCode
}
