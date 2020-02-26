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

const nexusOperator = "/"
const (
	macIgnore          = ".DS_Store"
	uploadFormat       = "/service/rest/v1/components?repository="
	continuation_token = "&continuationToken="
	download_repo      = "/service/rest/v1/search/assets?repository="
	download_group     = "&group="
)

type NexusSyncService struct {
}
type downloadModel struct {
	tarPath     string
	downloadUrl string
}

func (nexusSyncService *NexusSyncService) StartUpload(config *Config) {
	url := config.RemoteUrl + uploadFormat + config.Repository
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
		if strings.Index(file.FilePath, macIgnore) != -1 {
			wg.Done()
			continue
		}
		go func(wg *sync.WaitGroup, file util.PostModel) {
			err := util.NexusPost(url, config.Usr, config.Pwd, file.FilePath, file.LevelInfo)
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
	url := config.RemoteUrl + download_repo + config.Repository + download_group + config.RemoteDir + "*"
	//list component
	assets := getAssets(url, config.Usr, config.Pwd)
	if len(assets) == 0 {
		fmt.Println("get files error when request for the file list")
		return
	}
	//查找匹配路径的文件
	var downloadModels []downloadModel
	remotePath := config.RemoteDir[1:]
	for _, asset := range assets {
		if strings.HasPrefix(asset.Path, remotePath) {
			tarPath := strings.Replace(asset.Path, remotePath, "", 1)
			// /4/智能小A.gif
			tarPath = util.FormatOsPath(tarPath)
			downloadModels = append(downloadModels, downloadModel{tarPath, asset.DownloadUrl})
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

func getAssets(remoteUrl, usr, pwd string) []Asset {
	var assets []Asset
	assets = getAssetsRes(remoteUrl, "", usr, pwd, assets)
	if len(assets) < 0 {
		return nil
	}
	return assets

}

func getAssetsRes(remoteUrl string, continuationToken string, usr string, pwd string, assets []Asset) []Asset {
	url := remoteUrl
	if len(continuationToken) > 0 {
		url = url + continuation_token + continuationToken
	}
	resp, err := util.BasicAuthGet(url, usr, pwd)
	if err != nil {
		fmt.Println(err)
		return assets

	}
	jsonStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read response error, err:%s", err)
	}
	body, err := json2Struct(string(jsonStr))
	if err != nil {
		fmt.Println(err)
		return assets
	}
	assets = append(assets, body.Items...)
	if len(body.ContinuationToken) <= 0 {
		return assets
	} else {
		return getAssetsRes(remoteUrl, body.ContinuationToken, usr, pwd, assets)
	}
}

func json2Struct(jsonStr string) (*Body, error) {
	var bodyJSON Body
	err := json.Unmarshal([]byte(jsonStr), &bodyJSON)
	if err != nil {
		return nil, fmt.Errorf("json2Struct error,err:%s\n", err)
	}
	return &bodyJSON, nil
}
