package main

import (
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"strings"
	"text/template"
)

var funcs = template.FuncMap{
	"titleCap": strings.Title,
	"isBase":   isBaseType,
}

func main() {
	confFile, err := os.Open("v1/configs/config.yml")
	if err != nil {
		panic(err)
	}
	defer confFile.Close()

	config := Root{}
	dec := yaml.NewDecoder(confFile)
	if err := dec.Decode(&config); err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("").Funcs(funcs).
		ParseGlob("v1/tpl/*"))

	for _, mp := range config.Maps {
		for _, def := range mp.Values {
			def.Package = config.Package
			def.Key = mp.Key

			iFile, err := os.Create(path.Join(config.Dir, def.Name+".go"))
			if err != nil {
				panic(err)
			}

			err = tpl.ExecuteTemplate(iFile, "interface", def)
			if err != nil {
				panic(err)
			}

			_ = iFile.Close()
		}
	}
}

type Root struct {
	// Package name for generated types
	Package string `yaml:"package" json:"package"`

	// Target output dir for generated types
	Dir string `yaml:"dir" json:"dir"`

	// Map definition list
	Maps []MapKeyDefinition `yaml:"maps" json:"maps"`
}

type MapKeyDefinition struct {

	// Type of the map type key
	Key string `yaml:"key" json:"key"`

	// Map value definitions
	Values []MapValDefinition `yaml:"values" json:"values"`
}

type MapValDefinition struct {
	Package string `yaml:"-" json:"-"`
	Key     string `yaml:"-" json:"-"`

	// Type of the map values
	Type string `yaml:"type" json:"type"`

	// Name of the generated map type
	Name string `yaml:"name" json:"name"`
}

func isBaseType(kind string) bool {
	return kind == "bool" ||
		kind == "int" ||
		kind == "int8" ||
		kind == "int16" ||
		kind == "int32" ||
		kind == "int64" ||
		kind == "uint" ||
		kind == "uint8" ||
		kind == "uint16" ||
		kind == "uint32" ||
		kind == "uint64" ||
		kind == "float32" ||
		kind == "float64" ||
		kind == "complex64" ||
		kind == "complex128" ||
		kind == "string" ||
		kind == "byte" ||
		kind == "rune" ||
		kind == "interface{}"
}
