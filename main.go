package main

import (
	"fmt"

	"github.com/ablece/mogo/mogo"
)

/*
 * File: main.go
 * Project: Mogo
 * File Created: Thursday, 16th May 2019 2:35:32 pm
 * Author: zxtang (1061225829@qq.com)
 * -----
 * Last Modified: Thursday, 16th May 2019 2:35:33 pm
 * Modified By: zxtang (1061225829@qq.com>)
 * -----
 * Copyright 2017 - 2019 Your Company, Your Company
 */

const (
	USAGE = `
gosk is a static site generator in Go

Usage:

        mogo command [args...]

The commands are:

    new                         create a new folder
	build	        			build and generate site.
	version         			print gosk version

`
)

const Version = "0.0.1"

var WebsitePath = "."

// func init() {
// 	yp := new(mogo.YamlParser)
// 	yamlData := yp.parse(".")
// }

func main() {

	// flag.Parse()
	// args := flag.Args()
	// argsLength := len(args)
	// if argsLength == 0 || argsLength > 3 {
	// 	Usage()
	// 	os.Exit(1)
	// }
	// switch args[0] {
	// case "new":
	// 	if argsLength == 2 {

	// 	}
	// case "build":
	// 	mogo.Build()
	// case "usage":
	// 	Usage()
	// case "version":
	// 	fmt.Println(Version)
	// }

	mogo.Build()
}

func Usage() {
	fmt.Println(USAGE)
}
