package eduyun

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

func Craw_to_end(province, pageNo int, pageSize int) error {
	page, err := craw_page(province, pageNo, pageSize)
	if err != nil {
		return err
	}
	lists := page.List
	db := tuuz.Db().Table("c_eduyun")
	//data := map[string]any{
	//	"schoolLicenseNo": schoolLicenseNo,
	//	"corpProvince":    schoolLicenseNo,
	//	"subject":         subject,
	//	"corpName":        corpName,
	//	"corpLogo":        corpLogo,
	//	"pkSpInfo":        pkSpInfo,
	//	"corpArea":        corpArea,
	//	"certificateNo":   certificateNo,
	//	"corpCity":        corpCity,
	//	"cityName":        cityName,
	//	"areaName":        areaName,
	//	"provinceName":    provinceName,
	//	"businessType":    businessType,
	//	"object":          object,
	//}
	db.Data(lists)
	_, err = db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return err
	}
	if page.PageNo >= page.TotalPage {
		Craw_to_end(province, pageNo+1, pageSize)
	}
	fmt.Println("eduyun采集完成")
	return nil
}

func craw_page(province, pageNo, pageSize any) (Page, error) {
	var page Page
	post := map[string]interface{}{
		"province": province,
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
