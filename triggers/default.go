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
)

type DefaultTrigger struct {
	TriggerBase
}

type defaultTriggerData struct {
	State		string	`json:"state" binding:"required" validate:"required"`
	Entity		string	`json:"entity" binding:"-"`
	Name		string	`json:"name" binding:"required"`
	Data		string	`json:"data" binding:"required"`
}

func (t *DefaultTrigger) ProcessEvent(c *gin.Context, event *TriggerEvent) error {
	data := defaultTriggerData{}
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	if data.State != "warning" && data.State != "critical" && data.State != "acknowledged" && data.State != "ok" {
		return fmt.Errorf("Default trigger cannot handle state: %s", data.State)
	}

	event.State = data.State
	event.Entity = data.Entity
	event.Name = data.Name
	event.Data = data.Data

	return nil
}

func (t *DefaultTrigger) Response(c *gin.Context, event *TriggerEvent) error {
	c.JSON(http.StatusOK, gin.H{
		"state": event.State,
		"entity": event.Entity,
		"name": event.Name,
		"data": event.Data,
	})

	return nil
}
