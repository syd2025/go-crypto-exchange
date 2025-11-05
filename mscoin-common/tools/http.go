package tools

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Post(url string, params any) ([]byte, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelFunc()
	marshal, _ := json.Marshal(params)
	s := string(marshal)
	reqbody := strings.NewReader(s)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, reqbody)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer httpResp.Body.Close()
	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func Get(url string, params any) ([]byte, error) {
	marshal, _ := json.Marshal(params)
	s := string(marshal)
	reqbody := strings.NewReader(s)
	httpReq, err := http.NewRequest(http.MethodGet, url, reqbody)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer httpResp.Body.Close()
	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func GetWithHeader(path string, m map[string]string, proxy string) ([]byte, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelFunc()
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range m {
		httpReq.Header.Add(k, v)
	}
	client := http.DefaultClient
	if proxy != "" {
		proxyAddress, _ := url.Parse(proxy)
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyAddress),
			},
		}
	}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil

}
