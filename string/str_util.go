package estrings

import (
	"fmt"
	"net/url"
	"path"
	"regexp"

	"golang.org/x/net/publicsuffix"
)

// IsStrEmpty 判断字符串是否为空
func IsStrEmpty(str string) bool {
	return str == ""
}

// IsAnyStrEmpty 判断多个字符串中是否有任意一个为空
func IsAnyStrEmpty(strs ...string) bool {
	if strs != nil {
		for _, str := range strs {
			if IsStrEmpty(str) {
				return true
			}
		}
	}
	return false
}

// IsAllNotEmpty 所有字符串均非空
func IsAllNotEmpty(strs ...string) bool {
	return !IsAnyStrEmpty(strs...)
}

// IsAnyStrNotEmpty 判断多个字符串中是否有任意一个不为空
func IsAnyStrNotEmpty(strs ...string) bool {
	for _, str := range strs {
		if !IsStrEmpty(str) {
			return true
		}
	}
	return false
}

func GetDomainByUrl(urlStr string) (tld string, err error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	if IsStrEmpty(u.Host) {
		return urlStr, nil
	}
	var domain string
	for i := len(u.Host) - 1; i >= 0; i-- {
		if u.Host[i] == ':' {
			domain = u.Host[:i]
			break
		} else if u.Host[i] < '0' || u.Host[i] > '9' {
			domain = u.Host
			break
		}
	}
	tld, err = publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return "", err
	}
	//suffix, icann := publicsuffix.PublicSuffix(strings.ToLower(domain))
	//if err != nil && !icann && suffix == domain {
	//	tld = domain
	//	err = nil
	//}
	//if err != nil {
	//	return "", err
	//}

	return tld, nil
}

// RenderTemplate 渲染文本中的占位符，按顺序替换为 values 数组中的内容
func RenderTemplate(text string, values ...string) string {
	// 正则表达式匹配所有的 {{}} 占位符
	re := regexp.MustCompile(`\{\{}}`)
	// 使用替换函数依次替换每个 {{}} 为 values 中的相应值
	// 使用闭包按顺序返回替换值
	result := re.ReplaceAllStringFunc(text, func(s string) string {
		if len(values) == 0 {
			return s // 如果没有更多的值可以替换，直接返回原占位符
		}
		// 获取第一个值并移除它
		val := values[0]
		values = values[1:]
		return val
	})
	return result
}

// URLPathJoin 拼接基础 URL 和路径部分，返回拼接后的完整 URL。
func URLPathJoin(baseURL, pathSegment string) (string, error) {
	// 解析基础 URL
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("URL parsing error: %w", err)
	}
	// 使用 path.Join 来拼接路径，确保路径正确处理
	fullPath := path.Join(base.Path, pathSegment)
	// 拼接后的完整 URL
	fullURL := base.ResolveReference(&url.URL{
		Path: fullPath,
	})
	// 返回拼接后的完整 URL
	return fullURL.String(), nil
}
