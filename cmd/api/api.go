package main

import (
	"log"
	"net/http"

	helper "github.com/andrew-j-price/journey/pkg/helpers"
	"github.com/andrew-j-price/journey/pkg/logger"
	"github.com/andrew-j-price/journey/pkg/random"
)

func main() {
	listen_address := helper.GetEnv("LISTEN_ADDRESS", ":8080")
	// logger.Info.Printf("Startup binding to %s\n", listen_address)
	log.Printf("Startup binding to %s\n", listen_address)
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerRoot)
	mux.HandleFunc("/dogs", routeDog)
	mux.HandleFunc("/employees", routeEmployees)
	mux.HandleFunc("/ping", helper.HttpRoutePing)
	// logger.Info.Printf("Web server starting\n")
	log.Printf("Web server starting\n")
	log.Fatal(http.ListenAndServe(listen_address, mux))
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	// NOTE: the "else if" is presently useless here
	if r.URL.Path == "/" {
		helper.HttpRouteDefault(w, r)
	} else if r.URL.Path == "/default" {
		helper.HttpRouteDefault(w, r)
	} else {
		helper.HttpRoute404(w, r)
	}
}

func routeDog(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("%v received %v call from %v using %v", r.URL.Path, r.Method, r.RemoteAddr, r.UserAgent())
	hobbies := []string{"sleeping", "eating", "playing"}
	message := map[string]interface{}{"jack": 15, "murphy": 7, "are": "good", "hobbies": hobbies}
	helper.HttpJsonResponse(w, r, 200, message)
}

// init and store employees here
var employees = random.EmployeeSliceOfStructs()

func routeEmployees(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("%v received %v call", r.URL.Path, r.Method)
	switch r.Method {
	case "GET":
		message := map[string]interface{}{"employees": employees}
		helper.HttpJsonResponse(w, r, 200, message)
		/*employeesJSON, _ := json.Marshal(employees)
		fmt.Println(string(employeesJSON))
		w.Write(employeesJSON)*/
	case "POST":
		message := map[string]interface{}{"todo": "todo"}
		helper.HttpJsonResponse(w, r, 201, message)
	default:
		message := map[string]interface{}{"error": "unsupported method"}
		helper.HttpJsonResponse(w, r, 400, message)
	}
}
