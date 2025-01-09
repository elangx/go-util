package tool

import (
	"strings"
	"unicode"
)

// 类三元运算符
func IF[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}

// 国家码转换成对应国家的国旗emoji
func CountryCodeToEmoji(countryCode string) string {
	if len(countryCode) != 2 {
		return ""
	}

	//台湾要转成CN，否则会出民国旗
	if strings.ToUpper(countryCode) == "TW" {
		countryCode = "CN"
	}

	var emoji string
	for _, r := range countryCode {
		//保证只有英文字母
		if !unicode.IsLetter(r) {
			return ""
		}
		// 计算Regional Indicator Symbol的Unicode码点
		emoji += string(unicode.ToUpper(r) + 0x1F1E6 - 'A')
	}
	return emoji
}
