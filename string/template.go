package estrings

import (
	"strings"
)

// FillTemplate 使用 params 中的键值对填充模板字符串 tpl。
// 占位符的格式为 {key}。
func FillTemplate(tpl string, params map[string]string) string {
	// 如果参数为空，直接返回原始模板
	if len(params) == 0 {
		return tpl
	}

	// strings.NewReplacer 对于多个替换操作性能更高
	// 它会构建一个高效的替换器
	args := make([]string, 0, len(params)*2)
	for key, value := range params {
		placeholder := "{" + key + "}"
		args = append(args, placeholder, value)
	}

	replacer := strings.NewReplacer(args...)
	return replacer.Replace(tpl)
}
