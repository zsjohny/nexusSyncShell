package core

import (
	"bytes"
	"encoding/json"
	"errors"
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
	fmt.Println("start upload")
	fmt.Println("traverse the dir")
	localPath := config.LocalDir
	files, err := util.GetAllFiles(localPath)
	if err != nil {
		fmt.Errorf("read localDir fail")
		return
	}

	usr, pwd, err := checkAuthValid(config)
	if err != nil {
		fmt.Errorf(err.Error())
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
				fmt.Printf("file %s upload fail\n", file)
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
		return "", "", errors.New("auth info error, plearse split with `:`")

	}
	usr := auth[0]
	pwd := auth[1]
	return usr, pwd, nil
}

func (*NexusSyncService) StartDownload(config *Config) {
	fmt.Println("start download")
	// valid data
	localDir := config.LocalDir
	if err := util.PathExists(localDir); err != nil {
		fmt.Printf("localDir error, err : %s\n", err)
		return
	}
	remoteUrl := config.RemoteUrl
	usr, pwd, err := checkAuthValid(config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//list component
	resp, err := util.BasicAuthGet(remoteUrl, usr, pwd)
	if err != nil {
		return
	}
	jsonStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read response error")
		return
	}
	components := json2Struct(string(jsonStr))
	if components == nil {
		return
	}
	//find correct path
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
		fmt.Println(assets[idx])
		go func(wg *sync.WaitGroup, asset *Asset) {
			fileName := path.Base(asset.DownloadUrl)
			fileName = localDir + string(os.PathSeparator) + fileName
			fmt.Printf("ready to get fileName:%s", fileName)
			if resp, err := http.Get(asset.DownloadUrl); err == nil {
				buffer := bytes.NewBuffer(make([]byte, 4096))
				_, err := io.Copy(buffer, resp.Body)
				defer resp.Body.Close()
				if err != nil {
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

func json2Struct(jsonStr string) *[]Component {
	var bodyJSON Body
	err := json.Unmarshal([]byte(jsonStr), &bodyJSON)
	if err != nil {
		fmt.Printf("json2Struct error,err:%s\n", err)
		return nil
	}
	return &bodyJSON.Items
}
