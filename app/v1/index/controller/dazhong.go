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
		city := ""
		re = regexp.MustCompile(`<span class="J-current-city">(.*?)</span>`)
		match := re.FindStringSubmatch(html)

		if len(match) > 1 {
			city = match[1]
		} else {
			fmt.Println("未找到城市")
			return
		}
		cata := ""

		re = regexp.MustCompile(`<a href="[^"]+" class="cur" data-cat-id="[^"]+" data-click-name=".*."><span>([^<]+)</span></a>`)
		match = re.FindStringSubmatch(html)
		if len(match) > 1 {
			cata = match[1] // Output: 书法
		} else {
			fmt.Println("未找到cata")
			return
		}
		for _, matches := range matchess {
			if len(matches) > 0 {
				shopid := matches[1]
				title := matches[2]
				href := matches[3]
				data := map[string]any{
					"name":   title,
					"cata":   cata,
					"shopid": shopid,
					"url":    href,
					"city":   city,
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
		shopid := ""
		re := regexp.MustCompile(`\/shop\/(\w+)\/review`)
		match := re.FindStringSubmatch(html)
		if len(match) > 1 {
			shopid = match[1]
		} else {
			re = regexp.MustCompile(`shopId:'(\w+)'`)
			match := re.FindStringSubmatch(html)
			if len(match) > 1 {
				shopid = (match[1])

			} else {
				fmt.Println("shopid No match found.")
				return
			}
			//fmt.Println(html)
		}

		re = regexp.MustCompile(`<div class="address">\s+<span class="item">地址：</span>\s+(.*?)\s+</div>`)

		match = re.FindStringSubmatch(html)
		address := ""
		phone := ""
		if len(match) > 1 {
			address = match[1]
		} else {
			fmt.Println("address No match found")
			return
		}

		re = regexp.MustCompile(`data-phone="(\d+.\d+)"`)
		result := re.FindStringSubmatch(html)

		if len(result) == 2 {
			phone = result[1]
		} else {
			fmt.Println("phone No match found.")
			fmt.Println(result)
		}

		tuuz.Db().Table(Table).Where("shopid", shopid).Data(map[string]any{
			"address": address,
			"phone":   phone,
		}).Update()
		fmt.Println("update", shopid, address, phone)
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
