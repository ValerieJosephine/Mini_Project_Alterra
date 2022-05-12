package models

import (
	db "MINI_PROJECT_ALTERRA/database"

	"net/http"
)

type product struct {
	Id_product          int    `json:"id"`
	description_product string `json:"description"`
	status              string `json:"status"`
	restock_at          int    `json:"restock"`
	price               int    `json:"price"`
	updated_at          int    `json:"updated"`
}

func FetchAllproduct() (Response, error) {
	var obj product
	var arrobj []product
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM product"

	rows, err := con.Query(sqlStatement)
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

func StoreProduct(Id_product int, description_product string, status string, restock_at int, price int, updated_at int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT product (Id_product, description_product, status, restock_at, price, updated_at) VALUES (?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Id_product, description_product, status, restock_at, price, updated_at)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
}
