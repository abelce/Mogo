/*
 * File: yamlParser.go
 * Project: Mogo
 * File Created: Thursday, 16th May 2019 3:16:33 pm
 * Author: zxtang (1061225829@qq.com)
 * -----
 * Last Modified: Thursday, 16th May 2019 3:16:33 pm
 * Modified By: zxtang (1061225829@qq.com>)
 * -----
 * Copyright 2017 - 2019 Your Company, Your Company
 */
package mogo

import (
	"log"
	"os"
	"strings"

	"github.com/scottkiss/go-gypsy/yaml"
)

type YamlParser struct{}

var YAML_FILES = [3]string{"global.yml", "nav.yml", "pages.yml"}

func (yp *YamlParser) parse(root string) map[string]interface{} {
	var yamlFilesConfig = make(map[string]interface{})

	if !strings.HasSuffix(root, "/") {
		root += "/"
	}
	for _, yamlFile := range YAML_FILES {
		path := root + yamlFile
		if !isExists(path) {
			log.Panic(path + " file not found!")
			os.Exit(1)
		}

		config, err := yaml.ReadFile(path)
		if err != nil {
			log.Panic(err)
			os.Exit(1)
		}

		yamlFilesConfig[yamlFile] = config
	}

	return yamlFilesConfig
}
