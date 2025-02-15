package LghTool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type UploadFileResp struct {
	Code    int    `json:"code"`
	Status  int    `json:"status"`
	Data    string `json:"data"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

// filePath, File local path; fileName, Save to cloudstore
func UploadFormFile(url string, fileName, filePath string) (*UploadFileResp, error) {
	fileName = strings.Replace(fileName, ">", "", -1)
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file %s: %w", filePath, err)
	}
	defer file.Close()

	// 创建一个缓冲区来存储我们的表单数据
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 创建一个新的文件字段
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("could not create form file: %w", err)
	}

	// 将文件内容复制到 part
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("could not copy file to form: %w", err)
	}
	if err := writer.WriteField("FileName", fileName); err != nil {
		fmt.Println("Error adding form field:", err)
		return nil, err
	}

	// 关闭 multipart 编码器以写入尾部边界
	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("could not close writer: %w", err)
	}

	// 创建一个新的 HTTP 请求
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "text/plain")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	// 输出响应
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", fileName, string(respBody))
	respData := UploadFileResp{}
	_ = json.Unmarshal(respBody, &respData)
	return &respData, nil
}
