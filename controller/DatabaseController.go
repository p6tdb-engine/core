package controller

import (
	"fmt"
	"log"

	"github.com/lhsribas/P6TDB/model"
	"gopkg.in/yaml.v2"
)

var data = `
database:
  kind: oracle
  tables:
  - table:
      name: monthly_savings
      fields:
      - field: name
        dataType: string
        size: 255
      - field: supplier_id
        dataType: integer
        size: 10
      - field: price
        dataType: numeric
		size: 0
`

func Controller(database model.Database) {

}

func unmarshall() {
	db := model.Database{}

	err := yaml.Unmarshal([]byte(data), &db)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("--- t:\n%v\n\n", db)
}
