package word

import (
	"strings"
	"unicode"
)

// 单词全部转为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// 单词全部转为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 下划线单词转为大写驼峰
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// 下划线单词转为小写驼峰
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// 驼峰转为下划线
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 { // 首字母小写
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) { // 大写字母前加下划线
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r)) // 字母转小写
	}
	return string(output)
}
