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
	"io/ioutil"
	"os"
	"path"
)

// Config file struct
type Config struct {
	Type        string
	Compression int
}

// config file skeleton
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

// generate config file paths
var home = os.Getenv("HOME")
var configDirPath = path.Join(home, ".config", "gomin")
var configFilePath = path.Join(configDirPath, "config.json")

// create config file and write skeleton to it
func createConfigFile(p string) {
	f, err := os.Create(p)
	check(err)
	defer f.Close()
	f.Write([]byte(configJSON))
	f.Sync()
}

// parse config file and create a struct from it
func parseConfigFile(p string) []Config {
	f, err := os.Open(p)
	check(err)
	defer f.Close()
	fileSlice, err := ioutil.ReadAll(f)
	check(err)

	var config []Config
	json.Unmarshal(fileSlice, &config)

	return config
}

// init config everytime the program is run
func initConfig() []Config {

	if _, err := os.Open(configFilePath); os.IsNotExist(err) {
		os.Mkdir(configDirPath, 0700)
		createConfigFile(configFilePath)
	}
	config := parseConfigFile(configFilePath)
	return config
}
