package utils

import (
	"regexp"
	"strings"
	"bytes"
)

//检查字符串,去掉特殊字符
func CheckCharDoSpecial(s string, char byte, regs string) string {
	reg := regexp.MustCompile(regs)
	var result []string
	if arr := reg.FindAllString(s, -1); len(arr) > 0 {
		buf := bytes.Buffer{}
		for key, val := range arr {
			if val != string(char) {
				buf.WriteString(val)
			}
			if val == string(char) && buf.Len() > 0 {
				result = append(result, buf.String())
				buf.Reset()
			}
			//处理最后一批数据
			if buf.Len() > 0 && key == len(arr) - 1 {
				result = append(result, buf.String())
			}
		}
	}
	return strings.Join(result, string(char))
}
