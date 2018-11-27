package gutil

import (
	"encoding/json"
	"github.com/usthooz/gutil/http"
)

// SendMsgToDinTalk 发送消息到钉钉
func SendMsgToDinTalk(title, text, url string) {
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]string{
			"title": title,
			"text":  text,
		},
	}
	// 数据编码
	dataJson, err := json.Marshal(data)
	if err != nil {
		tp.Errorf("SendMsgToDinTalk: data marshal err-> %v", err)
		return
	}
	// 发送
	_, err = xhttp.PostJson(url, dataJson)
	if err != nil {
		tp.Errorf("SendMsgToDinTalk: Post json err-> %v", err)
	}
	return
}
