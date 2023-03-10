package SystemParamModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "system_param"

func Api_find_val(key interface{}) interface{} {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"key": key,
	}
	db.Where(where)
	ret, err := db.Value("val")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_set_val(key, val interface{}) bool {
	db := tuuz.Db().Table(Table)
	db.Where("key", key)
	db.Data(map[string]interface{}{
		"val": val,
	})
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_KV() map[string]string {
	db := tuuz.Db().Table(Table)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		arr := map[string]string{}
		for _, data := range ret {
			arr[data["key"].(string)] = data["val"].(string)
		}
		return arr
	}
}

func Api_select() []gorose.Data {
	db := tuuz.Db().Table(Table)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
