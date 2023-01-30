package eduyun

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

func Craw_to_end(province, pageNo, pageSize int) {
	craw_init_db()
	craw_to_end(province, pageNo, pageSize)
	craw_db_dealer()
}
func craw_init_db() {
	db := tuuz.Db().Table("c_eduyun_fujian")
	db.Truncate()
}

func craw_db_dealer() {
	tuuz.Db().Execute("update c_eduyun_fujian set `subject`=\"学科,非学科\" where `subject` not in (\"0\",\"1\");")
	tuuz.Db().Execute("update c_eduyun_fujian set `subject`=\"非学科\" where `subject`=\"0\";")
	tuuz.Db().Execute("update c_eduyun_fujian set `subject`=\"学科\" where `subject`=\"1\";")

	tuuz.Db().Execute("update c_eduyun_fujian set `businessType`=\"线上\" where `businessType`=\"0\";")
	tuuz.Db().Execute("update c_eduyun_fujian set `businessType`=\"线下\" where `businessType`=\"1\";")

	tuuz.Db().Execute("UPDATE c_eduyun_fujian set object=REPLACE(object, \"0\", \"学龄前\");")
	tuuz.Db().Execute("UPDATE c_eduyun_fujian set object=REPLACE(object, \"1\", \"义务教育阶段\");")
	tuuz.Db().Execute("UPDATE c_eduyun_fujian set object=REPLACE(object, \"2\", \"高中阶段\");")
	tuuz.Db().Execute("UPDATE c_eduyun_fujian set object=REPLACE(object, \"3\", \"成人\");")
	tuuz.Db().Execute("UPDATE c_eduyun_fujian set object=REPLACE(object, \"4\", \"中职\");")
}
func craw_to_end(province, pageNo int, pageSize int) error {
	fmt.Println("province", province, "pageNo", pageNo, "pageSize", pageSize)
	page, err := craw_page(province, pageNo, pageSize)
	if err != nil {
		return err
	}
	lists := page.List
	db := tuuz.Db().Table("c_eduyun_fujian")
	db.Data(lists)
	_, err = db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return err
	} else {
		fmt.Println("eduyun采集完成")
	}
	if page.PageNo < page.TotalPage {
		return craw_to_end(province, pageNo+1, pageSize)
	}
	return nil
}

func craw_page(province, pageNo, pageSize any) (Page, error) {
	var page Page
	post := map[string]interface{}{
		//"province": province,
		"pageNo":   pageNo,
		"pageSize": pageSize,
	}
	ret, err := Net.Post("https://xwpx.eduyun.cn/bmp-web/tolSpInfo/getSpInfoList", nil, post, nil, nil)
	if err != nil {
		return page, err
	}
	err = jsoniter.UnmarshalFromString(ret, &page)
	if err != nil {
		return page, err
	}
	return page, err
}

type Page struct {
	PageNo      int    `json:"pageNo"`
	PageSize    int    `json:"pageSize"`
	Count       int    `json:"count"`
	TotalPage   int    `json:"totalPage"`
	First       int    `json:"first"`
	Last        int    `json:"last"`
	Prev        int    `json:"prev"`
	Next        int    `json:"next"`
	FirstPage   bool   `json:"firstPage"`
	LastPage    bool   `json:"lastPage"`
	List        []List `json:"list"`
	FirstResult int    `json:"firstResult"`
}
type List struct {
	SchoolLicenseNo string `json:"schoolLicenseNo"`
	CorpProvince    string `json:"corpProvince"`
	Subject         string `json:"subject"`
	CorpName        string `json:"corpName"`
	CorpLogo        string `json:"corpLogo"`
	PkSpInfo        string `json:"pkSpInfo"`
	CorpArea        string `json:"corpArea,omitempty"`
	CertificateNo   string `json:"certificateNo"`
	CorpCity        string `json:"corpCity"`
	CityName        string `json:"cityName"`
	AreaName        string `json:"areaName,omitempty"`
	ProvinceName    string `json:"provinceName"`
	BusinessType    string `json:"businessType"`
	Object          string `json:"object"`
}
