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
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/tobias-urdin/alertie/log"
	"github.com/tobias-urdin/alertie/config"
	"github.com/tobias-urdin/alertie/database"
	"github.com/tobias-urdin/alertie/factory"
	_ "github.com/tobias-urdin/alertie/transports"
	"github.com/tobias-urdin/alertie/utils"
	"github.com/tobias-urdin/alertie/triggers"
	"github.com/tobias-urdin/alertie/services"
	"github.com/tobias-urdin/alertie/alerts"
	"github.com/tobias-urdin/alertie/contacts"
)

func databaseMigrate(db *gorm.DB) {
	db.AutoMigrate(&services.Service{})
	db.AutoMigrate(&alerts.Alert{})
	db.AutoMigrate(&contacts.Contact{})
}

func routeRegister(r *gin.RouterGroup) {
	triggers.Register(r)
	services.Register(r)
	alerts.Register(r)
	contacts.Register(r)
}

var configFile string

func main() {
	flag.StringVar(&configFile, "config", "alertie.ini", "Path to config file")
	flag.Parse()

	log.Init("api")
	log.Info("Starting alertie-api")
	config.Init(configFile)

	db := database.Init(config.Connection)
	defer db.Close()
	databaseMigrate(db)

	transport := factory.New("transports.nsq")
	if transport == nil {
		log.Fatal("Invalid transport")
	}

	triggers.Init()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	binding.Validator.RegisterValidation("phone", utils.PhoneValidation)
	r.RedirectTrailingSlash = false

	r.GET("/", func(c *gin.Context) {
		versions := make(map[string]string)
		versions["1"] = "/v1"

		c.JSON(http.StatusOK, gin.H{
			"app": "Alertie API",
			"versions": versions,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		healty := true
		databaseHealth := true
		db := database.GetDB()

		if err := db.DB().Ping(); err != nil {
			healty = false
			databaseHealth = false
		}

		c.JSON(http.StatusOK, gin.H{
			"healty": healty,
			"database": databaseHealth,
		})
	})

	r.GET("/v1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"app": "Alertie API",
			"version": "1",
		})
	})

	v1 := r.Group("/v1")
	routeRegister(v1)

	r.Run()
}
