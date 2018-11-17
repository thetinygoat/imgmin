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
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// error check helper
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// pwd helper
func getPwd() string {
	pwd, err := os.Getwd()
	check(err)

	return pwd
}

// file list generator
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

// file mime type verificaton helper
func getFileType(f *os.File) string {
	buffer := make([]byte, 512)
	_, err := f.Read(buffer)
	check(err)
	contentType := http.DetectContentType(buffer)

	return contentType
}

// compression helper
func proceedToCompression(b bool, l []string, c []Config) {
	pwd := getPwd()
	if b {
		os.Mkdir(path.Join(pwd, "dist"), 0700)
		for _, f := range l {

			fileExt := filepath.Ext(f)
			switch fileExt {
			case ".jpg", ".jpeg":
				decodeAndCompressJpg(c[0].Compression, f, pwd)

			case ".png":
				decodeAndCompressPng(png.CompressionLevel(c[1].Compression), f, pwd)
			}

		}
	}
}

// verify mime type
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

// check for empty file list
func checkEmptyList(l []string) {
	if len(l) == 0 {
		fmt.Println("No image file(s) found in this directory!")
		os.Exit(1)
	}
}
