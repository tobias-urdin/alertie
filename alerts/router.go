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

package alerts

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup) {
	router.OPTIONS("/alerts", alertOptions)
	router.GET("/alerts", alertIndex)
	router.OPTIONS("/alerts/:id", alertOptions)
	router.GET("/alerts/:id", alertView)
	router.DELETE("/alerts/:id", alertDelete)
}

func alertOptions(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}

func alertIndex(c *gin.Context) {
	alerts, count, err := FindMany()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get alerts",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"alerts": alerts,
		"count": count,
	})
}

func alertView(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad alert id",
		})

		return
	}

	alert, err := FindOne(&Alert{ID: int64(id)})
	if err != nil || alert.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Could not find alert",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"alert": alert,
	})
}

func alertDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request id",
		})

		return
	}

	alert, err := FindOne(Alert{ID: int64(id)})
	if err != nil || alert.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Could not find alert",
		})

		return
	}

	if err := alert.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete alert",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
