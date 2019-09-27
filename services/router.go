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

package services

import (
	"time"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func Register(router *gin.RouterGroup) {
	router.OPTIONS("/services", serviceOptions)
	router.GET("/services", serviceIndex)
	router.OPTIONS("/services/:id", serviceOptions)
	router.GET("/services/:id", serviceView)
	router.POST("/services", serviceAdd)
	router.PUT("/services/:id", serviceEdit)
	router.DELETE("/services/:id", serviceDelete)
}

func serviceOptions(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,DELETE,POST,PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}

func serviceIndex(c *gin.Context) {
	services, count, err := FindMany()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get services",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"services": services,
		"count": count,
	})
}

func serviceView(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad service id",
		})

		return
	}

	service, err := FindOne(&Service{ID: int64(id)})
	if err != nil || service.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Could not find service",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"service": service,
	})
}

func serviceAdd(c *gin.Context) {
	var service Service

	if err := c.BindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})

		return
	}

	token, err := uuid.NewV4()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token for service",
		})

		return
	}

	service.Token = token.String()

	service.CreatedAt = time.Now()
	service.UpdatedAt = time.Now()

	if err := SaveOne(&service); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save service",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"service": service,
	})
}

func serviceEdit(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad service id",
		})

		return
	}

	service, err := FindOne(Service{ID: int64(id)})
	if err != nil || service.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Could not find service",
		})

		return
	}

	if err := c.BindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})

		return
        }

	service.UpdatedAt = time.Now()

	if err := service.Update(&service); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update service",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"service": service,
	})
}

func serviceDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request id",
		})

		return
	}

	service, err := FindOne(Service{ID: int64(id)})
	if err != nil || service.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Could not find service",
		})

		return
	}

	if err := service.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete service",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
