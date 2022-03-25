package random

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"syreclabs.com/go/faker"
)

func getUrl(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("resp: is of type: %v, with value: %v\n", reflect.TypeOf(resp), resp)
	fmt.Printf("# URL %v return code: %v %v\n", url, resp.StatusCode, http.StatusText(resp.StatusCode))
	return resp
}

func getRespData(resp *http.Response) []byte {
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("respData: is of type: %v, with value: %v\n", reflect.TypeOf(respData), respData)
	return respData
}

func getAttestDefault() {
	fmt.Println("# FUNC: getAttestDefault")
	// define the data structure
	type attestDefault struct {
		Date          string `json:"date"`
		HostIP        string `json:"host_ip"`
		HostName      string `json:"host_name"`
		RemoteIP      string `json:"remote_ip"`
		ServerAddress string `json:"server_address"`
		URI           string `json:"uri"`
		UUID          string `json:"uuid"`
	}
	// get data
	resp := getUrl("http://attest.linecas.com/default")
	respData := getRespData(resp)
	// structure data
	var respObject attestDefault
	json.Unmarshal(respData, &respObject)
	// do something with data
	// fmt.Printf("respObject: is of type: %v, with value: %v\n", reflect.TypeOf(respObject), respObject)
	fmt.Printf("URI called: %v at: %v\n", respObject.URI, respObject.Date)
}

func getAttestDefaultWithoutStructs() {
	fmt.Println("# FUNC: getAttestDefaultWithoutStructs")
	// get data
	resp := getUrl("http://attest.linecas.com/default")
	respData := getRespData(resp)
	// structure data
	var data map[string]interface{}
	err := json.Unmarshal([]byte(respData), &data)
	if err != nil {
		panic(err)
	}
	// do something with data
	fmt.Printf("data: is of type: %v, with value: %v\n", reflect.TypeOf(data), data)
	// value would be <nil> if desired element did not exist
	fmt.Printf("host_name is: %v, uuid is: %v", data["host_name"], data["uuid"])
}

func getAttestDogs() {
	fmt.Println("# FUNC: getAttestDogs")
	// define the data structure
	type attestDogs []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	// get data
	resp := getUrl("http://attest.linecas.com/dogs")
	respData := getRespData(resp)
	// structure data
	var respObject attestDogs
	json.Unmarshal(respData, &respObject)
	// do something with data
	// fmt.Printf("respObject: is of type: %v, with value: %v\n", reflect.TypeOf(respObject), respObject)
	fmt.Printf("Index 0 Dog: %v with name: %v\n", respObject[0], respObject[0].Name)
}

func getAttestTodos() {
	fmt.Println("# FUNC: getAttestTodos")
	// define the data structure
	type attestTodos []struct {
		UID    string `json:"uid"`
		Task   string `json:"task"`
		Status bool   `json:"status"`
	}
	// get data
	resp := getUrl("http://attest.linecas.com/todos")
	respData := getRespData(resp)
	// structure data
	var respObject attestTodos
	json.Unmarshal(respData, &respObject)
	// do something with data
	// fmt.Printf("respObject: is of type: %v, with value: %v\n", reflect.TypeOf(respObject), respObject)
	completedTodos := 0
	for i := 0; i < len(respObject); i++ {
		if respObject[i].Status {
			completedTodos++
		}
	}
	fmt.Printf("There are: %v todos: with %v completed and %v open.\n", len(respObject), completedTodos, len(respObject)-completedTodos)
}

func postAttestTodoByteSlice(task_msg string) {
	fmt.Println("# FUNC: postAttestTodoByteSlice")
	toUrl := "http://attest.linecas.com/todos"
	fmt.Println("Sending POST request to:", toUrl)

	var jsonData = []byte(fmt.Sprintf(`{
		"task": "%v"
	}`, task_msg))

	req, _ := http.NewRequest("POST", toUrl, bytes.NewBuffer(jsonData))
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
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))
}

// NOTE: Preferring this method, as I find it simpler and more straightforward.
func postAttestTodoMapInterface(task_msg string) {
	fmt.Println("# FUNC: postAttestTodoMapInterface")
	toUrl := "http://attest.linecas.com/todos"
	fmt.Println("Sending POST request to:", toUrl)

	data := map[string]interface{}{
		"task": task_msg,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", toUrl, bytes.NewBuffer(jsonData))
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
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))
}

func postAttestTodoStruct(task_msg string) {
	fmt.Println("# FUNC: postAttestTodoStruct")
	toUrl := "http://attest.linecas.com/todos"
	fmt.Println("Sending POST request to:", toUrl)

	type jsonDataStruct struct {
		Task string `json:"task"`
	}
	jsonPayload := &jsonDataStruct{Task: task_msg}
	jsonStr, err := json.Marshal(jsonPayload)
	if err != nil {
		log.Fatalf("could not marshal JSON: %s", err)
	}

	req, _ := http.NewRequest("POST", toUrl, bytes.NewBuffer(jsonStr))
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
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))
}

func AttestMain() {
	fmt.Println("# FUNC: AttestMain")
	if false {
		// This will obviously not run
		// but prevents fmt from removing imports when commenting out functions
		getAttestDefault()
		getAttestDefaultWithoutStructs()
		getAttestDogs()
		getAttestTodos()
		task_msg := faker.Hacker().SaySomethingSmart()
		postAttestTodoByteSlice(task_msg)
		postAttestTodoMapInterface(task_msg)
		postAttestTodoStruct(task_msg)
	}
	// actual items to run are here
	getAttestDefaultWithoutStructs()
	task_msg := faker.Hacker().SaySomethingSmart()
	postAttestTodoMapInterface(task_msg)
}
