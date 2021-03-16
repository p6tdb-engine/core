package model

type Table struct {
	Name   string `yaml:"name"`
	Fields []struct {
		Field Field
	} `yaml:"fields"`
}
