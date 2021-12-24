package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Request struct {
	HttpMethod      string            `json:"httpMethod"`
	Header          map[string]string `json:"headers"`
	Body            string            `json:"body"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
}

type ResponseBody struct {
	Method string `json:"method"`
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type Response struct {
	StatusCode      int               `json:"statusCode"`
	Header          map[string]string `json:"headers"`
	Body            string            `json:"body"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
}

func Handler(request []byte) ([]byte, error) {

	req := Request{}

	err := json.Unmarshal(request, &req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	upd := Update{}

	err = json.Unmarshal([]byte(req.Body), &upd)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var respText string
	var respChatID int64

	if upd.Message != nil {

		respText, err = ProcMessage(upd.Message)
		respChatID = upd.Message.Chat.ID

	} else if upd.EditedMessage != nil {

		respText, err = ProcMessage(upd.EditedMessage)
		respChatID = upd.EditedMessage.Chat.ID

	} else {
		err = errors.New("empty Message and EditedMessage")

	}

	if err != nil {

		respText = err.Error()
		err = nil

	}

	respbody, err := json.Marshal(&ResponseBody{
		Method: "sendMessage",
		ChatID: respChatID,
		Text:   respText,
	},
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resp := Response{
		StatusCode: 200,
		Header: map[string]string{
			"Content-Type": "application/json",
		},
		Body:            string(respbody),
		IsBase64Encoded: false,
	}

	respbytesl, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return respbytesl, nil
}
