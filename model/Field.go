package model

type Filed struct {
	Field    string `yaml:"field"`
	DataType string `yaml:"dataType"`
	Size     int    `yaml:"size"`
}
