package fileReader

import (
	"OfflineSearchEngine/internals/apiServer/CreateEngine"
	"OfflineSearchEngine/internals/configuration"
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

func ListTextFiles(engine CreateEngine.SearchEngine, path string) error {
	m := make(map[int]string)
	filesPath := ReadDir(path)
	for _, v := range filesPath {
		file, err := ReadFile(v)
		if err != nil {
			return err
		}
		configuration.DocId++
		m[configuration.DocId] = v
		engine.AddDoc(file, configuration.DocId)
	}
	return nil
}
