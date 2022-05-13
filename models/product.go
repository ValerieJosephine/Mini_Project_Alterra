package models

import (
	db "MINI_PROJECT_ALTERRA/database"
	"errors"
	"fmt"

	"net/http"

	"gorm.io/gorm"
)

//>menulis tipe data
type Product struct {
	Id_product          int    `json:"id"`
	description_product string `json:"description"`
	status              string `json:"status"`
	restock_at          int    `json:"restock"`
	price               int    `json:"price"`
	updated_at          int    `json:"updated"`
}

// type AProductRepo interface {
// 	FetchAllproduct() ([]product.ProductCostum, error)
// }

type ProductRepo struct {
	db *gorm.DB
}

// >untuk get product dari db
func FetchAllproduct() (Response, error) {
	var obj Product
	var arrobj []Product
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM product"
	rows, err := con.Raw(sqlStatement).Rows()
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_product, &obj.description_product, &obj.status, &obj.restock_at, &obj.price, &obj.updated_at)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil

}

// >untuk fungsi post produk
func StoreProduct(Id_product int, description_product string, status string, restock_at int, price int, updated_at int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT product (Id_product, description_product, status, restock_at, price, updated_at) VALUES (?,?,?,?,?,?)"

	stmt := con.Exec(sqlStatement)
	if stmt.Error != nil {
		return res, stmt.Error
	}

	// result, err := stmt.Exec(Id_product, description_product, status, restock_at, price, updated_at)
	// if err != nil {
	// 	return result, err
	// }

	res.Status = http.StatusOK
	res.Message = "Success"
}

//> fungsi update data product
func UpdateProduct(Id_product int, description_product string, status string, restock_at int, price int, updated_at int) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "UPDATE product SET description_product = ?, status = ?, restock_at = ?, price = ?, updated_at = ? WHERE id = ?"

	// stmt, err := con.Prepare(sqlStatement)
	// if err != nil {
	// 	return res, err
	// }

	stmt := con.Exec(sqlStatement)
	if stmt.Error != nil {
		return res, stmt.Error
	} else if stmt.RowsAffected == 0 {
		errorText := "data tidak ter update"
		fmt.Println(errorText)
		return res, errors.New(errorText)
	}

	// result, err := stmt.Exec(description_product, status, restock_at, price, updated_at, Id_product)
	// if err != nil {
	// 	return result, err
	// }
	// rowsAffected, err := result.RowsAffected()
	// if err != nil {
	// 	return result, err
	// 	}
	// }

	res.Status = http.StatusOK
	res.Message = "Success"

	return res, nil
}

//> fungsi delete data product
func DeleteProduct(Id_product int) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "DELETE FROM product WHERE id = ?"

	// stmt, err := con.Prepare(sqlStatement)
	// if err != nil {
	// 	return res, err
	// }
	stmt := con.Exec(sqlStatement)
	if stmt.Error != nil {
		return res, stmt.Error
	} else if stmt.RowsAffected == 0 {
		errorText := "data tidak ter update"
		fmt.Println(errorText)
		return res, errors.New(errorText)
	}

	// result, err := stmt.Exec(Id_product)
	// if err != nil {
	// 	return res, err
	// }
	// rowsAffected, err := result.RowsAffected()
	// if err != nil {
	// 	return res, err
	// }

	res.Status = http.StatusOK
	res.Message = "Success"

	return res, nil
}
