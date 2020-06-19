package util

import (
	"fmt"
	"testing"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/19 12:14
 * @Title:
 * --- --- ---
 * @Desc:
 */

func TestSplitScheme(t *testing.T) {
	scheme := SplitScheme("http://asuka.red")
	println(scheme)
}

func TestSplitSlashArr(t *testing.T) {
	arr := SplitSlashArr("localhost:4000/bbs/getUser")
	fmt.Printf("%+v\n", arr)
}

func TestGetRouteKey(t *testing.T) {
	key := GetRouteKey("localhost:4000/bbs/getUser")
	println(key)
}
