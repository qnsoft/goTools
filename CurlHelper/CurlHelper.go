package Curl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpRequest(url, method string, data []byte) (resp *http.Response, err error) {
	client := &http.Client{}
	request, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		fmt.Println("err:", err)
	}
	request.Header.Set("Content-Type", "application/xjson")
	request.Header.Set("Connection", "Keep-Alive")
	resp, err = client.Do(request)
	return resp, err
}

func RequestResult(url, method string, data []byte) (string, error) {
	response, err := HttpRequest(url, method, data)
	if err != nil {
		return "", err
	}
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
