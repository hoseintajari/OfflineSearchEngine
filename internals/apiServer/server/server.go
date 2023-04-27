package server

import (
	"OfflineSearchEngine/internals/apiServer/CreateEngine"
	"OfflineSearchEngine/internals/apiServer/config"
	"OfflineSearchEngine/internals/apiServer/routing"
	"OfflineSearchEngine/internals/configuration"
	"OfflineSearchEngine/internals/fileReader"
	"github.com/gin-gonic/gin"
)

func RunServer(engine CreateEngine.SearchEngine, pathConfig string) {
	serverAddress, err := config.LoadConfig(pathConfig)
	if err != nil {
		panic(err)
	}
	s := gin.Default()
	err = fileReader.ListTextFiles(engine, configuration.PathReadDir)

	if err != nil {
		return
	}
	routing.Router(s)
	err = s.Run(serverAddress)
	if err != nil {
		panic("Failed to run server")
	}
}
