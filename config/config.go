// Alertie
// Copyright (c) 2018 Tobias Urdin
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package config

import (
	"os"
	"gopkg.in/ini.v1"
	"github.com/tobias-urdin/alertie/log"
)

var (
	Cfg *ini.File
	Connection string
	Lookups []string
)

func Init(configFile string) {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Error("Config file does not exist: %s", configFile)
		os.Exit(1)
	}

	Cfg, err := ini.Load(configFile)
	if err != nil {
		log.Error("Failed to parse config file: %s error: %s", configFile, err)
		os.Exit(1)
		return
	}

	Cfg.BlockMode = false

	database := Cfg.Section("database")
	Connection = database.Key("connection").String()

	nsq := Cfg.Section("nsq")
	Lookups = nsq.Key("lookups").Strings(",")
}
