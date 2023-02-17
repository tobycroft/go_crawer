package mt

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/tobycroft/Calc"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/tuuz"
	"strings"
	"sync"
)

type MtCraw struct {
	c     *colly.Collector
	maxid int64
}

var wg sync.WaitGroup

func (self *MtCraw) Craw_Init() {
	self.c = colly.NewCollector()
}

func (self *MtCraw) ManualVisit(addr string) {
	self.c.Visit(addr)
}

func (self *MtCraw) Craw_ready() {

	self.c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	self.c.OnResponse(func(e *colly.Response) {
		wg.Done()
		go func(bbody []byte) {
			body := string(bbody)
			//fmt.Println(body)
			bodys1 := strings.Split(body, "window.__INITIAL_STATE__ = ")
			bodys2 := bodys1[len(bodys1)-1]
			bodys3 := strings.Split(bodys2, "</script>")
			bodys4 := bodys3[0]
			s6 := strings.TrimSpace(bodys4)
			var datas Data
			err := jsoniter.UnmarshalFromString(s6, &datas)
			if err != nil {
				fmt.Println(self.maxid, "无数据")
				//Log.Crrs(errors.New(Calc.Any2String(self.maxid)+"无数据"), tuuz.FUNCTION_ALL())
				return
			}

			var bff BffData
			err2 := jsoniter.UnmarshalFromString(datas.BffData[0], &bff)
			if err2 != nil {
				fmt.Println(self.maxid, "无数据")
				//Log.Crrs(errors.New(Calc.Any2String(self.maxid)+"无数据"), tuuz.FUNCTION_ALL())
				return
			}
			fmt.Println("姓名:", bff.ResponseData[0].Data.Data.AttrValues.Name)
			if bff.ResponseData[0].Data.Data.AttrValues.Name == "" {
				fmt.Println(self.maxid, "姓名空数据")
				//Log.Crrs(errors.New(Calc.Any2String(self.maxid)+"姓名空数据"), tuuz.FUNCTION_ALL())
				return
			}
			fmt.Println(bff.ResponseData[0].Data.Data.AttrValues.Skills)
			fmt.Println(bff.ResponseData[0].Data.Data.AttrValues.WorkYears)
			fmt.Println(bff.ResponseData[0].Data.Data.AttrValues.WorkYearsStr)
			fmt.Println(bff.ResponseData[0].Data.Data.TechnicianID)
			if bff.ResponseData[0].Data.Data.TechnicianID == 0 {
				fmt.Println(Calc.Any2String(self.maxid) + "技师id=0")
				//Log.Crrs(errors.New(Calc.Any2String(self.maxid)+"技师id=0"), tuuz.FUNCTION_ALL())
				return
			}
			fmt.Println(bff.ResponseData[0].Data.Data.AttrValues.PhotoURL)
			fmt.Println(bff.ResponseData[0].Data.Data.ShopIDForFe)

			fmt.Println(bff.ResponseData[0].Data.Data.Share.Title)
			fmt.Println(bff.ResponseData[0].Data.Data.Share.Desc)
			fmt.Println(bff.ResponseData[0].Data.Data.Share.URL)
			crawData <- bff
		}(e.Body)
	})

}

var crawData = make(chan BffData, 1024)

func (self *MtCraw) Craw_insert() {
	for bff := range crawData {
		db := tuuz.Db().Table("mt_craw")
		db.Where("techid", bff.ResponseData[0].Data.Data.TechnicianID)
		ret, err := db.Find()
		if err != nil {
			panic(err)
			return
		}
		id := "0"
		if len(ret) > 1 {
			id = Calc.Any2String(ret["id"])
		}
		db = tuuz.Db().Table("mt_craw")
		data := map[string]interface{}{
			"name":         bff.ResponseData[0].Data.Data.AttrValues.Name,
			"skills":       strings.Join(bff.ResponseData[0].Data.Data.AttrValues.Skills, ","),
			"workyears":    bff.ResponseData[0].Data.Data.AttrValues.WorkYears,
			"workyearsstr": bff.ResponseData[0].Data.Data.AttrValues.WorkYearsStr,
			"techid":       bff.ResponseData[0].Data.Data.TechnicianID,
			"photo":        bff.ResponseData[0].Data.Data.AttrValues.PhotoURL,
			"shopidforfe":  bff.ResponseData[0].Data.Data.ShopIDForFe,
			"title":        bff.ResponseData[0].Data.Data.Share.Title,
			"desc":         bff.ResponseData[0].Data.Data.Share.Desc,
			"url":          bff.ResponseData[0].Data.Data.Share.URL,
			"repeat_id":    id,
		}
		db.Data(data)
		_, err = db.Insert()
		if err != nil {
			panic(err)
			return
		}

	}
}

func (self *MtCraw) Craw_start() {
	wg.Add(5)
	mtid := SystemParamModel.Api_find_val("mtid")
	self.maxid = Calc.Any2Int64(mtid)

	time := int64(5)

	fmt.Println("jishi:", self.maxid+time)
	for i := int64(1); i <= time; i++ {
		go self.c.Visit("https://g.meituan.com/domino/craftsman-app/craftsman-detail.html?technicianId=" + Calc.Any2String(self.maxid+i))
	}
	wg.Wait()
	SystemParamModel.Api_set_val("mtid", self.maxid+time)
	self.Craw_start()
}
