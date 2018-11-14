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
)

func main() {
	dir := getPwd()
	fileList := genFileSlice(dir)
	verifyjpeg(fileList)
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

	var filelist []string

	for _, f := range files {
		filelist = append(filelist, f.Name())
	}
	return filelist
}
