package filter

import (
	"github.com/gin-gonic/gin"
	httpServer "github.com/gutrse3321/aki/pkg/transports/http"
	"gopkg.in/oauth2.v4/server"
	"net/http"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/22 15:07
 * @Title:
 * --- --- ---
 * @Desc:
 */

var (
	TwinklingList = make(map[string]bool)
)

func init() {
	TwinklingList["/uaa/oauth/login"] = true
}

func UserAuthMiddleware(oauthServer *server.Server) httpServer.Middleware {
	return func(r *gin.Engine) {
		r.Use(func(c *gin.Context) {
			if twinkling := TwinklingList[c.Request.URL.Path]; twinkling {
				c.Next()
				return
			}

			_, err := oauthServer.ValidationBearerToken(c.Request)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{
					"error_msg": "认证错误:" + err.Error(),
				})
				return
			}
			c.Next()
		})
	}
}
