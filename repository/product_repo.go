package repository

import (
	"MINI_PROJECT_ALTERRA/models"
	// "gorm.io/gorm"
)

// type ProductRepository interface {
// 	FetchProducts() ([]models.Product, error)
// 	StoreProduct(input models.Product) (models.Product, error)
// }

// type Repository struct {
// 	db *gorm.DB
// }

// func NewRepository(db *gorm.DB) *Repository {
// 	return &Repository{db: db}

// }

func (r *Repository) FetchProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil

}

func (r *Repository) FetchProductById(id int) error {
	err := r.db.Where("id = ?").Find(&models.Product{}).Error
	if err != nil {
		return err
	}
	return err

}

func (r *Repository) StoreProduct(input models.Product) (models.Product, error) {
	err := r.db.Create(&input).Error
	if err != nil {
		return input, err
	}

	return input, nil
}

func (r *Repository) DeleteProduct(id int) error {
	err := r.db.Where("id = ?").Delete(&models.Product{}).Error
	if err != nil {
		return err
	}
	return err

}

func (r *Repository) UpdateProduct([]models.Product) error {
	var products []models.Product
	err := r.db.Model(&products).Updates(map[string]interface{}{"description": "?", "status": "?", "restock": '?', "updated": '?'}).Error
	if err != nil {
		return err
	}
	return err

}

// func FetchAllproduct() (Response, error) {
// 	var obj Product
// 	var arrobj []Product
// 	var res Response

// 	con := db.GetConnection()

// 	sqlStatement := "SELECT * FROM product"
// 	rows, err := con.Raw(sqlStatement).Rows()
// 	defer rows.Close()

// 	if err != nil {
// 		return res, err
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(&obj.Id_product, &obj.description_product, &obj.status, &obj.restock_at, &obj.price, &obj.updated_at)
// 		if err != nil {
// 			return res, err
// 		}

// 		// arrobj = append(arrobj, obj)
// 	}

// 	arrobj = append(arrobj, obj)

// 	res.Status = http.StatusOK
// 	res.Message = "Success"
// 	res.Data = arrobj

// 	return res, nil

// }
