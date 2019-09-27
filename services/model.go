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
	"github.com/tobias-urdin/alertie/database"
)

type Service struct {
	ID		int64		`json:"id" binding:"-" validate:"-"`
	Name		string		`gorm:"type:varchar(64);not null;unique" json:"name" form:"name" validate:"required,max=64"`
	Description	string		`gorm:"type:text" json:"description" form:"description" validate:"required"`
	Token		string		`gorm:"type:varchar(36);not null;index:idx_service_token" json:"token" binding:"-" validate:"-"`
	Maintenance	bool		`gorm:"not null" json:"maintenance" form:"maintenance" binding:"-" validate:"-"`
	//UserID		int64		`gorm:"type:int;not null" json:"user_id" binding:"-" validate:"-"`
	//User		users.User
	CreatedAt	time.Time	`json:"created_at" binding:"-" validate:"-"`
	UpdatedAt	time.Time	`json:"updated_at" binding:"-" validate:"-"`
}

func FindMany() ([]Service, int, error) {
	db := database.GetDB()
	var models []Service
	var count int

	tx := db.Begin()
	db.Model(&models).Count(&count)
	db.Find(&models)
	err := tx.Commit().Error

	return models, count, err
}

func FindOne(condition interface{}) (Service, error) {
	db := database.GetDB()
	var model Service

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

func (s *Service) Update(data interface{}) error {
	db := database.GetDB()
	err := db.Model(s).Update(data).Error

	return err
}

func (s *Service) Delete() error {
	db := database.GetDB()
	err := db.Where(Service{
		ID: s.ID,
	}).Delete(Service{}).Error

	return err
}
