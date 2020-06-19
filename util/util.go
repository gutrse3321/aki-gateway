package util

import "strings"

/**
 * @Author: Tomonori
 * @Date: 2020/6/19 12:09
 * @Title:
 * --- --- ---
 * @Desc:
 */

/*
	分隔http:// 返回后面的网址
*/
func SplitScheme(url string) string {
	scheme := strings.HasPrefix(url, "http://")
	if scheme {
		return url[7:]
	}
	return ""
}

/*
	分隔斜杠 返回数组
*/
func SplitSlashArr(url string) []string {
	scheme := SplitScheme(url)
	if scheme != "" {
		url = scheme
	}
	slash := strings.Split(url, "/")
	return slash
}

/*
	获取路由键
*/
func GetRouteKey(url string) string {
	arr := SplitSlashArr(url)
	if len(arr) < 2 {
		return ""
	}
	return arr[1]
}
