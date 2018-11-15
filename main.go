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
	"os"
	"path/filepath"
)

func main() {
	dir := getPwd()
	jpegList, _ := genFileSlice(dir)
	willProceed := verifyJpeg(jpegList)
	if willProceed {
		minifyJpeg(50, jpegList, dir)
	}
}

func getPwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return pwd
}

func genFileSlice(d string) ([]string, []string) {
	files, err := ioutil.ReadDir(d)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var jpegSlice []string
	var pngSlice []string

	for _, f := range files {

		fileExt := filepath.Ext(f.Name())
		switch fileExt {
		case ".jpg", ".jpeg":
			jpegSlice = append(jpegSlice, f.Name())
		case ".png":
			pngSlice = append(pngSlice, f.Name())
		}
	}

	return jpegSlice, pngSlice
}
