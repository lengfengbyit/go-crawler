package util

import "strings"

// 删除所有空格和换行
func TrimSpaceAndLinefeed(str string) string {
	return strings.ReplaceAll(strings.ReplaceAll(str, "\n", ""), " ", "")
}
