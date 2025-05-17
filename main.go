package main

import (
	"simple-store-management/configs"
	"simple-store-management/databases/connection"
	"simple-store-management/databases/migration"
	_ "simple-store-management/docs"
	"simple-store-management/routers"

	_ "github.com/lib/pq"
)

// @title           Simple Store Management API
// @version         1.0
// @description     This is a simple store management server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      simple-store-management-golang-production.up.railway.app
// @BasePath

// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	configs.Initiator()
	connection.Initiator()
	defer connection.SqlDBConnections.Close()
	migration.Initiator(connection.SqlDBConnections)
	routers.StartServer()
}
