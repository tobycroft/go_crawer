package main

import (
	"main.go/config/app_conf"
	"main.go/craw/mt"
	"os"
	"time"
)

func init() {
	if app_conf.TestMode == false {
		s, err := os.Stat("./log/")

		if err != nil {
			os.Mkdir("./log", 0755)
		} else if s.IsDir() {
			os.Mkdir("./log", 0755)
		}
	}
}

func main() {

	//Calc.RefreshBaseNum()
	//mainroute := gin.Default()
	////gin.SetMode(gin.ReleaseMode)
	////gin.DefaultWriter = ioutil.Discard
	//mainroute.SetTrustedProxies([]string{"0.0.0.0/0"})
	//mainroute.SecureJsonPrefix(app_conf.SecureJsonPrefix)
	//route.OnRoute(mainroute)
	//mainroute.Run(":80")
	//post := map[string]interface{}{
	//	"province": 350000,
	//	"pageNo":   2,
	//	"pageSize": 1,
	//}
	//eduyun.Craw_to_end(350000, 1, 500)
	var mtc mt.MtCraw
	go mtc.Craw_insert()
	mtc.Craw_Init()
	mtc.Craw_ready()
	mtc.Craw_start()
	for {
		time.Sleep(1 * time.Second)
	}
}
