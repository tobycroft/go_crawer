package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"regexp"
)

func DazhongController(route *gin.RouterGroup) {

	route.Any("1", dazhong_1)
}

type MtCraw struct {
	c     *colly.Collector
	maxid int64
}

func dazhong_1(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		fmt.Println(err)
	}
	html := string(body)
	re := regexp.MustCompile(`<a .*? data-shopid="(.*?)" .*? title="(.*?)" .*? href="(.*?)">`)
	matchess := re.FindAllStringSubmatch(html, -1)

	for _, matches := range matchess {
		if len(matches) > 0 {
			shopid := matches[1]
			title := matches[2]
			href := matches[3]
			data := map[string]any{
				"name":   title,
				"shopid": shopid,
				"url":    href,
			}
			db := tuuz.Db().Table("c_dazhong_xuexipeixun")
			thedata, err := db.Where("name", title).Find()
			if err != nil {
				Log.Dbrr(err, tuuz.FUNCTION_ALL())
			}
			if len(thedata) < 1 {
				tuuz.Db().Table("c_dazhong_xuexipeixun").Data(data).Insert()
				fmt.Println(shopid, title, href)
			}
		} else {
			fmt.Println("No match found")
		}
	}

}
