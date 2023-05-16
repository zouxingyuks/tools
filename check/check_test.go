package check

import (
	"testing"
)

func TestCheckChinese(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{"test1", args{input: "s"}, true},
		{"test2", args{input: "测试"}, true},
		{"test3", args{input: "Hello, 世界"}, true},
		{"test4", args{input: "12345"}, true},
		{"test5", args{input: "中文校验函数"}, true},
		{"test6", args{input: "!@#$%^&*()_+"}, false},
		{"test7", args{input: "This is a test."}, true},
		{"test8", args{input: "123abcDEF"}, true},
		{"test9", args{input: "空白字符 \t \n"}, true},
		{"test10", args{input: "常见标点符号：！？。"}, true},
		{"sql 注入 1", args{input: "1 OR 1=1; --"}, false},
		{"sql 注入 2", args{input: "SELECT * FROM users WHERE username='admin' OR 1=1; --"}, false},
		{"sql 注入 3", args{input: "DROP TABLE users; --"}, false},
		{"sql 注入 4", args{input: "INSERT INTO users (username, password) VALUES ('admin', 'password'); --"}, false},
		{"sql 注入 5", args{input: "UPDATE users SET password='newpassword' WHERE username='admin'; --"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := CheckChinese(&tt.args.input); result != tt.wantResult {
				t.Errorf("CheckChinese() result = %v, wantResult %v", result, tt.wantResult)
			}
		})
	}
}

func TestCheckPhoneNumber(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{"合法电话号码-纯数字", args{input: "1234567890"}, true},
		{"合法电话号码-带横线", args{input: "123-456-7890"}, true},
		//此号码过于复杂，不予支持
		{"不支持手机号码-带括号", args{input: "(123) 456-7890"}, false},
		{"合法电话号码-国际格式", args{input: "+1 123-456-7890"}, true},
		{"非法电话号码-不完整", args{input: "123456789"}, false},
		{"非法电话号码-包含字母", args{input: "abcdefghij"}, false},
		{"非法电话号码-错误分隔符", args{input: "12-34-567-890"}, false},
		{"非法电话号码-非法字符", args{input: "1.800.123.4567"}, false},
		{"非法电话号码-括号不匹配", args{input: "+1 (123) 456-7890"}, false},
		{"合法电话号码-特殊字符空格", args{input: "123 456 7890"}, true},
		{"不支持电话号码-特殊字符点号", args{input: "123.456.7890"}, false},
		{"不支持电话号码-特殊字符斜杠", args{input: "123/456/7890"}, false},
		{"不支持电话号码-特殊字符下划线", args{input: "123_456_7890"}, false},
		{"合法电话号码-带国家代码", args{input: "+86 10 12345678"}, true},
		{"非法电话号码-空字符串", args{input: ""}, false},
		{"非法电话号码-包含特殊符号", args{input: "!@#$%^&*()"}, false},
		//此测试点是因为对手机号码做了格式化
		{"合法电话号码-带有空格", args{input: "1234567 890"}, true},
		{"非法电话号码-太长的号码", args{input: "1234567890123456789012345678901234567890"}, false},
		{"合法电话号码-带加号", args{input: "+1234567890"}, true},
		{"合法电话号码-带国际长途拨号前缀", args{input: "0111234567890"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := CheckPhoneNumber(&tt.args.input); result != tt.wantResult {
				t.Errorf("CheckPhoneNumber() result = %v, wantResult %v", result, tt.wantResult)
			}
		})
	}
}
