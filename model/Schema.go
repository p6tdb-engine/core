package model

// Schema
type Schema struct {
	Name   string   `yaml:"name"`
	Tables []Tables `yaml:"tables"`
}
