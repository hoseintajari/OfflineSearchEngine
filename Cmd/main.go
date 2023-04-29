package main

import (
	"OfflineSearchEngine/internals/apiServer"
	"OfflineSearchEngine/internals/apiServer/CreateEngine"
	"OfflineSearchEngine/internals/dataBase/DBmodels"
)

func main() {
	DBmodels.CreateUserTable()
	engine := CreateEngine.NewSearchEngine("v1", 100)
	apiServer.RunServer(engine, "./Cmd")
}
