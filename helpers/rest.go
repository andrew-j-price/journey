package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/andrew-j-price/journey/logger"
)

func RestPerformGetUrl(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// defer resp.Body.Close()
	logger.Info.Printf("Reponse %v with return code: %v\n", http.StatusText(resp.StatusCode), resp.StatusCode)
	return resp
}

func RestPerformPostUrl(url string, jsonData []byte) *http.Response {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	// TODO: figure out method for setting custom headers more dynamically
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)

	body := RestHttpResponseToByte(resp)
	fmt.Println("Response Body:", string(body))
	return resp
}

func RestHttpResponseToByte(resp *http.Response) []byte {
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// logger.Debug.Printf("respData: is of type: %v, with value: %v\n", reflect.TypeOf(respData), respData)
	// NOTE: single line alternative
	// respData, _ := ioutil.ReadAll(resp.Body)
	return respData
}

func RestJsonMarshalData(payload map[string]interface{}) []byte {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	return jsonData
}

func RestJsonUnmarshalData(respData []byte) map[string]interface{} {
	var payload map[string]interface{}
	err := json.Unmarshal([]byte(respData), &payload)
	if err != nil {
		panic(err)
	}
	return payload
}

// Takes URL and returns JSON
// Params (string) of URL
// Returns (map[string]interface{}) of JSON response
func RestGetUrl(url string) map[string]interface{} {
	logger.Info.Printf("Performing Get call to: %v\n", url)
	resp := RestPerformGetUrl(url)
	respData := RestHttpResponseToByte(resp)
	payload := RestJsonUnmarshalData(respData)
	return payload
}

func RestPostUrl(url string, payload map[string]interface{}) map[string]interface{} {
	logger.Info.Printf("Performing Post call to: %v with data: %v\n", url, payload)
	jsonData := RestJsonMarshalData(payload)
	resp := RestPerformPostUrl(url, jsonData)
	respData := RestHttpResponseToByte(resp)
	respPayload := RestJsonUnmarshalData(respData)
	return respPayload
}
