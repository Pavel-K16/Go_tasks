package entities

type ProductCategory struct {
	Id          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:VARCHAR(25);not null"`
	Description string `gorm:"type:text;column:description"`
}

func (item ProductCategory) NotNull() map[string]interface{} {
	toupdate := make(map[string]interface{})

	if item.Name != "" {
		toupdate["name"] = item.Name
	}
	if item.Description != "" {
		toupdate["description"] = item.Description
	}
	return toupdate
}

type Product struct {
	Id          uint    `gorm:"primaryKey;autoIncrement"`
	Name        string  `gorm:"type:VARCHAR(25);not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"type:Decimal(15,3)"`
	CategoryId  int     `gorm:"foreignKey; references: productCategory; Id;column:categoryid"`
}

func (item Product) NotNull() map[string]interface{} {
	toupdate := make(map[string]interface{})

	if item.Name != "" {
		toupdate["name"] = item.Name
	}
	if item.Description != "" {
		toupdate["description"] = item.Description
	}
	if item.Price != 0 {
		toupdate["price"] = item.Price
	}
	if item.CategoryId != 0 {
		toupdate["categoryid"] = item.CategoryId
	}
	return toupdate
}

type Table interface {
	NotNull() map[string]interface{}
}
