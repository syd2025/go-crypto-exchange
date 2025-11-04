package tools

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func Post(url string, params any) ([]byte, error) {
	marshal, _ := json.Marshal(params)
	s := string(marshal)
	reqbody := strings.NewReader(s)
	httpReq, err := http.NewRequest(http.MethodPost, url, reqbody)
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
