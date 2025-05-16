package routers

import "github.com/gin-gonic/gin"

func StartServer() {
	router := gin.Default()

	branchItemInitiator(router)
	branchInitiator(router)
	employeeInitiator(router)
	itemInitiator(router)
	itemTypeInitiator(router)
	salesDataInitiator(router)
	userInitiator(router)

	router.Run(":8080")
}
