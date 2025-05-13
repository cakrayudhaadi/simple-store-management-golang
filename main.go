package main

import (
	"simple-store-management/configs"
	"simple-store-management/databases/connection"
	"simple-store-management/databases/migration"

	_ "github.com/lib/pq"
)

func main() {
	configs.Initiator()
	connection.Initiator()
	defer connection.SqlDBConnections.Close()
	migration.Initiator(connection.SqlDBConnections)
}
