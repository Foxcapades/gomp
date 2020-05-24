package main

import (
	"math/rand"
	"os"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/Foxcapades/Argonaut/v0"
	"github.com/Foxcapades/gomp/v1/internal/gen"
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/yaml.v3"
)

var funcs = template.FuncMap{
	"titleCap": strings.Title,
	"isBase":   gen.IsBaseType,
	"pad":      gen.Pad,
	"defVal":   gen.DefaultValue,
	"trimR":    strings.TrimRight,
}

var skipTests = false

func main() {
	var file string

	logrus.SetFormatter(new(prefixed.TextFormatter))

	_, err := cli.NewCommand().
		Flag(cli.LFlag("skip-tests", "Skip generating tests").
			Bind(&skipTests, true)).
		Arg(cli.NewArg().
			Name("config-file").
			Description("Configuration file containing definitions of the map types to generate").
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

	rand.Seed(time.Now().UnixNano())
	pack := path.Join(config.Repo, config.Dir)
	for _, mp := range config.Maps {
		for _, def := range mp.Values {
			def.Package = config.Package
			def.Key = mp.Key
			def.Repo = pack
			execTemplate(tpl, &def, config.Dir)
		}
	}
}

func execTemplate(tpl *template.Template, def *MapValDefinition, dir string) {
	fName := path.Join(dir, def.Name)
	oFile, err := os.Create(fName + ".go")
	check(err)
	defer oFile.Close()

	check(tpl.ExecuteTemplate(oFile, "interface", def))

	if !skipTests {
		tFile, err := os.Create(fName + "_test.go")
		check(err)
		defer tFile.Close()

		check(tpl.ExecuteTemplate(tFile, "tests", def))
	}
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

	Repo string `yaml:"repo" json:"repo"`

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
	Repo    string `yaml:"-" json:"-"`

	// Type of the map values
	Type string `yaml:"type" json:"type"`

	// Name of the generated map type
	Name string `yaml:"name" json:"name"`
}
