package fileReader

import (
	"OfflineSearchEngine/internals/apiServer/server"
	"OfflineSearchEngine/internals/linguisticModule"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadDir(path string) []string {
	var Res []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".txt") {
			Res = append(Res, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return Res
}

func ReadFile(path string) (*bufio.Scanner, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	return scanner, err
}

func ListTextFiles(server *server.Server, li linguisticModule.LinguisticModule) error {
	m := make(map[int]string)
	filesPath := ReadDir(server.Engine.PathReadDir)
	for _, v := range filesPath {
		file, err := ReadFile(v)
		if err != nil {
			return err
		}
		server.Engine.DocId++
		m[server.Engine.DocId] = v
		server.Engine.Engine.AddDoc(bufConvertString(file), server.Engine.DocId, li)
	}
	return nil
}
