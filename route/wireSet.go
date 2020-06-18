package route

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	httpServer "github.com/gutrse3321/aki/pkg/transports/http"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/18 17:22
 * @Title:
 * --- --- ---
 * @Desc:
 */

func CreateInitControllersFn(entry *Entry) httpServer.InitControllers {
	return func(r *gin.Engine) {
		r.POST("/*entry", entry.In)
	}
}

var WireSet = wire.NewSet(
	NewEntry,
	CreateInitControllersFn,
)
