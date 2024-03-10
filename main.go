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
  }` // Your JSON data here

  var input map[string]any
  if err := json.Unmarshal([]byte(jsonData), &input); err != nil {
      log.Fatalf("Error parsing JSON: %v", err)
  }

  transformed, err := transform.ParseMap(input)
  if err != nil {
      log.Fatalf("Error transforming JSON: %v", err)
  }

  transformedJSON, err := json.MarshalIndent(transformed, "", "  ")
  if err != nil {
      log.Fatalf("Error marshalling transformed JSON: %v", err)
  }

  fmt.Println(string(transformedJSON))
}