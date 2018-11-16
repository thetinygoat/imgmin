// Copyright (C) 2018  Sachin Saini

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func main() {
	dir := getPwd()
	imgList := genFileSlice(dir)
	willProceed := verifyImg(imgList)
	if willProceed {
		os.Mkdir(path.Join(dir, "dist"), 0700)
		for _, f := range imgList {

			fileExt := filepath.Ext(f)
			switch fileExt {
			case ".jpg", ".jpeg":
				minifyJpeg(50, f, dir)

			case ".png":
				minifyPng(50, f, dir)
			}

		}
	}

	parseJSON()
}

func getPwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return pwd
}

func genFileSlice(d string) []string {
	files, err := ioutil.ReadDir(d)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var imgSlice []string

	for _, f := range files {

		fileExt := filepath.Ext(f.Name())
		switch fileExt {
		case ".jpg", ".jpeg", ".png":
			imgSlice = append(imgSlice, f.Name())
		}

	}
	return imgSlice
}

func verifyImg(fs []string) bool {
	for _, file := range fs {
		imgfile, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer imgfile.Close()
		contentType, err := getFileContentType(imgfile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		switch contentType {
		case "image/jpeg", "image/png":
			fmt.Println("verified mime type of", file, "as", contentType)
		default:
			fmt.Println("invalid format ", contentType, " of file ", file)
			os.Exit(1)
		}
	}
	return true
}

// Config struct
type Config struct {
	Type        string
	Compression int
}

func parseJSON() {
	f, err := os.Open("config.json")

	if err != nil {
		panic(err)
	}
	defer f.Close()
	bv, _ := ioutil.ReadAll(f)

	var conf []Config

	json.Unmarshal([]byte(bv), &conf)

	fmt.Println(conf)

}
