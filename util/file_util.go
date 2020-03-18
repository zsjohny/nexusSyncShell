package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const nexusSep = "/"

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

//c:\test /data/a basePath - > /
func GetAllFiles(pathname string, postModels []PostModel, basePath string) ([]PostModel, error) {
	//判断所给路径是否为文件
	if IsFile(pathname){
		postModels = append(postModels, PostModel{pathname, basePath})
		return postModels, nil
	}
	sep := string(os.PathSeparator)
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return postModels, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + sep + fi.Name()
			postModels, err = GetAllFiles(fullDir, postModels, basePath+nexusSep+fi.Name())
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

//调用os.MkdirAll递归创建文件夹
func CreateDirs(filePath string) error {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

func MkDirs(tempDir string) error {

	_, err := os.Stat(tempDir)
	if err != nil {
		fmt.Println("stat temp dir error,maybe is not exist, maybe not")
		if os.IsNotExist(err) {

			err := os.Mkdir(tempDir, os.ModePerm)
			if err != nil {
				return fmt.Errorf("mkdir failed![%v]\n", err)
			}
			return fmt.Errorf("temp dir is not exist")
		}

		return fmt.Errorf("stat file error")
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//standard format : os -> c:\test linux -> /data/soft
func FormatLocalPath(tarPath string) string {
	tarPath = FormatOsPath(tarPath)
	length := len(tarPath)
	if tarPath[length-1] == os.PathSeparator {
		tarPath = tarPath[:length-1]
	}
	return tarPath
}

//standard format : /lsj_test
func FormatNexusPathSeparator(path string) string {
	length := len(path)
	if length != 1 && path[length-1] == '/' {
		path = path[:length-1]
	}
	return strings.ReplaceAll(path, "\\", "/")
}

func FormatOsPath(path string) string {
	if os.PathSeparator == '/' {
		return strings.ReplaceAll(path, "\\", "/")
	} else {
		return strings.ReplaceAll(path, "/", "\\")
	}
}
