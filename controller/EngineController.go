package controller

import (
	"log"

	"github.com/lhsribas/P6TDB/engines"
	"github.com/lhsribas/P6TDB/model"
	"gopkg.in/yaml.v2"
)

type EngineController struct{}

var postgresql = engines.PostgreSQLEngine{}
var mysql = engines.MySQLEngine{}
var oracle = engines.OracleEngine{}
var sqlserver = engines.SQLServer{}

func Engine(data []byte, extension string) {

	db := unmarshall(data)

	switch db.Database.Kind {

	case "oracle":
		oracle.Engine(db)
		break

	case "mysql":
		mysql.Engine(db)
		break

	case "postgresql":
		postgresql.Engine(db)
		break

	case "sqlserver":
		sqlserver.Engine(db)
		break
	}
}

// unmarshall structure definition
func unmarshall(data []byte) model.Yaml2Go {
	db := model.Yaml2Go{}
	err := yaml.Unmarshal([]byte(data), &db)

	if err != nil {
		log.Fatalf("YAML_PARSER_ERROR: %v", err)
	}

	return db
}
