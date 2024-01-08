package main

import (
	"fmt"
	"net/http"

	"github.com/forevengetmathmmqqmm/goGinExample/global"
	"github.com/forevengetmathmmqqmm/goGinExample/initialize"
	"github.com/forevengetmathmmqqmm/goGinExample/pkg/setting"
	"github.com/forevengetmathmmqqmm/goGinExample/routers"
)

func main() {
	router := routers.InitRouter()
	global.GAV_DB = initialize.GormMySql()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
