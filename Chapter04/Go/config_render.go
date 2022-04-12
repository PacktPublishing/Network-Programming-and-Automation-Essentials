package main

import (
	"io/ioutil"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Router struct {
	Id     int    `yaml:"id"`
	Name   string `yaml:"name"`
	Toname string `yaml:"to_name"`
}

type RouterList []Router

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var routers RouterList

	yamlFile, err := ioutil.ReadFile("router_definitions.yaml")
	check(err)

	err = yaml.Unmarshal(yamlFile, &routers)
	check(err)

	templateFile, err := ioutil.ReadFile("cisco_template_go.txt")
	check(err)

	for _, router := range routers {

		outFile, err := os.Create(router.Name + "_router_config.txt")
		check(err)

		tmpl, err := template.New("render").Parse(string(templateFile))
		check(err)

		err = tmpl.Execute(outFile, router)
		check(err)
	}
}
