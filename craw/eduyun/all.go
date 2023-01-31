package eduyun

import (
	"main.go/tuuz"
)

func Craw_all(province, pageNo, pageSize int) {
	craw_init_db_all()
	craw_to_end(province, pageNo, pageSize)
	craw_db_dealer()
}
func craw_init_db_all() {
	db := tuuz.Db().Table("c_eduyun")
	db.Truncate()
}
