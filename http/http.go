package xhttp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

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

func Post(path string, form url.Values) (resp *http.Response, data map[string]interface{}, err error) {
	resp, err = http.PostForm(path, form)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("http post json error-> status = %d", resp.StatusCode)
		return
	}
	bs, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

func PostJSON(path string, form url.Values) (resp *http.Response, data map[string]interface{}, err error) {
	resp, err = http.PostForm(path, form)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("http post json error-> status = %d", resp.StatusCode)
		return
	}
	data = map[string]interface{}{}
	err = json.Unmarshal(bs, &data)
	if err != nil {
		return
	}
	return
}
