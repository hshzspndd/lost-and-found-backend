package utils

// 获取字符串前15个字符
func SubStr(s string, n int) string {
	front := []rune(s)
	if len(front) <= n {
		return s
	}
	return string(front[:n]) + "..."
}
