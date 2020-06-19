package reverseProxy

import "net/http"

/**
 * @Author: Tomonori
 * @Date: 2020/6/19 14:21
 * @Title:
 * --- --- ---
 * @Desc:
 */

type Req struct {
	*http.Request
}

func (r *Req) PatchHeaders() {
	r.Host = r.URL.Host
	return
}
