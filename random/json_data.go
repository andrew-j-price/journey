/*
From: https://stackoverflow.com/a/53998134/17207113
Playground: https://go.dev/play/p/P8cGP1mTDmD
Helpful: https://mholt.github.io/json-to-go/
*/

package random

import (
	"encoding/json"
	"fmt"
	"log"
)

func JsonDataMain() {
	jsonData := `{
    "string": "string_value",
    "number": 123.45,
    "js_array": ["a", "b", "c"],
    "integer": 678,
    "subtype": {
        "number_array": [1, 2, 3]
      }
    }`

	m := map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonData), &m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("# JSON Data analysis\n")
	for key, value := range m {
		switch v := value.(type) {
		case int:
			fmt.Printf("Key: %s, Integer: %d\n", key, v)
		case float64:
			fmt.Printf("Key: %s, Float: %v\n", key, v)
		case string:
			fmt.Printf("Key: %s, String: %s\n", key, v)
		case map[string]interface{}:
			fmt.Printf("Key: %s, Subtype: %+v\n", key, v)
		case []interface{}:
			//TODO: Read through each item in the interface and work out what type it is.
			fmt.Printf("Key: %s, []interface: %v\n", key, v)
		default:
			fmt.Printf("Key: %s, unhandled type: %+v\n", key, v)
		}
	}
}
