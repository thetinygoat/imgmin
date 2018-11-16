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
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getPwd() string {
	pwd, err := os.Getwd()
	check(err)

	return pwd
}

func getFileList(d string) []string {
	files, err := ioutil.ReadDir(d)
	check(err)

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

func getFileType(f *os.File) string {
	buffer := make([]byte, 512)
	_, err := f.Read(buffer)
	check(err)
	contentType := http.DetectContentType(buffer)

	return contentType
}

func proceedToCompression(b bool, l []string) {
	pwd := getPwd()
	if b {
		os.Mkdir(path.Join(pwd, "dist"), 0700)
		for _, f := range l {

			fileExt := filepath.Ext(f)
			switch fileExt {
			case ".jpg", ".jpeg":
				decodeAndCompressJpg(50, f, pwd)

			case ".png":
				decodeAndCompressPng(50, f, pwd)
			}

		}
	}
}

func verifyImg(l []string) bool {
	for _, file := range l {
		imgfile, err := os.Open(file)
		check(err)
		defer imgfile.Close()
		contentType := getFileType(imgfile)
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