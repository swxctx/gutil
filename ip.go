package gutil

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
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

// ExtranetIP get external IP addr.
func ExtranetIP() (ip string, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("Get external IP error: %v", p)
		} else if err != nil {
			err = errors.New("Get external IP error: " + err.Error())
		}
	}()
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	ip = string(b)
	return
}

// ParseIpToUint32 ip转换为uint32
func ParseIpToUint32(s string) (uint32, error) {
	var (
		ip uint32
	)
	ipObj := net.ParseIP(s)
	if ipObj == nil {
		return ip, errors.New("ip parse failed")
	}
	ip |= uint32(ipObj[12]) << 24
	ip |= uint32(ipObj[13]) << 16
	ip |= uint32(ipObj[14]) << 8
	ip |= uint32(ipObj[15])
	return ip, nil
}

// GetLocalInnerIP get host local inner ip
func GetLocalInnerIP() (string, error) {
	address, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range address {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}
	return "", nil
}
