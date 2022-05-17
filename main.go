package main

import (
	"fmt"
	"io/ioutil"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/ghodss/yaml"
)

const (
	swaggerFilePath = "swagger.yaml"
)

func main() {
	input, err := ioutil.ReadFile(swaggerFilePath)
	if err != nil {
		panic(err)
	}

	var doc openapi2.T
	if err = yaml.Unmarshal(input, &doc); err != nil {
		panic(err)
	}

	host := "{{host}}"
	basePath := doc.BasePath
	for path, items := range doc.Paths {
		for method, op := range items.Operations() {
			fmt.Printf("%s %s HTTP/1.1\n", method, host+basePath+path)
			if len(op.Consumes) > 0 {
				fmt.Printf("Content-Type: %s\n", op.Consumes[0])
			}
			if op.Security != nil {
				for _, auth := range *op.Security {
					if _, ok := auth["api_key"]; ok {
						fmt.Printf("Authorization: Bearer\n")
					}
				}
			}
			fmt.Println()
		}
	}
}
