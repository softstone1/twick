package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/softstone1/twick/transform"
)
func main() {
  // Example input JSON data.
  jsonData := `{
	"number_1": {
	  "N": "1.50"
	},
	"string_1": {
	  "S": "784498 "
	},
	"string_2": {
	  "S": "2014-07-16T20:55:46Z"
	},
	"map_1": {
	  "M": {
		"bool_1": {
		  "BOOL": "truthy"
		},
		"null_1": {
		  "NULL ": "true"
		},
		"list_1": {
		  "L": [
			{
			  "S": ""
			},
			{
			  "N": "011"
			},
			{
			  "N": "5215s"
			},
			{
			  "BOOL": "f"
			},
			{
			  "NULL": "0"
			}
		  ]
		}
	  }
	},
	"list_2": {
	  "L": "noop"
	},
	"list_3": {
	  "L": [
		"noop"
	  ]
	},
	"": {
	  "S": "noop"
	}
  }`

  var input any
  if err := json.Unmarshal([]byte(jsonData), &input); err != nil {
      log.Fatalf("Error parsing JSON: %v", err)
  }

  transformed, err := transform.ParseInput(input)
  if err != nil {
      log.Fatalf("Error transforming JSON: %v", err)
  }

  transformedJSON, err := json.MarshalIndent([]any{transformed}, "", "  ")
  if err != nil {
      log.Fatalf("Error marshalling transformed JSON: %v", err)
  }

  fmt.Println(string(transformedJSON))
}