package main

import (
	"github.com/gin-gonic/gin"
	"github.com/weinibear/admin-server/models"
	"github.com/weinibear/admin-server/pkg/setting"
	"github.com/weinibear/admin-server/routers"
)


func init() {
	setting.Setup()
	models.Setup()
}


func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
}