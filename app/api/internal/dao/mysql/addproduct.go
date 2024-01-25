package mysql

import (
	"database/sql"
	"dianshang/app/api/global"
	"dianshang/app/api/internal/model"
	"fmt"
	"go.uber.org/zap"
)

func AddProduct(db *sql.DB, userAccount string, productName string, productNumber int, productPrice float64) error {
	query := `
		INSERT INTO ` + userAccount + ` (ProductName, ProductNumber, ProductPrice) VALUES (?, ?, ?)
	`
	_, err := db.Exec(query, productName, productNumber, productPrice)
	if err != nil {
		global.Logger.Error("无法添加商品到商品库", zap.Error(err))
		return fmt.Errorf("无法添加商品到商品库")
	}
	global.Logger.Info("成功添加商品到商品库：" + userAccount)
	return nil
}
func CoutProductList(db *sql.DB, userAccount string) ([]model.ProductList, error) {
	query := "SELECT ProductID, ProductName, ProductNumber, ProductPrice FROM " + userAccount
	rows, err := db.Query(query)
	if err != nil {
		global.Logger.Error("查询商品库失败", zap.Error(err))
		return nil, fmt.Errorf("查询商品库失败")
	}
	defer rows.Close()

	var products []model.ProductList
	for rows.Next() {
		var product model.ProductList
		err := rows.Scan(&product.ProductID, &product.ProductName, &product.ProductNumber, &product.ProductPrice)
		if err != nil {
			global.Logger.Error("读取商品数据失败", zap.Error(err))
			return nil, fmt.Errorf("读取商品数据失败")
		}
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		global.Logger.Error("获取商品数据出错", zap.Error(err))
		return nil, fmt.Errorf("获取商品数据出错")
	}

	return products, nil
}
func ShowProductList(db *sql.DB, userAccount string) {
	products, err := CoutProductList(db, userAccount)
	if err != nil {
		fmt.Println("查询商品库失败:", err)
		return
	}
	for _, product := range products {
		fmt.Printf("ProductID: %d, ProductName: %s, ProductNumber: %d, ProductPrice: %.2f\n", product.ProductID, product.ProductName, product.ProductNumber, product.ProductPrice)
	}
}
