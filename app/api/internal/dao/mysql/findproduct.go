package mysql

import (
	"database/sql"
	"dianshang/app/api/internal/model"
	"fmt"
)

func CountFindProduct(db *sql.DB, ProductName string) ([]model.AllProductList, error) {
	query := "SELECT * FROM AllProductList WHERE ProductName LIKE ?"
	fmt.Println(ProductName)
	rows, err := db.Query(query, "%"+ProductName+"%") // 确保查询语句中通配符%正确添加在参数前后用于模糊匹配
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.AllProductList
	for rows.Next() {
		var product model.AllProductList
		err := rows.Scan(&product.AllProductID, &product.ProductListName, &product.ProductID, &product.ProductName, &product.ProductNumber, &product.ProductPrice)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
