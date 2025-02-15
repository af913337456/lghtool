package LghTool

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func SimpleHttpGET(urlStr string, data interface{}, setHeader func(header *http.Header)) (err error) {
	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		err = fmt.Errorf("http.NewRequest err: %v\n", err)
		return
	}

	defaultHttpClient := &http.Client{}
	defaultHttpClient.Timeout = 10 * time.Second

	// 配置请求头
	if setHeader != nil {
		setHeader(&req.Header)
	}
	resp, err := defaultHttpClient.Do(req)
	if err != nil {
		err = fmt.Errorf("defaultHttpClient.Do err: %v\n", err)
		return
	}
	// Do something with the response
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
		if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
			err = fmt.Errorf("json.NewDecoder err: %v\n", err)
			return
		}
	}
	return
}

func SimpleHttpPOST(urlStr string, data interface{}, setHeader func(header *http.Header)) (err error) {
	req, err := http.NewRequest(http.MethodPost, urlStr, nil)
	if err != nil {
		err = fmt.Errorf("http.NewRequest err: %v\n", err)
		return
	}

	defaultHttpClient := &http.Client{}
	defaultHttpClient.Timeout = 10 * time.Second

	// 配置请求头
	if setHeader != nil {
		setHeader(&req.Header)
	}
	resp, err := defaultHttpClient.Do(req)
	if err != nil {
		err = fmt.Errorf("defaultHttpClient.Do err: %v\n", err)
		return
	}
	// Do something with the response
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
		if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
			err = fmt.Errorf("json.NewDecoder err: %v\n", err)
			return
		}
	}
	return
}
