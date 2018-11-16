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
)

// Config file struct
type Config struct {
	Type        string
	Compression int
}

const configJSON = `[
    {
        "type": "jpg",
        "compression": 50
    },
    {
        "type": "png",
        "compression": -2
    }
]`

func createConfigFile() {
	f, err := os.Create("config.json")
	check(err)
	defer f.Close()
	f.Write([]byte(configJSON))
	f.Sync()
}

func parseConfigFile() []Config {
	f, err := os.Open("config.json")
	check(err)
	defer f.Close()
	fileSlice, err := ioutil.ReadAll(f)
	check(err)

	var config []Config
	json.Unmarshal(fileSlice, &config)

	return config
}

func configFileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err != nil, err
}

func initConfig(f string) []Config {

	if b, _ := configFileExists(f); !b {
		fmt.Println("does not")
		createConfigFile()
	}
	fmt.Println("already exists")
	config := parseConfigFile()
	return config
}
