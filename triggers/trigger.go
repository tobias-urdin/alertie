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
	"github.com/gin-gonic/gin"
	"github.com/tobias-urdin/alertie/factory"
)

type TriggerEvent struct {
	State	string	`json:"state"`
	Token	string	`json:"token"`
	Entity	string	`json:"entity"`
	Name	string	`json:"name"`
	Data	string	`json:"data"`
	Trigger	string	`json:"trigger"`
	IP	string	`json:"ip"`
}

type TriggerBase interface {
	ProcessEvent(*gin.Context, *TriggerEvent) error
	Response(*gin.Context, *TriggerEvent) error
}

type thisFactory struct {}

func (t *thisFactory) New(name string) interface{} {
	switch name {
	case "default":
		return &DefaultTrigger{}
	case "grafana":
		return &GrafanaTrigger{}
	}
	return nil
}

func init() {
	factory.Register("triggers", &thisFactory{})
}
