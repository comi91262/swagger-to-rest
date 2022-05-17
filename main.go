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

	fmt.Println(doc.BasePath)
	for path, items := range doc.Paths {
		fmt.Println(path)
		for op, hoge := range items.Operations() {
			fmt.Println(op)
			for _, hoge := range hoge.Parameters {
				fmt.Printf("debug: %v\n", hoge)
			}
			//for resp, hoge := range hoge.Responses {
			// fmt.Printf("debug: %v\n", resp)
			// fmt.Printf("debug: %v\n", hoge)
			//}
		}
	}
}

//swagger: "2.0"
//info:
//  title: Calculator Service
//  description: HTTP service for multiplying numbers, a goa teaser
//  version: ""
//host: localhost:8088
//consumes:
//- application/json
//- application/xml
//- application/gob
//produces:
//- application/json
//- application/xml
//- application/gob
//paths:
//  /multiply/{a}/{b}:
//    get:
//      tags:
//      - calc
//      summary: multiply calc
//      operationId: calc#multiply
//      parameters:
//      - name: a
//        in: path
//        description: Left operand
//        required: true
//        type: integer
//      - name: b
//        in: path
//        description: Right operand
//        required: true
//        type: integer
//      responses:
//        "200":
//          description: OK response.
//          schema:
//            type: integer
//            format: int64
//      schemes:
//      - http
