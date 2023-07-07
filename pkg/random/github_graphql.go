package random

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/andrew-j-price/journey/pkg/helpers"
	"github.com/andrew-j-price/journey/pkg/logger"
)

/*
https://docs.github.com/en/graphql/overview/explorer
query {
  viewer {
    login
    createdAt
  }
  rateLimit {
    limit
    cost
    remaining
    resetAt
  }
}

{
  "data": {
    "viewer": {
      "login": "andrew-j-price",
      "createdAt": "2015-06-07T18:14:26Z"
    },
    "rateLimit": {
      "limit": 5000,
      "cost": 1,
      "remaining": 4995,
      "resetAt": "2022-03-27T22:32:34Z"
    }
  }
}

*/

type GraphQLQuery struct {
	Query string `json:"query"`
}

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type GraphQLResponse struct {
	Data struct {
		Viewer struct {
			Login     string `json:"login"`
			CreatedAt string `json:"createdAt"`
		} `json:"viewer"`
		RateLimit struct {
			Limit     int    `json:"limit"`
			Cost      int    `json:"cost"`
			Remaining int    `json:"remaining"`
			ResetAt   string `json:"resetAt"`
		} `json:"rateLimit"`
	} `json:"data"`
}

func GithubGraphqlMain() {
	fmt.Println("github")
	token := helpers.GetEnv("GITHUB_TOKEN", "")
	if token == "" {
		logger.Fatal.Println("Environment variable GITHUB_TOKEN is missing or unset")
		os.Exit(1)
	}
	url := "https://api.github.com/graphql"
	/*
		// payload := map[string]interface{}{"jack": 15, "murphy": 7, "are": "good"}
		payload := map[string]interface{}{"query": "{ viewer { login } rateLimit { limit cost remaining resetAt }"}
		jsonData := helpers.RestJsonMarshalData(payload)

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	*/

	gqlReq := GraphQLQuery{
		Query: `
			query { 
				viewer { 
					login
					createdAt
				}
				rateLimit {
					limit
					cost
					remaining
					resetAt
				}
			}
		`,
	}

	// Encode request
	var requestBody bytes.Buffer
	if err := json.NewEncoder(&requestBody).Encode(gqlReq); err != nil {
		log.Fatal(err)
	}

	// By default, all gql requests are POSTs to
	req, err := http.NewRequest(http.MethodPost, url, &requestBody)
	if err != nil {
		log.Fatal(err)
	}

	//var bearerToken = "Bearer " + token
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	/*
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logger.Error.Println(err)
			// return nil
		}

		if !helpers.RestResponseCodeAnalysis(resp, 200) {
			fmt.Println("Does not match")
			// return nil
		} else {
			respData := helpers.RestResponseBody(resp)
			respPayload := helpers.RestJsonUnmarshalData(respData)
			fmt.Println("here")
			fmt.Println(respPayload)
			fmt.Printf("Login: %v\n", respPayload["viewer"])
			// return respPayload
		}
	*/
	// Make the request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Decode request response
	var response GraphQLResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	// Log Response
	log.Println(response) // {{{Brazil BRL 55}}}
	log.Println(response.Data.RateLimit.Remaining)
	log.Println(response.Data.Viewer.Login)
	fmt.Println(helpers.PrettyPrintStruct(response))
	os.Exit(0)

}
