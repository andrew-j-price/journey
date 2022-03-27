package main

import (
	"encoding/json"
	"log"
	"net/http"

	helper "github.com/andrew-j-price/journey/helpers"
	"github.com/andrew-j-price/journey/logger"
	"github.com/andrew-j-price/journey/random"
)

func apiMain() {
	listen_address := helper.GetEnv("LISTEN_ADDRESS", ":8080")
	logger.Info.Printf("Startup binding to %s\n", listen_address)
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerRoot)
	mux.HandleFunc("/dogs", handlerDog)
	mux.HandleFunc("/employees", handlerEmployees)
	mux.HandleFunc("/ping", handlerPing)
	logger.Info.Printf("Web server started\n")
	log.Fatal(http.ListenAndServe(listen_address, mux))
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	// NOTE: the "else if" is presently useless here
	if r.URL.Path == "/" {
		handlerDefault(w, r)
	} else if r.URL.Path == "/default" {
		handlerDefault(w, r)
	} else {
		handler404(w, r)
	}
}

func handler404(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("%v received %v call", r.URL.Path, r.Method)
	message := map[string]string{"404": r.URL.Path}
	jsonResponse(w, 404, message)
}

func handlerDefault(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("%v received %v call", r.URL.Path, r.Method)
	message := map[string]interface{}{
		"host_name":      "socket.gethostname()",
		"host_ip":        "socket.gethostbyname(socket.gethostname())",
		"server_address": "request.host",
		"uri":            r.URL.Path,
		"uuid":           "uuid4()",
		"date":           helper.TimeNow(),
		"remote_agent":   r.UserAgent(),
		"remote_ip":      r.RemoteAddr,
	}
	jsonResponse(w, 200, message)
}

func handlerDog(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("%v received %v call from %v using %v", r.URL.Path, r.Method, r.RemoteAddr, r.UserAgent())
	hobbies := []string{"sleeping", "eating", "playing"}
	message := map[string]interface{}{"jack": 15, "murphy": 7, "are": "good", "hobbies": hobbies}
	jsonResponse(w, 200, message)
}

// not using jsonResponse, but writing byte style.  Do this for example
func handlerPing(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("%v received %v call", r.URL.Path, r.Method)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"ping":"pong"}`))
}

func jsonResponse(w http.ResponseWriter, httpStatusCode int, httpResponse interface{}) {
	message, err := json.Marshal(httpResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(httpStatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}

// init and store employees here
var employees = random.EmployeeSliceOfStructs()

func handlerEmployees(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("%v received %v call", r.URL.Path, r.Method)
	switch r.Method {
	case "GET":
		message := map[string]interface{}{"employees": employees}
		jsonResponse(w, 200, message)
		/*employeesJSON, _ := json.Marshal(employees)
		fmt.Println(string(employeesJSON))
		w.Write(employeesJSON)*/
	case "POST":
		message := map[string]interface{}{"todo": "todo"}
		jsonResponse(w, 201, message)
	default:
		message := map[string]interface{}{"error": "unsupported method"}
		jsonResponse(w, 400, message)
	}
}
