package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func makeRequest(togen string) (string, error) {
	toSend := &GPT3Data{
		Text: togen,
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(toSend)
	log.Println(payloadBuf)
	resp, err := http.Post("https://api.aicloud.sbercloud.ru/public/v1/public_inference/gpt3/predict", "application/json", payloadBuf)

	if err != nil {
		log.Print(err)
		return "", err
	} else if resp.Status != "200 OK" {
		return "", errors.New(resp.Status)
	}

	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	log.Println(resp.Status)
	log.Println(string(respBody))
	return string(respBody), nil
}

func parseAPIAnswer(ans []byte) GPT3Answer {
	var answer GPT3Answer
	json.Unmarshal(ans, &answer)
	return answer
}
