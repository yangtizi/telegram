package tg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SendMessage(strBotSecret string, nChatID int64, strText string) {
	// 5326579020:AAF4obMIPDujmzK73puieTcSBfk4JTlX1YI

	strURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s",
		strBotSecret, nChatID, url.QueryEscape(strText),
	)
	go Bytes(strURL)
}

func SendMessageMarkdownV2(strBotSecret string, nChatID int64, strText string) {
	// 5326579020:AAF4obMIPDujmzK73puieTcSBfk4JTlX1YI

	strURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&parse_mode=MarkdownV2&text=%s",
		strBotSecret, nChatID, url.QueryEscape(strText),
	)
	go Bytes(strURL)
}

func Bytes(strURL string) ([]byte, error) {
	resp, err := http.Get(strURL)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return nil, err
	}

	if resp.StatusCode != 200 {
		return body, errors.New(resp.Status)
	}

	return body, nil
}
