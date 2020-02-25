package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"nexusSync/util"
	"path/filepath"
	"strings"
	"sync"
)

const Nexus_operator = "/"

type NexusSyncService struct {
}
type downloadModel struct {
	tarPath     string
	downloadUrl string
}

func (nexusSyncService *NexusSyncService) StartUpload(config *Config) {
	fmt.Println("traverse the dir")
	localPath := config.LocalDir

	var postModels []util.PostModel
	files, err := util.GetAllFiles(localPath, postModels, config.RemoteDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	//使用等待group等待所有coroutine完成
	var wg sync.WaitGroup
	wg.Add(len(files))

	for _, file := range files {
		go func(wg *sync.WaitGroup, file util.PostModel) {
			err := util.NexusPost(config.RemoteUrl, config.Usr, config.Pwd, file.FilePath, file.LevelInfo)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("file %s upload success\n", file)
			}
			wg.Done()
		}(&wg, file)
	}

	wg.Wait()

}

func (*NexusSyncService) StartDownload(config *Config) {
	//list component
	resp, err := util.BasicAuthGet(config.RemoteUrl, config.Usr, config.Pwd)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read response error, err:%s", err)
		return
	}
	components, err := json2Struct(string(jsonStr))
	if err != nil {
		fmt.Println(err)
		return
	}
	//查找匹配路径的文件
	var downloadModels []downloadModel
	remoteDir := config.RemoteDir
	remotePath := remoteDir[1:]
	for _, cpn := range *components {
		for _, asset := range *cpn.Assets {
			if strings.HasPrefix(asset.Path, remotePath) {
				tarPath := strings.Replace(asset.Path, remotePath, "", 1)
				// /4/智能小A.gif
				tarPath = util.FormatOsPath(tarPath)
				downloadModels = append(downloadModels, downloadModel{tarPath, asset.DownloadUrl})
			}

		}
	}
	//download list to target path
	fmt.Println("ready to request remote to get files")
	var wg sync.WaitGroup
	wg.Add(len(downloadModels))
	for idx := range downloadModels {
		//fmt.Println("download asset info:", assets[idx])
		go func(wg *sync.WaitGroup, downloadModel *downloadModel) {
			filePath := config.LocalDir + downloadModel.tarPath
			fmt.Printf("ready to get fileName:%s\n", filePath)
			if resp, err := http.Get(downloadModel.downloadUrl); err == nil {
				var byteBuf []byte
				buffer := bytes.NewBuffer(byteBuf)
				_, err := io.Copy(buffer, resp.Body)
				defer resp.Body.Close()
				if err != nil {
					fmt.Printf("io copy fail when get file %s, err:%s\n", filePath, err)
					return
				}
				err = util.CreateDirs(filepath.Dir(filePath))
				if err != nil {
					fmt.Printf("make dir error %s\n", filePath)
				}

				err = ioutil.WriteFile(filePath, buffer.Bytes(), 0777)
				if err != nil {
					fmt.Printf("write the file %s fail\n", filePath)
				}
				fmt.Printf("receive file %s success \n", filePath)
			} else {
				fmt.Printf("file %s download fail\n", filePath)
			}
			wg.Done()
		}(&wg, &downloadModels[idx])
	}

	wg.Wait()

}

func json2Struct(jsonStr string) (*[]Component, error) {
	var bodyJSON Body
	err := json.Unmarshal([]byte(jsonStr), &bodyJSON)
	if err != nil {
		return nil, fmt.Errorf("json2Struct error,err:%s\n", err)
	}
	return &bodyJSON.Items, nil
}
