package route

import (
	"akigate/reverseProxy"
	"akigate/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/18 17:21
 * @Title:
 * --- --- ---
 * @Desc:
 */
type Entry struct {
	rp *reverseProxy.ReverseProxy
	rm *RoutesMap
}

func NewEntry(rp *reverseProxy.ReverseProxy, rm *RoutesMap) *Entry {
	return &Entry{rp, rm}
}

func (e *Entry) In(ctx *gin.Context) {
	url := ctx.Request.URL
	if url.Path == "/" {
		ctx.JSON(http.StatusNotFound, &gin.H{
			"msg": "路径错误",
		})
	}
	key := util.GetRouteKey(url.Path)
	host := e.rm.Get(key)
	if host == "" {
		ctx.JSON(http.StatusNotFound, &gin.H{
			"msg": "请求服务不存在",
		})
		return
	}

	e.rp.SetTarget(host)
	e.rp.ServeProxy(ctx)
}

func (e *Entry) Test(ctx *gin.Context) {
	ctx.JSON(200, &gin.H{
		"data": "wtf?!",
	})
}
