package util

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
)

const (
	Post      string = "POST"
	Get       string = "GET"
	Directory        = "raw.directory"
	Asset_n          = "raw.assetN"
	File_name        = "raw.assetN.filename"
)

//文件路径只兼容linux
func NewMutipartPostRequest(postUrl, filePath string, tarPath string) (req *http.Request, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// TODO 正式环境去除,兼容win
	filePath = strings.ReplaceAll(filePath, "\\", "/")
	fileName := path.Base(filePath)

	//创建一个模拟的form中的一个选项,这个form项现在是空的
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作, 设置文件的上传参数叫uploadfile, 文件名是filename,
	//相当于现在还没选择文件, form项里选择文件的选项
	fileWriter, err := bodyWriter.CreateFormFile(Asset_n, fileName)
	if err != nil {
		fmt.Println("error writing to buffer")
		return nil, err
	}

	//iocopy 这里相当于选择了文件,将文件放到form中
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, err
	}

	bodyWriter.WriteField(Directory, tarPath)
	bodyWriter.WriteField(File_name, path.Base(filePath))

	//这个很关键,必须这样写关闭,不能使用defer关闭,不然会导致错误
	bodyWriter.Close()

	request, err := http.NewRequest(Post, postUrl, bodyBuf)
	request.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	return request, nil
}

func NexusPost(postUrl, username, password, filePath, tarPath string) error {
	req, err := NewMutipartPostRequest(postUrl, filePath, tarPath)
	if err != nil {
		fmt.Printf("file %s generate request fail\n", filePath)
		return err
	}
	req.SetBasicAuth(username, password)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("file %s request fail\n", filePath)
		return err
	}
	if response.StatusCode != http.StatusNoContent {
		fmt.Printf("response status code error,code = %d, file : %s\n", response.StatusCode, filePath)
		return err
	}
	return nil
}

func BasicAuthGet(url string, username, password string) (*http.Response, error) {
	request, err := http.NewRequest(Get, url, nil)
	if err != nil {
		fmt.Println("get request generation fail")
		return nil, err
	}
	request.SetBasicAuth(username, password)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("get request fail，code= %d\n", resp.StatusCode)
		return nil, err
	}
	return resp, nil
}
