package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Exchanges struct {
	Id       uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"id"`
	Name     string    `gorm:"size:255;not null" json:"name"`
	ImageSrc string    `gorm:"not null;unique" json:"image_src"`
	Short    string    `gorm:"not null;unique" json:"short"`
}

func (e *Exchanges) ValidateExchange() error {
	if e.Name == "" {
		return errors.New("name required")
	}
	if e.ImageSrc == "" {
		return errors.New("image src required")
	}
	if e.Short == "" {
		return errors.New("short alias required")
	}
	return nil
}

func (e *Exchanges) SaveExchange(db *gorm.DB) (*Exchanges, error) {
	err := db.Debug().Create(&e).Error
	if err != nil {
		return &Exchanges{}, err
	}
	return e, nil
}

func (e *Exchanges) FindAllExchanges(db *gorm.DB) (*[]Exchanges, error) {
	Exchange := []Exchanges{}
	err := db.Debug().Model(&Exchanges{}).Limit(100).Find(&Exchange).Error
	if err != nil {
		return &[]Exchanges{}, err
	}
	return &Exchange, nil
}
