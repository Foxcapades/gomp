package main

import (
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/Foxcapades/Argonaut/v0"
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/yaml.v3"
)

var funcs = template.FuncMap{
	"titleCap": strings.Title,
	"isBase":   isBaseType,
	"pad":      pad,
}

func main() {
	var file string

	logrus.SetFormatter(new(prefixed.TextFormatter))

	_, err := cli.NewCommand().
		Arg(cli.NewArg().
			Name("config-file").
			Description("Configuration file containing").
			Bind(&file).
			Require()).
		Parse()
	check(err)

	confFile, err := os.Open(file)
	check(err)
	defer confFile.Close()

	config := Root{}
	dec := yaml.NewDecoder(confFile)
	check(dec.Decode(&config))

	tpl := template.Must(template.New("").Funcs(funcs).
		ParseGlob("v1/tpl/*"))

	for _, mp := range config.Maps {
		for _, def := range mp.Values {
			def.Package = config.Package
			def.Key = mp.Key
			execTemplate(tpl, &def, config.Dir)
		}
	}
}

func execTemplate(tpl *template.Template, def *MapValDefinition, dir string) {
	oFile, err := os.Create(path.Join(dir, def.Name+".go"))
	check(err)
	defer oFile.Close()

	check(tpl.ExecuteTemplate(oFile, "interface", def))
}

func check(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}

// Root Config Node
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

func isBaseType(k string) bool {
	return k == "bool" ||
		k == "int" || k == "int8" || k == "int16" || k == "int32" || k == "int64" ||
		k == "uint" || k == "uint8" || k == "uint16" || k == "uint32" || k == "uint64" ||
		k == "float32" || k == "float64" ||
		k == "complex64" || k == "complex128" ||
		k == "string" ||
		k == "byte" ||
		k == "rune" ||
		k == "interface{}"
}

func pad(a, b string) string {
	cl := 0

	if len(a) < len(b) {
		cl = len(b) - len(a)
	}

	out := make([]byte, cl)

	for i := range out {
		out[i] = ' '
	}

	return string(out)
}
