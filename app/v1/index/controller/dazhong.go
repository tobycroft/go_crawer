package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

func DazhongController(route *gin.RouterGroup) {

	route.Any("1", dazhong_1)
}

type MtCraw struct {
	c     *colly.Collector
	maxid int64
}

func dazhong_1(c *gin.Context) {

}
