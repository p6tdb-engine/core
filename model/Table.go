package model

type Table struct {
	Name   string  `yaml:"name"`
	Fields []Filed `yaml:"fields"`
}
