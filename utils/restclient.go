package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Post request type
const (
	REQUEST_BODY_JSON  string = "JSON"
	REQUEST_BODY_XFORM string = "X_WWW_FORM"
)

// DoRequest Create http request
func DoRequest(method string, urlStr string, headers map[string]string, body map[string]interface{}, typeBody string) (int, []byte) {
	var err error
	var reqBody io.Reader = nil

	if headers == nil {
		headers = make(map[string]string)
	}

	if body != nil {
		reqBody = buildRequestBody(body, typeBody, headers)
	}

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest(method, urlStr, reqBody)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if err != nil {
		ShowErrorLogs(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		ShowErrorLogs(err)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		ShowErrorLogs(err)
	}

	return resp.StatusCode, respBody
}

// DoGet sends a GET request with params encoded as query string
func DoGet(urlStr string, headers map[string]string, params map[string]interface{}) (int, []byte) {
	if params != nil {
		urlStr = appendQueryParams(urlStr, params)
	}
	return DoRequest(http.MethodGet, urlStr, headers, nil, "")
}

// DoUpdate sends a POST request with JSON body
func DoUpdate(method string, urlStr string, headers map[string]string, params map[string]interface{}, body map[string]interface{}, typeBody string) (int, []byte) {
	if params != nil {
		urlStr = appendQueryParams(urlStr, params)
	}
	return DoRequest(method, urlStr, headers, body, typeBody)
}

// appendQueryParams encodes params into the query string of urlStr.
func appendQueryParams(urlStr string, params map[string]interface{}) string {
	u, parseErr := url.Parse(urlStr)
	if parseErr != nil {
		ShowErrorLogs(parseErr)
		return urlStr
	}
	q := u.Query()
	for key, element := range params {
		switch v := element.(type) {
		case []string:
			for _, item := range v {
				q.Add(key, item)
			}
		default:
			q.Set(key, fmt.Sprintf("%v", v))
		}
	}
	u.RawQuery = q.Encode()
	return u.String()
}

// buildRequestBody marshals body according to typeBody and sets appropriate headers.
func buildRequestBody(body map[string]interface{}, typeBody string, headers map[string]string) io.Reader {
	if body == nil {
		return nil
	}
	if typeBody == REQUEST_BODY_JSON {
		headers["Content-Type"] = "application/json"
		jBody, err := json.Marshal(body)
		if err != nil {
			ShowErrorLogs(err)
			return nil
		}
		return bytes.NewBuffer(jBody)
	}
	if typeBody == REQUEST_BODY_XFORM {
		headers["Content-Type"] = "application/x-www-form-urlencoded"
		var payload string
		for key, element := range body {
			payload = payload + fmt.Sprintf("%s=%s&", key, element)
		}
		payload = strings.TrimRight(payload, "&")
		return strings.NewReader(payload)
	}
	return nil
}
