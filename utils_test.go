package utils

import (
	"testing"
)

//检查字符串,去掉特殊字符
//\w 匹配字母或数字或下划线或汉字 等价于 ‘[A-Za-z0-9_]’。
func TestCheckCharDoSpecial(t *testing.T) {
	tt := []struct {
		s        string //输入的字符串
		split    byte   //需要分隔的标识
		reg      string //需要的正则
		expected string //预期的值
	}{
		{"w$$w               !@#!w.!@s#g^f^oot.co))(世界你好m", '.', `[a-z\.]`, "www.sgfoot.com"},
		{"he!@#!ll*(*白蛇,缘起owor  ld", ' ', "[a-z]", "helloworld"},
		{",json,xml,gorm,", ',', `[a-z\,]`, "json,xml,gorm"},
		{"洛a阳b城c里d见e秋f风，欲g作h家i书j意k万l重", ' ', `[a-z]`, "abcdefghijkl"},        //查找连续的小写字母
		{"洛A阳B城C里D见E秋F风，欲G作H家I书J意K万L重", ' ', `[A-Z]+`, "ABCDEFGHIJKL"},       //查找连续的大写字母
		{"洛A阳B城C里D见E秋F风，欲G作H家I书J意K万L重", ' ', `[[:upper:]]+`, "ABCDEFGHIJKL"}, //查找连续的大写字母
		{",1ab**c,2a%^-b,3a_b$,", '_', `[\w]`, "1abc2ab3a_b"},                //匹配字母或数字或下划线或汉字
		{",1ab**c,2a%^-b,3a_b$,", ',', `[\w\,]`, "1abc,2ab,3a_b"},            //匹配字母或数字或下划线或汉字
	}
	for _, item := range tt {
		actual := CheckCharDoSpecial(item.s, item.split, item.reg)
		if item.expected != actual {
			t.Errorf("[%s] splie: %c actual: %s => expected %s", item.s, item.split, actual, item.expected)
		}
	}
}

func TestInArray(t *testing.T) {
	tests := []struct{
		need interface{}
		haystack interface{}
		expected bool
	} {
		{1, []int{1,2,3}, true},
		{"a", []string{"a", "b"}, true},
		{22, []int{1,2,3}, false},
		{1.78, []float64{1,2,3}, false},
		{3.78, []float64{1,2,3, 3.78}, true},
	}
	for _, item := range tests {
		actual := InArray(item.need, item.haystack)
		if actual != item.expected {
			t.Errorf("need: %v, haystack: %v, actual: %v", item.need, item.haystack, actual)
		}
	}
}