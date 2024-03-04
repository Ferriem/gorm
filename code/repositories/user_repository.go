package repositories

import (
	"github.com/Ferriem/gorm/code/common"
	"github.com/Ferriem/gorm/code/datamodels"
	"gorm.io/gorm"
)

type UserRepository interface {
	Conn() error
	CreateUser(*datamodels.User) (uint, error)
	CreateUsers([]*datamodels.User) (int64, error)
	SelectByID(uint) (*datamodels.User, error)
	SelectAll() ([]*datamodels.User, error)
	Update(*datamodels.User) error
	Delete(uint) error
}

type UserManager struct {
	table string
	conn  *gorm.DB
}

func NewUserManager(tableName string, conn *gorm.DB) UserRepository {
	return &UserManager{
		table: tableName,
		conn:  conn,
	}
}

func (u *UserManager) Conn() error {
	if u.conn == nil {
		db, err := common.NewConn()
		if err != nil {
			return err
		}
		u.conn = db
	}
	if u.table == "" {
		u.table = "users"
	}
	return nil
}

func (u *UserManager) CreateUser(user *datamodels.User) (uint, error) {
	if err := u.Conn(); err != nil {
		return 0, err
	}
	result := u.conn.Create(user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (u *UserManager) CreateUsers(users []*datamodels.User) (int64, error) {
	if err := u.Conn(); err != nil {
		return 0, err
	}
	result := u.conn.Session(&gorm.Session{SkipHooks: true}).Create(users)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (u *UserManager) SelectByID(id uint) (*datamodels.User, error) {
	if err := u.Conn(); err != nil {
		return &datamodels.User{}, err
	}
	var user datamodels.User
	if err := u.conn.Preload("CreditCard").First(&user, id).Error; err != nil {
		return &datamodels.User{}, err
	}

	return &user, nil
}

func (u *UserManager) SelectAll() ([]*datamodels.User, error) {
	if err := u.Conn(); err != nil {
		return nil, err
	}
	var users []*datamodels.User
	if err := u.conn.Preload("CreditCard").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Only use save, won't judge whether the user is existed.
func (u *UserManager) Update(user *datamodels.User) error {
	if err := u.Conn(); err != nil {
		return err
	}
	result := u.conn.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserManager) Delete(id uint) error {
	if err := u.Conn(); err != nil {
		return err
	}
	result := u.conn.Delete(&datamodels.User{CreditCard: datamodels.CreditCard{UserID: id}, ID: id})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
