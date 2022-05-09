package models

import (
	"MINI_PROJECT_ALTERRA/db"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
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
	obj := product
	arrobj := []product
	res := Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM product"

	rows,err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return ress,err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_product, &obj.description_product, &obj.status, &obj.restock_at, &obj.price, &obj.updated_at)
		if err != nil{
			return res,err
		}

		arrobj = append(arrobj,obj)
	}
}

res.Status = http.StatusOK
res.Message = "Success"
res.Data = arrobj

return res,nil
