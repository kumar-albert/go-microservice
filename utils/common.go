package utils

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func ParseYaml(file string, output map[interface{}]interface{}) {
	yfile, err := ioutil.ReadFile(file)

	if err != nil {

		log.Fatal(err)
	}

	err2 := yaml.Unmarshal(yfile, &output)

	if err2 != nil {

		log.Fatal(err2)
	}

}
