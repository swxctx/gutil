package gutil

import "regexp"

var (
	phoneRx, _ = regexp.Compile("^[1][3456789]\\d{9}")
	emailRegex = `^(|(([A-Za-z0-9]+_+)|([A-Za-z0-9]+\-+)|([A-Za-z0-9]+\.+)|([A-Za-z0-9]+\++))*[A-Za-z0-9]+@((\w+\-+)|(\w+\.))*\w{1,63}\.[a-zA-Z]{2,6})$`
	EmailRegex *regexp.Regexp
)

func init() {
	EmailRegex = regexp.MustCompile(emailRegex)
}

// IsValidPhone 识别手机号码是否符合正确格式
func IsValidPhone(phone string) bool {
	if len(phone) != 11 || (phone[0] != '1') {
		return false
	}
	return phoneRx.Match([]byte(phone))
}

// IsValidEmail 邮箱格式校验
func IsValidEmail(email string) bool {
	if !EmailRegex.MatchString(email) || regexp.MustCompile(`^www\.`).MatchString(email) || regexp.MustCompile(`con$`).MatchString(email) {
		return false
	}
	return true
}
