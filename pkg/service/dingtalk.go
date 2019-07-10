package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type Request struct {
	WebHook string `json:"-"`
	MsgType string `json:"msgtype"`
	Text    text   `json:"text"`
	At      At     `json:"at"`
}
type text struct {
	Content string `json:"content"`
}
type At struct {
	IsAtAll bool `json:"isAtAll"`
}

func NewRequest(webhook, subject, content string) *Request {
	return &Request{
		WebHook: webhook,
		MsgType: "text",
		Text: text{
			Content: subject + "\n\n" + content + "\n",
		},
		At: At{
			IsAtAll: true,
		},
	}
}
func (r *Request) Name() string {
	return "dingtalk"
}
func (r *Request) Send() (bool, error) {
	httpclient := http.DefaultClient
	stream, err := json.Marshal(r)
	if err != nil {
		return false, err
	}
	resp, err := httpclient.Post(r.WebHook, "application/json", bytes.NewReader(stream))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		stream, _ = ioutil.ReadAll(resp.Body)
		return false, errors.Wrap(err, "ding talk :: send")
	}
	return true, nil
}
