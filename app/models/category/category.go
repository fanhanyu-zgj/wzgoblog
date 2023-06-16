package category

import (
	"strconv"
	"wz/app/models"
	"wz/pkg/route"
)

type Category struct {
	models.BaseModel

	Name string `gorm:"type:varchar(255);not null;" valid:"name"`
}

func (category Category) Link() string {
	return route.RouteName2URL("categories.show", "id", strconv.FormatUint(category.ID, 10))
}
