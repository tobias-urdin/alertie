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

package triggers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"github.com/satori/go.uuid"
	"github.com/tobias-urdin/alertie/factory"
	"github.com/tobias-urdin/alertie/services"
	"github.com/tobias-urdin/alertie/utils"
)

func Register(router *gin.RouterGroup) {
	router.OPTIONS("/triggers", triggerOptions)
	router.POST("/triggers/:trigger/:token", triggerPost)
}

func triggerOptions(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}

func triggerPost(c *gin.Context) {
	triggerName := c.Param("trigger")
	token := c.Param("token")

	// TODO: Fix, ugly hack to make sure we can use factory here
	factory.Register("triggers", &thisFactory{})

	triggerPath := fmt.Sprintf("triggers.%s", triggerName)
	trigger := factory.New(triggerPath)
	if trigger == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Trigger %s (%s) does not exist", triggerName, triggerPath),
		})
		return
	}

	service, err := services.FindOne(&services.Service{Token: token})
	if err != nil || service.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Could not find service with that token",
			"details": err.Error(),
		})
		return
	}

	var event TriggerEvent
	event.Token = token
	event.Trigger = triggerName
	event.IP = c.ClientIP()

	if err := trigger.ProcessEvent(c, &event); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Trigger %s could not process event", triggerName),
				"details": utils.NewValidationError(err),
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Trigger %s could not process event", triggerName),
				"details": err.Error(),
			})
		}

		return
	}

	if event.Entity == "" {
		entuuid, err := uuid.NewV4()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to generate entity uuid",
				"details": err.Error(),
			})
			return
		}

		event.Entity = entuuid.String()
	}

	if err := trigger.Response(c, &event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Trigger %s could not create response", triggerName),
			"details": err.Error(),
		})
		return
	}
}
