package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/andrew-j-price/journey/logger"
)

func RestPerformGetUrl(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error.Println(err)
		return nil
	}
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
		logger.Error.Println(err)
		return nil
	}
	// defer resp.Body.Close()
	// resp.Body.Close()
	return resp
}

func RestResponseBody(resp *http.Response) []byte {
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error.Println(err)
		return nil
	}
	// logger.Debug.Printf("respData: is of type: %v, with value: %v\n", reflect.TypeOf(respData), respData)
	// NOTE: single line alternative
	// respData, _ := ioutil.ReadAll(resp.Body)
	return respData
}

// Get specific Response Header
func RestResponseGetHeader(resp *http.Response, name string) (string, error) {
	value := resp.Header.Get(name)
	logger.Debug.Printf("Value is %v\n", value)
	if value == "" {
		return "", errors.New("empty response")
	} else {
		return value, nil
	}
}

// Get HTTP response Status Code
func RestResponseCodeAnalysis(resp *http.Response, expectedReturnCode int) bool {
	if resp.StatusCode != expectedReturnCode {
		logger.Error.Printf("Expected %v, received %v\n", expectedReturnCode, resp.StatusCode)
		return false
	} else {
		return true
	}
}

// Get HTTP response Status Code
func RestResponseReturnCode(resp *http.Response) int {
	return resp.StatusCode
}

func RestJsonMarshalData(payload map[string]interface{}) []byte {
	data, err := json.Marshal(payload)
	if err != nil {
		logger.Error.Println(err)
		return nil
	}
	return data
}

func RestJsonUnmarshalData(response []byte) map[string]interface{} {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(response), &data)
	if err != nil {
		logger.Error.Println(err)
		return nil
	}
	return data
}

// Takes URL and returns JSON
// Params (string) of URL
// Params (int) of expected HTTP Response Code
// Returns (map[string]interface{}) of JSON response
func RestGetUrl(url string, expectedReturnCode int) map[string]interface{} {
	logger.Info.Printf("Performing Get call to: %v\n", url)
	resp := RestPerformGetUrl(url)
	if !RestResponseCodeAnalysis(resp, expectedReturnCode) {
		return nil
	} else {
		respData := RestResponseBody(resp)
		respPayload := RestJsonUnmarshalData(respData)
		return respPayload
	}
}

func RestPostUrl(url string, payload map[string]interface{}, expectedReturnCode int) map[string]interface{} {
	logger.Info.Printf("Performing Post call to: %v with data: %v\n", url, payload)
	jsonData := RestJsonMarshalData(payload)
	resp := RestPerformPostUrl(url, jsonData)
	if !RestResponseCodeAnalysis(resp, expectedReturnCode) {
		return nil
	} else {
		respData := RestResponseBody(resp)
		respPayload := RestJsonUnmarshalData(respData)
		return respPayload
	}
}
