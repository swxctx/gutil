package xhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Get
func Get(path string) (resp *http.Response, bs []byte, err error) {
	resp, err = http.Get(path)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("http get error-> status = %d", resp.StatusCode)
		return
	}
	bs, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

// GetJSONByParams
func GetJSONByParams(path string, params map[string]string) (resp *http.Response, data map[string]interface{}, err error) {
	params_s := ""
	for k, v := range params {
		params_s += fmt.Sprintf("%s=%s&", k, QueryEscape(v))
	}
	params_s = params_s[:len(params_s)-1]
	return GetJSON(path + "?" + params_s)
}

// JoinParams
func JoinParams(path string, params map[string]string) string {
	params_s := ""
	for k, v := range params {
		params_s += fmt.Sprintf("%s=%s&", k, QueryEscape(v))
	}
	params_s = params_s[:len(params_s)-1]
	return path + "?" + params_s
}

// GetJSON
func GetJSON(path string) (resp *http.Response, data map[string]interface{}, err error) {
	resp, err = http.Get(path)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		bs, _ := ioutil.ReadAll(resp.Body)
		data = map[string]interface{}{}
		json.Unmarshal(bs, &data)
		err = fmt.Errorf("http get error-> status = %d, resp = %v", resp.StatusCode, string(bs))
		return
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	data = map[string]interface{}{}
	err = json.Unmarshal(bs, &data)
	if err != nil {
		return
	}
	return
}

// Post
func Post(path string, form url.Values) (resp *http.Response, bs []byte, err error) {
	resp, err = http.PostForm(path, form)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bs, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("http post json error-> status = %d", resp.StatusCode)
		return
	}
	return
}

// PostJSON
func PostJSON(URL string, jsonByte []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonByte))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// QueryEscape
func QueryEscape(s string) string {
	s = url.QueryEscape(s)
	s = strings.Replace(s, "+", "%20", -1)
	return s
}

// QueryUnescape
func QueryUnescape(s string) (string, error) {
	var err error
	s = strings.Replace(s, "%20", "+", -1)
	s, err = url.QueryUnescape(s)
	return s, err
}
