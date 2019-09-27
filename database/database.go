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

package database

import (
	"github.com/tobias-urdin/alertie/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var globDb *gorm.DB

func Init(uri string) *gorm.DB {
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		log.Fatal("Could not open database connection: %s", err)
	}

	err = db.DB().Ping()
	if err != nil {
		log.Fatal("Could not ping database: %s", err)
	}

	globDb = db
	return globDb
}

func GetDB() *gorm.DB {
	return globDb
}
