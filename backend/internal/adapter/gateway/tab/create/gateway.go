package gateway

import (
	"fmt"

	model "github.com/asaringo99/task_management/internal/adapter/gateway/tab"
	"gorm.io/gorm"
)

type TabCreateGateway struct {
	db *gorm.DB
}

func (c *TabCreateGateway) Create(input model.TabModel) (model.TabModel, error) {
	tab := model.TabModel{
		Id:       0,
		Userid:   input.Userid,
		IsActive: input.IsActive,
		Title:    input.Title,
	}
	fmt.Println(tab)
	fmt.Println("das")
	c.db.Table("tabs").Create(&tab)
	return tab, nil
}

func NewTabCreateGateway(db *gorm.DB) TabCreateGatewayInterface {
	return &TabCreateGateway{
		db: db,
	}
}
