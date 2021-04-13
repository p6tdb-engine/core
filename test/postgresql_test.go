package test

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/lhsribas/P6TDB/commands"
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
          - name: person_id
            type: PK
            description: PK
            nullabel: false
      - table:
        name: document
        fields:
          - name: document
            type: string
            description: Some name
            nullabel: true
          - name: Type
            type: string
            description: Type of Document
            nullabel: true 
`

func TestYamlCreateSchema(t *testing.T) {

	db := model.Yaml2Go{}

	err := yaml.Unmarshal([]byte(dataYaml), &db)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var buffer bytes.Buffer

	tables := db.Database.Schema.Tables

	for _, v := range tables {

		buffer.WriteString(commands.CREATE_TABLE)
		buffer.WriteString(" ")
		buffer.WriteString(commands.EXIST)
		buffer.WriteString(" ")
		buffer.WriteString(v.Name)
		buffer.WriteString(" ")
		buffer.WriteString("{")

		for _, f := range v.Fields {

			buffer.WriteString("\n")
			buffer.WriteString(f.Name)

			if f.Type == "PK" || f.Type == "pk" {
				buffer.WriteString(" PRIMARY KEY")
			}

			if f.Nullabel {
				buffer.WriteString(" NULL")
			} else {
				buffer.WriteString(" NOT NULL")
			}
		}

		buffer.WriteString("\n")

		buffer.WriteString("}")

		buffer.WriteString("\n")

	}

	//fmt.Printf("--- t:\n%v\n\n", db)
	fmt.Println(buffer.String())
}
