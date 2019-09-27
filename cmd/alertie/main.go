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

package main

import (
	"flag"
	"github.com/tobias-urdin/alertie/log"
	"github.com/tobias-urdin/alertie/config"
	"github.com/tobias-urdin/alertie/database"
)

var (
	configFile string
	createAdmin bool
	resetAdmin bool
)

func main() {
	flag.StringVar(&configFile, "config", "alertie.ini", "Path to config file")
	flag.BoolVar(&createAdmin, "createadmin", false, "Create admin user if it does not exist")
	flag.BoolVar(&resetAdmin, "resetadmin", false, "Reset the admin user password")

	flag.Parse()

	log.Init("cli")
	config.Init(configFile)

	db := database.Init(config.Connection)
	defer db.Close()

	if createAdmin {
		return
	}

	if resetAdmin {
		return
	}
}
