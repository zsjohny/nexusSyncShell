package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"nexusSync/util"
	"os"
	"path"
	"strings"
	"sync"
)

const Nexus_operator = "/"

type NexusSyncService struct {
}

func (nexusSyncService *NexusSyncService) StartUpload(config *Config) {
	fmt.Println("traverse the dir")
	localPath := config.LocalDir
	files, err := util.GetAllFiles(localPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	usr, pwd, err := checkAuthValid(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	remoteUrl := config.RemoteUrl
	remoteDir := Nexus_operator
	if len(config.RemoteDir) > 0 {
		remoteDir = config.RemoteDir
	}

	//使用等待group等待所有coroutine完成
	var wg sync.WaitGroup
	wg.Add(len(files))

	for _, file := range files {
		go func(wg *sync.WaitGroup, file string) {
			err := util.NexusPost(remoteUrl, usr, pwd, file, remoteDir)
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

func checkAuthValid(config *Config) (string, string, error) {
	//验证信息
	auth := strings.Split(config.Auth, ":")
	if len(auth) < 1 {
		return "", "", fmt.Errorf("auth info error, plearse split with `:`")

	}
	usr := auth[0]
	pwd := auth[1]
	return usr, pwd, nil
}

func (*NexusSyncService) StartDownload(config *Config) {
	// valid data
	localDir := config.LocalDir
	if err := util.PathExists(localDir); err != nil {
		fmt.Println(err)
		return
	}
	remoteUrl := config.RemoteUrl
	usr, pwd, err := checkAuthValid(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	//list component
	resp, err := util.BasicAuthGet(remoteUrl, usr, pwd)
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
	var assets []Asset
	for _, cpn := range *components {
		if cpn.Group == config.RemoteDir {
			assets = append(assets, *cpn.Assets...)
		}
	}
	//download list to target path
	fmt.Println("ready to request remote to get files")
	var wg sync.WaitGroup
	wg.Add(len(assets))
	for idx := range assets {
		//fmt.Println("download asset info:", assets[idx])
		go func(wg *sync.WaitGroup, asset *Asset) {
			fileName := path.Base(asset.DownloadUrl)
			fileName = localDir + string(os.PathSeparator) + fileName
			fmt.Printf("ready to get fileName:%s\n", fileName)
			if resp, err := http.Get(asset.DownloadUrl); err == nil {
				var byteBuf []byte
				buffer := bytes.NewBuffer(byteBuf)
				_, err := io.Copy(buffer, resp.Body)
				defer resp.Body.Close()
				if err != nil {
					fmt.Printf("io copy fail when get file %s, err:%s\n", fileName, err)
					return
				}
				ioutil.WriteFile(fileName, buffer.Bytes(), 0777)
				fmt.Printf("receive file %s success \n", fileName)
			} else {
				fmt.Printf("file %s download fail\n", fileName)
			}
			wg.Done()
		}(&wg, &assets[idx])
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
