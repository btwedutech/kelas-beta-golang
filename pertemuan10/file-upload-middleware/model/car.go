package model

import "gorm.io/gorm"

type Car struct {
	Model
	Nama         string `gorm:"not null" json:"nama"`
	Tipe         string `gorm:"not null" json:"tipe"`
	Tahun        string `gorm:"not null" json:"tahun"`
	Color        string `gorm:"not null" json:"color"`
	Condition    string `gorm:"not null" json:"condition"`
	UUID         string `gorm:"not null" json:"uuid"`
	SellingPrice int    `gorm:"not null;default:0" json:"selling_price"`
}

func (cr *Car) Create(db *gorm.DB) error {
	err := db.
		Model(Car{}).
		Create(&cr).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *Car) GetByID(db *gorm.DB) (Car, error) {
	res := Car{}

	err := db.
		Model(Car{}).
		Where("id = ?", cr.Model.ID).
		Take(&res).
		Error

	if err != nil {
		return Car{}, err
	}

	return res, nil
}

func (cr *Car) GetBySpecific(db *gorm.DB) (Car, error) {
	res := Car{}

	query := db.Model(Car{})

	if cr.UUID != "" {
		query = query.Where("uuid = ?", cr.UUID)
	}

	// if cr.Tipe != "" {
	// 	query = query.Where("tipe = ?", cr.Tipe)
	// }

	// if cr.Tahun != "" {
	// 	query = query.Where("tahun = ?", cr.Tahun)
	// }

	// if cr.Color != "" {
	// 	query = query.Where("color = ?", cr.Color)
	// }

	// if cr.Condition != "" {
	// 	query = query.Where("condition = ?", cr.Condition)
	// }

	err := query.Take(&res).Error

	if err != nil {
		return Car{}, err
	}

	return res, nil
}

func (cr *Car) GetAll(db *gorm.DB) ([]Car, error) {
	res := []Car{}

	err := db.
		Model(Car{}).
		Find(&res).
		Error

	if err != nil {
		return []Car{}, err
	}

	return res, nil
}

func (cr *Car) UpdateOneByID(db *gorm.DB) error {
	err := db.
		Model(Car{}).
		Select("nama", "tipe", "tahun", "color",
			"condition", "uuid", "selling_price").
		Where("id = ?", cr.Model.ID).
		Updates(map[string]any{
			"nama":          cr.Nama,
			"tipe":          cr.Tipe,
			"tahun":         cr.Tahun,
			"color":         cr.Color,
			"condition":     cr.Condition,
			"uuid":          cr.UUID,
			"selling_price": cr.SellingPrice,
		}).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *Car) DeleteByID(db *gorm.DB) error {
	err := db.
		Model(Car{}).
		Where("id = ?", cr.Model.ID).
		Delete(&cr).
		Error

	if err != nil {
		return err
	}

	return nil
}
