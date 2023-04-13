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
	route.Any("link", dazhong_link)
}

type MtCraw struct {
	c     *colly.Collector
	maxid int64
}

const Table = "c_dazhong_xuexipeixun"

func dazhong_1(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		fmt.Println(err)
	}
	html := string(body)
	re := regexp.MustCompile(`<a .*? data-shopid="(.*?)" .*? title="(.*?)" .*? href="(.*?)">`)
	matchess := re.FindAllStringSubmatch(html, -1)
	if len(matchess) > 0 {
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
				db := tuuz.Db().Table(Table)
				thedata, err := db.Where("name", title).Find()
				if err != nil {
					Log.Dbrr(err, tuuz.FUNCTION_ALL())
				}
				if len(thedata) < 1 {
					tuuz.Db().Table(Table).Data(data).Insert()
					fmt.Println(shopid, title, href)
				}
			} else {
				fmt.Println("No match found")
			}
		}
	} else {
		//fmt.Println(html)
		re := regexp.MustCompile(`\/shop\/(\w+)\/review`)
		match := re.FindStringSubmatch(html)
		if len(match) > 1 {
			shopid := match[1]
			re = regexp.MustCompile(`<div class="address">\s+<span class="item">地址：</span>\s+(.*?)\s+</div>`)

			match := re.FindStringSubmatch(html)
			address := ""
			phone := ""
			if len(match) > 1 {
				address = match[1]
			} else {
				fmt.Println("address No match found")
				return
			}

			re = regexp.MustCompile(`data-phone="(\d+-\d+)"`)
			result := re.FindStringSubmatch(html)

			if len(result) == 2 {
				phone = result[1]
			} else {
				fmt.Println("phone No match found.")
				//return
			}
			tuuz.Db().Table(Table).Where("shopid", shopid).Data(map[string]any{
				"address": address,
				"phone":   phone,
			}).Update()
		} else {
			fmt.Println("shopid No match found.")
		}
		//re := regexp.MustCompile(`<div class="address">\s+<span class="item">地址：</span>\s+(.*?)\s+</div>`)

	}

}

func dazhong_link(c *gin.Context) {
	datas, err := tuuz.Db().Table(Table).Where("address", "=", "").Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return
	}
	c.Header("Content-Type", "text/html")
	c.String(200, "<html><body>")
	for _, data := range datas {
		c.String(200, "<a href=\"%s\">%s</a></br>\n", data["url"], data["url"])
	}
	c.String(200, "</body></html>")
}
