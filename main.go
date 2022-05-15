package main

import (
	"fmt"
	"io/ioutil"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/ghodss/yaml"
)

const (
	swaggerFilePath = "gen/http/openapi.json"
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

	//if err = json.Unmarshal(input, &doc); err != nil {
	//	panic(err)
	//}
	fmt.Println(doc.BasePath)
	for path, items := range doc.Paths {
		fmt.Println(path)
		fmt.Println(items.Operations())
		fmt.Println(items.Parameters)
	}

	//outputYAML, err := yaml.Marshal(doc)
	//if err != nil {
	//	panic(err)
	//}
	////outputYAML
	//fmt.Println(outputYAML)

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
