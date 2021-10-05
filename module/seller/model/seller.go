package model

import (
	"strings"

	"github.com/caohieu04/real-estate/common"
)

const EntityName = "Seller"

type Seller struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"address" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (Seller) TableName() string {
	return "sellers"
}

type SellerUpdate struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"address" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (SellerUpdate) TableName() string {
	return Seller{}.TableName()
}

type SellerCreate struct {
	ID    int            `json:"id" gorm:"column:id;"`
	Name  string         `json:"name" gorm:"column:name;"`
	Addr  string         `json:"address" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (SellerCreate) TableName() string {
	return Seller{}.TableName()
}

func (p *SellerCreate) Validate() error {
	p.Name = strings.TrimSpace(p.Name)

	if len(p.Name) == 0 {
		return ErrNameCannotBeEmpty
	}

	return nil
}

var (
	ErrNameCannotBeEmpty = common.NewCustomError(nil, "seller name can't be blank", "ErrNameCannotBeEmpty")
)
