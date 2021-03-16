package model

// Database
type Schema struct {
	Name   string
	Tables []struct {
		Table Table
	} `yaml:"tables"`
}
