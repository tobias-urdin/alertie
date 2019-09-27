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

type GrafanaTrigger struct {
	TriggerBase
}

type grafanaTriggerData struct {
	RuleID		int64	`json:"ruleId" binding:"required"`
	RuleName	string	`json:"ruleName" binding:"required"`
	State		string	`json:"state" binding:"required"`
	Message		string	`json:"message" binding:"required"`
}

func (t *GrafanaTrigger) ProcessEvent(c *gin.Context, event *TriggerEvent) error {
	data := grafanaTriggerData{}
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	if data.State == "alerting" {
		event.State = "critical"
	} else if data.State == "ok" {
		event.State = "ok"
	} else {
		return fmt.Errorf("Grafana trigger cannot handle state: %s", data.State)
	}

	event.Entity = fmt.Sprintf("%s/%s", data.RuleID, data.RuleName)
	event.Name = data.RuleName
	event.Data = data.Message

	return nil
}

func (t *GrafanaTrigger) Response(c *gin.Context, event *TriggerEvent) error {
	c.JSON(http.StatusOK, gin.H{
		"state": event.State,
		"entity": event.Entity,
		"name": event.Name,
		"data": event.Data,
	})

	return nil
}
