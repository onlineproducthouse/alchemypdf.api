package httprequtil

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"alchemypdf.api/lib/alchemypdf.api.util/errorlocal"
)

type (
	HTTPReqUtilCfg struct {
		Method      string
		URL         string
		ContentType string
		Payload     any // MUST be JSON
	}

	IHTTPReqUtil interface {
		SetHeader(key, value string)
		Execute() (*http.Response, []byte, error)
	}

	HTTPReqUtil struct {
		req *http.Request
	}
)

func NewHTTPReqUtil(cfg HTTPReqUtilCfg) (*HTTPReqUtil, *errorlocal.AppError) {
	const op = "NewHTTPReqUtil"

	var payload io.Reader = nil

	if cfg.Method == "POST" && cfg.Payload != nil {
		reqJson, reqJsonErr := json.Marshal(cfg.Payload)
		if reqJsonErr != nil {
			return nil, errorlocal.UnknownErr(reqJsonErr.Error(), op, reqJsonErr)
		}

		payload = bytes.NewBuffer(reqJson)
	}

	req, reqErr := http.NewRequest(cfg.Method, cfg.URL, payload)
	if reqErr != nil {
		return nil, errorlocal.UnknownErr(reqErr.Error(), op, reqErr)
	}

	if cfg.ContentType == "" {
		req.Header.Set("content-type", "application/json")
	} else {
		req.Header.Set("content-type", cfg.ContentType)
	}

	return &HTTPReqUtil{req}, nil
}

func (svc HTTPReqUtil) SetHeader(key, value string) {
	svc.req.Header.Set(key, value)
}

func (svc HTTPReqUtil) Execute() (*http.Response, []byte, *errorlocal.AppError) {
	const op = "HTTPReqUtil.Execute"

	client := &http.Client{Timeout: 0}

	executeResp, executeRespErr := client.Do(svc.req)
	if executeRespErr != nil {
		return nil, nil, errorlocal.UnknownErr(executeRespErr.Error(), op, executeRespErr)
	}

	defer executeResp.Body.Close()

	respBody, respBodyErr := io.ReadAll(executeResp.Body)
	if respBodyErr != nil {
		return nil, nil, errorlocal.UnknownErr(respBodyErr.Error(), op, respBodyErr)
	}

	return executeResp, respBody, nil
}
