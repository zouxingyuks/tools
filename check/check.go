package check

import (
	"fmt"
	"regexp"
	"strings"
)

// CheckChinese 中文检查
// 仅允许中文、英文字母、数字和空白字符，和常见标点符号，过滤特殊字符，特别是sql注入
func CheckChinese(input *string) bool {
	// 去除首位空格
	userInput := strings.TrimSpace(*input)
	if ok, err := regexp.MatchString(`^([\p{L}\p{N}\s.,;：'！？。]|\p{Han})+$`, userInput); !ok {
		fmt.Println(err, ok)
		return false
	}
	return true
}

// 手机号码检查
func CheckPhoneNumber(input *string) bool {
	// 去除首位空格
	userInput := strings.TrimSpace(*input)
	re1 := regexp.MustCompile(`^\+?[\d\s-]+$`)
	re2 := regexp.MustCompile(`^\+\d{1,2}\s?\(\d{1,3}\)\s?\d{1,4}[\s-]?\d{1,4}[\s-]?\d{1,9}$`)
	re3 := regexp.MustCompile(`^\+\d{1,2}\s?\d{1,4}[\s-]?\d{1,4}[\s-]?\d{1,9}$`)
	re4 := regexp.MustCompile(`^(\+\d{1,2}\s?)?(\(\d{1,3}\)\s?)?(\d{1,4}[\s-]?)?\d{1,4}[\s-]?\d{1,9}$`)

	// 第一步基本格式校验
	ok1 := re1.MatchString(userInput)
	if !ok1 {
		return false
	}

	// 第二步更具体的电话号码格式校验
	ok2 := re2.MatchString(userInput)
	ok3 := re3.MatchString(userInput)
	if !(ok2 || ok3) {
		// 不符合带括号的格式，再进行位数校验
		ok4 := re4.MatchString(userInput)
		if !ok4 {
			return false
		}
		// 进行位数校验，此处假设电话号码的有效位数范围是10-15
		numberOnly := regexp.MustCompile(`\d`)
		digits := numberOnly.FindAllString(userInput, -1)
		if len(digits) < 10 || len(digits) > 15 {
			return false
		}
	}

	return true
}
