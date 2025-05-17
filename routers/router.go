package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartServer() {
	router := gin.Default()

	branchItemInitiator(router)
	branchInitiator(router)
	employeeInitiator(router)
	itemInitiator(router)
	itemTypeInitiator(router)
	salesDataInitiator(router)
	userInitiator(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
