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

package alerters

import (
	"github.com/tobias-urdin/alertie/factory"
	"gopkg.in/ini.v1"
)

type AlerterBase interface {
	Init(*ini.Section) error
	Alert() (bool, error)
}

type thisFactory struct {}

func (t *thisFactory) New(name string) interface{} {
	switch name {
	case "email":
		return &EmailAlerter{}
	}
	return nil
}

func init() {
        factory.Register("alerters", &thisFactory{})
}
