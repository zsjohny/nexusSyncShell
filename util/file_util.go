package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

type PostModel struct {
	FilePath  string
	LevelInfo string
}

func PathExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		return fmt.Errorf("file doesn't exists or error isn't be defined in os, path = %s", path)
	}
	return fmt.Errorf("file doesn't exists ,path = %s", path)
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

func GetAllFiles(pathname string, postModels []PostModel, basePath string) ([]PostModel, error) {
	sep := string(os.PathSeparator)
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return postModels, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + sep + fi.Name()
			postModels, err = GetAllFiles(fullDir, postModels, basePath+fi.Name()+"/")
			if err != nil {
				fmt.Println("read dir fail:", err)
				return postModels, err
			}
		} else {
			fullName := pathname + sep + fi.Name()
			postModels = append(postModels, PostModel{fullName, basePath})
		}
	}
	return postModels, nil
}
