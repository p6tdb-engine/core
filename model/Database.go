package model

type Database struct {
	Kind   string `yaml:"kind"`
	Schema Schema `yaml:"schema"`
}
