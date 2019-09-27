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
	"time"
	"github.com/tobias-urdin/alertie/database"
	"github.com/tobias-urdin/alertie/services"
)

const (
	ALERT_OK = "ok"
	ALERT_WARNING = "warning"
	ALERT_CRITICAL = "critical"
)

type Alert struct {
	ID		int64		`json:"id" binding:"-" validate:"-"`
	ServiceID	int64		`gorm:"not null" json:"service_id" form:"service_id" validate:"required,max=36"`
	Service		services.Service
	Entity		string		`gorm:"type:varchar(64);not null" json:"entity" form:"entity" validate:"required,max=64"`
	Name		string		`gorm:"type:varchar(128);not null" json:"summary" form:"summary" validate:"required,max=128"`
	Data		string		`gorm:"type:text;not null" json:"description" form:"description" validate:"required"`
	State		string		`gorm:"type:varchar(32);not null" json:"state" binding:"state" validate:"required,max=32"`
	Acknowledged	bool		`gorm:"not null;default:0" json:"acknowledged" binding:"acknowledged" binding:"-" validate:"-"`
	SourceTrigger	string		`gorm:"type:varchar(32);not null" json:"source_trigger" binding:"-" validate:"max=32"`
	SourceIP	string		`gorm:"type:varchar(64)" json:"source_ip" binding:"-" validate:"-"`
	CreatedAt	time.Time	`json:"created_at" binding:"-" validate:"-"`
	UpdatedAt	time.Time	`json:"updated_at" binding:"-" validate:"-"`
}

func FindMany() ([]Alert, int, error) {
	db := database.GetDB()
	var models []Alert
	var count int

	tx := db.Begin()
	db.Model(&models).Count(&count)
	db.Find(&models)
	err := tx.Commit().Error

	return models, count, err
}

func FindOne(condition interface{}) (Alert, error) {
	db := database.GetDB()
	var model Alert

	tx := db.Begin()
	tx.Where(condition).First(&model)
	err := tx.Commit().Error

	return model, err
}

func SaveOne(data interface{}) error {
	db := database.GetDB()
	err := db.Save(data).Error

	return err
}

func (s *Alert) Update(data interface{}) error {
	db := database.GetDB()
	err := db.Model(s).Update(data).Error

	return err
}

func (s *Alert) Delete() error {
	db := database.GetDB()
	err := db.Where(Alert{
		ID: s.ID,
	}).Delete(Alert{}).Error

	return err
}
