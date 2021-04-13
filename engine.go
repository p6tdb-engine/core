package main

import (
	"log"

	"github.com/lhsribas/P6TDB/engines"
	"github.com/lhsribas/P6TDB/model"
	"gopkg.in/yaml.v2"
)

var dataYaml = `
database:
  kind: postgresql
  schema:
    name: db02
    tables:
      - table:
        name: person
        fields:
          - name: name
            type: string
            description: Some name
            nullabel: true
          - name: age
            type: int
            description: Age Person
            nullabel: true
          - name: age
		    type: int
			description: Age Person
            nullabel: true
      - table:
        name: document
        fields:
          - name: document
            type: string
            description: Some name
            nullabel: true
          - name: age
            type: int
            description: Age Person
            nullabel: true 
`

var postgresql = engines.PostgreSQLEngine{}

func main() {

	yaml2Go := parser()

	switch yaml2Go.Database.Kind {
	case "postgresql":
		postgresql.Engine(yaml2Go)
		break

	case "oracle":
		break

	case "mysql":
		break
	}

}

func parser() model.Yaml2Go {

	// model
	m := model.Yaml2Go{}

	err := yaml.Unmarshal([]byte(dataYaml), &m)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return m
}
