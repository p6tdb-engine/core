package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/lhsribas/P6TDB/model"
	"gopkg.in/yaml.v2"
)

var data = `
database:
  kind: oracle
  schema:
    name: teste
`

func TestCreateTables(t *testing.T) {

	db := model.Engine{}

	err := yaml.Unmarshal([]byte(data), &db)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("--- t:\n%v\n\n", db)

}
