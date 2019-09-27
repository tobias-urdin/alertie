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
	"fmt"
	"github.com/tobias-urdin/alertie/database"
)

type Contact struct {
	ID		int64		`json:"id" binding:"-" validate:"-"`
	Firstname	string		`gorm:"type:varchar(64);not null" json:"firstname" form:"firstname" validate:"required,max=64"`
	Lastname	string		`gorm:"type:varchar(64);not null" json:"lastname" form:"lastname" validate:"required,max=64"`
	Email		string		`gorm:"type:varchar(64);not null" json:"email" form:"email" validate:"required,max=64,email"`
	Phone		string		`gorm:"type:varchar(32);not null" json:"phone" form:"phone" validate:"required,max=32,phone"`
	CreatedAt	time.Time	`json:"created_at" binding:"-" validate:"-"`
	UpdatedAt	time.Time	`json:"updated_at" binding:"-" validate:"-"`
}

func (c *Contact) Fullname() string {
	return fmt.Sprintf("%s %s", c.Firstname, c.Lastname)
}

func FindMany() ([]Contact, int, error) {
	db := database.GetDB()
	var models []Contact
	var count int

	tx := db.Begin()
	db.Model(&models).Count(&count)
	db.Find(&models)
	err := tx.Commit().Error

	return models, count, err
}

func FindOne(condition interface{}) (Contact, error) {
	db := database.GetDB()
	var model Contact

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

func (c *Contact) Update(data interface{}) error {
	db := database.GetDB()
	err := db.Model(c).Update(data).Error

	return err
}

func (c *Contact) Delete() error {
	db := database.GetDB()
	err := db.Where(Contact{
		ID: c.ID,
	}).Delete(Contact{}).Error

	return err
}
