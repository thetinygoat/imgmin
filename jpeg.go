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

func minifyJpeg(q int, f string, pwd string) {
	decodeAndCompressJpg(q, f, pwd)
}

func decodeAndCompressJpg(q int, f string, pwd string) {
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
	fmt.Println("successfully compressed", f)
}
