package main

import (
	"OfflineSearchEngine/internals/apiServer/CreateEngine"
	"OfflineSearchEngine/internals/apiServer/server"
	"OfflineSearchEngine/internals/configuration"
	"OfflineSearchEngine/internals/dataBase/DBmodels"
)

func main() {
	DBmodels.CreateUserTable()
	str := "v1"
	configuration.MainEngine = CreateEngine.NewSearchEngine(str, 1000)
	server.RunServer(configuration.MainEngine, "./Cmd")
}
