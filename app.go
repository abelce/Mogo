package mogo

import (
	"io"
	"io/ioutil"
	"log"
	"os"
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
	RENDER_DIR = "./root"
)

func Build() {
	if !isExists(PUBLICSH_DIR) {
		err := os.Mkdir(PUBLICSH_DIR, 0777)
		if err != nil {
			log.Panic("create publish dir error -- " + err.Error())
		}
	}
	rf := new(RenderFactory)
	rf.Render(RENDER_DIR)

	//copy res
	err := copyDir(RENDER_DIR+"/assets", PUBLICSH_DIR+"/assets")
	if err != nil {
		log.Println(err)
	}
	log.Println("blog process ok！")
}

func copyFile(src, dst string) (w int64, err error) {
	f, err := os.Open(src)
	if err != nil {
		return
	}
	defer f.Close()
	dstf, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return
	}
	defer dstf.Close()
	return io.Copy(dstf, f)
}

func copyDir(source, dest string) (err error) {
	fi, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return &CustomError{"Source is not a directory"}
	}

	err = os.MkdirAll(dest, fi.Mode())
	if err != nil {
		return err
	}
	entries, err := ioutil.ReadDir(source)
	for _, entry := range entries {
		sfp := source + "/" + entry.Name()
		dfp := dest + "/" + entry.Name()
		if entry.IsDir() {
			err = copyDir(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		} else {
			_, err = copyFile(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		}

	}
	return
}

type CustomError struct {
	msg string
}

func (e *CustomError) Error() string {
	return e.msg
}
