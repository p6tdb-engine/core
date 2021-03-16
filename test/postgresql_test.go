package test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/lhsribas/P6TDB/model"
	"gopkg.in/yaml.v2"
)

var dataYaml = `
database:
  kind: oracle
  schema:
    name: db02
    tables:
      - table:
          name: monthly_savings
          fields:
            - field:
                name: name
                dataType: string
            - field:
                name: price
                dataType: numeric
      - table:
          name: day_savings
          fields:
            - field:
                name: name
                dataType: string
            - field:
                name: price
                dataType: numeric
`
var dataJson = `
{
    "database" : {
        "kind": "oracle",
        "schema" : {
            "name" : "db02",
            "tables" : [
                {
                    "table" : {
                        "name": "monthly_savings",
                        "fields": [ 
                            
                            {
                                "field" : {"name": "name", "dataType": "string"}
                            },
                            {
                                "field" : {"name": "price", "dataType": "numeric"}
                            }
                            
                        ]
                    } 
                } ,
                {
                    "table" : {
                        "name": "day_savings",
                        "fields": [ 
                            
                            {
                                "field" : {"name": "name", "dataType": "string"}
                            },
                            {
                                "field" : {"name": "price", "dataType": "numeric"}
                            }
                            
                        ]
                        
                    } 
                } 
            ]
        }
    }
}
`

func TestYamlCreateSchema(t *testing.T) {

	db := model.Engine{}

	err := yaml.Unmarshal([]byte(dataYaml), &db)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("--- t:\n%v\n\n", db)

}

func TestJsonCreateSchema(t *testing.T) {

	db := model.Engine{}

	err := json.Unmarshal([]byte(dataJson), &db)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("--- t:\n%v\n\n", db)
}
