package datamodels

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type CreditCard struct {
	Number string
	UserID uint
}

type User struct {
	ID           uint           `sql:"id"`            // Standard field for the primary key
	Name         string         `sql:"name"`          // A regular string field
	Email        *string        `sql:"email"`         // A pointer to a string, allowing for null values
	Age          uint8          `sql:"age"`           // An unsigned 8-bit integer
	Birthday     *time.Time     `sql:"birthday"`      // A pointer to time.Time, can be null
	MemberNumber sql.NullString `sql:"member_number"` // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   `sql:"activated_at"`  // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      `sql:"create_at"`     // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      `sql:"update_at"`     // Automatically managed by GORM for update time
	CreditCard   CreditCard     `sql:"credit_card"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Name == "Jack" {
		return errors.New("invalid role")
	}
	return
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("BeforeDelete")
	tx.Model(u).Association("CreditCard").Find(&u.CreditCard)
	tx.Where("user_id = ?", u.ID).Delete(u.CreditCard)
	return
}
