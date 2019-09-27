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

package contacts

import (
	"time"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup) {
	router.OPTIONS("/contacts", contactOptions)
	router.GET("/contacts", contactIndex)
	router.OPTIONS("/contacts/:id", contactOptions)
	router.GET("/contacts/:id", contactView)
	router.POST("/contacts", contactAdd)
	router.PUT("/contacts/:id", contactEdit)
	router.DELETE("/contacts/:id", contactDelete)
}

func contactOptions(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,DELETE,POST,PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}

func contactIndex(c *gin.Context) {
	contacts, count, err := FindMany()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get contacts",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"contacts": contacts,
		"count": count,
	})
}

func contactView(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad contact id",
		})

		return
	}

	contact, err := FindOne(&Contact{ID: int64(id)})
	if err != nil || contact.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Could not find contact",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"contact": contact,
	})
}

func contactAdd(c *gin.Context) {
	var contact Contact

	if err := c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})

		return
	}

	contact.CreatedAt = time.Now()
	contact.UpdatedAt = time.Now()

	if err := SaveOne(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save contact",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"contact": contact,
	})
}

func contactEdit(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad contact id",
		})

		return
	}

	contact, err := FindOne(Contact{ID: int64(id)})
	if err != nil || contact.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Could not find contact",
		})

		return
	}

	if err := c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})

		return
        }

	contact.UpdatedAt = time.Now()

	if err := contact.Update(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update contact",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"contact": contact,
	})
}

func contactDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request id",
		})

		return
	}

	contact, err := FindOne(Contact{ID: int64(id)})
	if err != nil || contact.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Could not find contact",
		})

		return
	}

	if err := contact.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete contact",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
