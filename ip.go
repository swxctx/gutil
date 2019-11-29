package gutil

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetLocalPublicIp get external IP addr.
// NOTE: Query IP information from the service API: http://pv.sohu.com/cityjson?ie=utf-8
func GetLocalPublicIp() (ip string, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("Get external IP error: %v", p)
		} else if err != nil {
			err = errors.New("Get external IP error: " + err.Error())
		}
	}()
	resp, err := http.Get("http://pv.sohu.com/cityjson?ie=utf-8")
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	idx := bytes.Index(b, []byte(`"cip": "`))
	b = b[idx+len(`"cip": "`):]
	idx = bytes.Index(b, []byte(`"`))
	b = b[:idx]
	ip = string(b)
	return
}
