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
	"os"
)

func verifyjpeg(f []string) {
	for _, file := range f {
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
		if contentType != "image/jpeg" || contentType != "image/png" {
			fmt.Println("invalid image format ", contentType, " of file ", file)
		} else {
			fmt.Println("false")
		}
		fmt.Println(contentType)
	}
}
