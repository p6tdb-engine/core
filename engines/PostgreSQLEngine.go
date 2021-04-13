package engines

import (
	"bytes"
	"fmt"

	"github.com/lhsribas/P6TDB/commands"
	"github.com/lhsribas/P6TDB/model"
)

type PostgreSQLEngine struct{}

func (p *PostgreSQLEngine) Engine(db model.Yaml2Go) {

	tables := createTable(db.Database.Schema.Tables)

	fmt.Println(tables)
}

func createTable(tables []model.Tables) string {

	var buffer bytes.Buffer

	for _, v := range tables {

		fmt.Println(v.Name)

		buffer.WriteString(commands.CREATE_TABLE + " ")
		buffer.WriteString(commands.EXIST + " ")
		buffer.WriteString(v.Name + " ")
		buffer.WriteString("{")

		buffer.WriteString("}")
		buffer.WriteString(" ")
	}

	return buffer.String()
}
