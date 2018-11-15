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
	"image/jpeg"
	"os"
	"path"
)

func verifyJpeg(js []string) bool {
	for _, file := range js {
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
		case "image/jpeg":
			fmt.Println("verified mime type of", file, "as", contentType)
		case "image/png":
			fmt.Println("error : expected ", file, " to be of type image/jpeg but recieved", contentType)
			os.Exit(1)
		default:
			fmt.Println("invalid format ", contentType, " of file ", file)
			os.Exit(1)
		}
	}
	return true
}

func minifyJpeg(q int, fl []string, pwd string) {
	os.Mkdir(path.Join(pwd, "dist"), 0700)
	for _, f := range fl {
		imgFile, err := os.Open(f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		img, err := jpeg.Decode(imgFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		imgC, err := os.Create(path.Join(pwd, "dist", f))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = jpeg.Encode(imgC, img, &jpeg.Options{Quality: q})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = imgC.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = imgFile.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}
}
