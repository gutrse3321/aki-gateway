package route

import (
	"github.com/gin-gonic/gin"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/18 17:21
 * @Title:
 * --- --- ---
 * @Desc:
 */
type Entry struct {
}

func NewEntry() *Entry {
	return &Entry{}
}

func (e *Entry) In(ctx *gin.Context) {

}
