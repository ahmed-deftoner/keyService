package models

import (
	"errors"
	"html"
	"strings"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Key struct {
	Keyid     uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"key_id"`
	Uid       string    `gorm:"size:255;not null" json:"uid"`
	Service   string    `gorm:"size:255;not null" json:"service"`
	ApiKey    string    `gorm:"not null;unique" json:"api_key"`
	SecretKey string    `gorm:"not null;unique" json:"secret_key"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func Verify(hashedpassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
}

func (u *Key) BeforeSave() error {
	hashapi, err := Hash(u.ApiKey)
	if err != nil {
		return err
	}
	hashsecret, err := Hash(u.SecretKey)
	if err != nil {
		return err
	}
	u.ApiKey = string(hashapi)
	u.SecretKey = string(hashsecret)
	return nil
}

func (u *Key) Prepare() {
	u.ApiKey = html.EscapeString(strings.TrimSpace(u.ApiKey))
	u.SecretKey = html.EscapeString(strings.TrimSpace(u.SecretKey))
}

func (u *Key) Validate() error {
	if u.Uid == "" {
		return errors.New("uid required")
	}
	if u.Service == "" {
		return errors.New("service required")
	}
	if u.SecretKey == "" {
		return errors.New("secret key required")
	}
	if u.ApiKey == "" {
		return errors.New("api key required")
	}
	return nil
}

func (u *Key) SaveKey(db *gorm.DB) (*Key, error) {
	err := db.Debug().Create(&u).Error
	if err != nil {
		return &Key{}, err
	}
	return u, nil
}

func (u *Key) FindAllKeys(db *gorm.DB) (*[]Key, error) {
	Keys := []Key{}
	err := db.Debug().Model(&Key{}).Limit(100).Find(&Keys).Error
	if err != nil {
		return &[]Key{}, err
	}
	return &Keys, nil
}

func (u *Key) FindKeyById(db *gorm.DB, kid uuid.UUID) (*Key, error) {
	err := db.Debug().Model(Key{}).Where("keyid = ?", kid).Take(&u).Error
	if err != nil {
		return &Key{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Key{}, errors.New("Key not found")
	}
	return u, nil
}
