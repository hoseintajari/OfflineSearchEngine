package apiServer

import (
	"OfflineSearchEngine/internals/apiServer/CreateEngine"
	"OfflineSearchEngine/internals/apiServer/config"
	server2 "OfflineSearchEngine/internals/apiServer/server"
	"OfflineSearchEngine/internals/fileReader"
	"OfflineSearchEngine/internals/linguisticModule"
)

func RunServer(engine CreateEngine.SearchEngine, pathConfig string) {
	serverAddress, err := config.LoadConfig(pathConfig)
	if err != nil {
		panic(err)
	}

	li := linguisticModule.NewLinguisticModule(
		&linguisticModule.ToLower{},
		&linguisticModule.PunctuationRemover{},
		&linguisticModule.StopWords{},
	)
	iEngine := CreateEngine.CreateISearchEngine(engine, pathConfig, li)
	server := server2.CreateServer(&iEngine)
	if err != nil {
		return
	}

	if err = fileReader.ListTextFiles(server, li); err != nil {
		panic(err)
	}
	if err = server.Run(serverAddress); err != nil {
		panic(err)
	}
	if err != nil {
		panic("Failed to run server")
	}
}
