package mysql

import (
	"database/sql"
	"dianshang/app/api/global"
	"dianshang/app/api/internal/model"
	"fmt"
	"go.uber.org/zap"
)

func UpdateAllProductList(db *sql.DB) error {
	DropList(db, "AllProductList")
	MakeAllProductList(db)
	// 查询所有用户账户
	rows, err := db.Query("SELECT UserAccount FROM AccountList WHERE Identity = 'boss'")
	if err != nil {
		global.Logger.Error("无法查询用户账户", zap.Error(err))
		return fmt.Errorf("无法查询用户账户: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var userAccount string
		err := rows.Scan(&userAccount)
		if err != nil {
			global.Logger.Error("扫描用户账户失败", zap.Error(err))
			continue
		}

		// 获取用户商品表格中的数据
		productRows, err := db.Query("SELECT ProductID, ProductName, ProductNumber, ProductPrice FROM " + userAccount)
		if err != nil {
			global.Logger.Error("无法查询用户商品表格", zap.String("UserAccount", userAccount), zap.Error(err))
			continue
		}
		defer productRows.Close()

		// 将商品表格中的数据插入到 AllProductList 表中
		for productRows.Next() {
			var productID int
			var productName string
			var productNumber int
			var productPrice float64
			err := productRows.Scan(&productID, &productName, &productNumber, &productPrice)
			if err != nil {
				global.Logger.Error("扫描商品表格失败", zap.String("UserAccount", userAccount), zap.Error(err))
				continue
			}

			// 插入数据到 AllProductList 表中
			insertQuery := "INSERT INTO AllProductList (ProductListName, ProductID, ProductName, ProductNumber, ProductPrice) VALUES (?, ?, ?, ?, ?)"
			_, err = db.Exec(insertQuery, userAccount, productID, productName, productNumber, productPrice)
			if err != nil {
				global.Logger.Error("无法插入商品数据到 AllProductList", zap.String("UserAccount", userAccount), zap.Error(err))
				continue
			}
		}
	}

	global.Logger.Info("已完成填充 AllProductList 表")
	return nil
}
func ShowAllProductList(db *sql.DB) error {
	rows, err := db.Query("SELECT AllProductID, ProductListName, ProductID, ProductName, ProductNumber, ProductPrice FROM AllProductList")

	if err != nil {
		return fmt.Errorf("无法从 AllProductList 中获取数据: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var allProductID, productID, productNumber int
		var productListName, productName string
		var productPrice float64
		err := rows.Scan(&allProductID, &productListName, &productID, &productName, &productNumber, &productPrice)
		if err != nil {
			return fmt.Errorf("无法扫描行: %v", err)
		}

		fmt.Printf("AllProductID: %d, ProductListName: %s, ProductID: %d, ProductName: %s, ProductNumber: %d, ProductPrice: %f\n", allProductID, productListName, productID, productName, productNumber, productPrice)
	}

	return nil
}

func CoutAllProductList(db *sql.DB) ([]model.AllProductList, error) {
	// 声明一个用于存储所有产品列表数据的切片
	var allProductLists []model.AllProductList

	// 查询AllProductList表中的所有数据
	rows, err := db.Query("SELECT AllProductID, ProductListName, ProductID, ProductName, ProductNumber, ProductPrice FROM AllProductList")

	if err != nil {
		global.Logger.Error("无法查询AllProductList表", zap.Error(err))
		return nil, fmt.Errorf("无法查询AllProductList表: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var allProductID int
		var productListName string
		var productID int
		var productName string
		var productNumber int
		var productPrice float64
		err := rows.Scan(&allProductID, &productListName, &productID, &productName, &productNumber, &productPrice)
		if err != nil {
			global.Logger.Error("扫描AllProductList表失败", zap.Error(err))
			continue
		}

		// 创建一个新的 AllProductList 对象并添加到切片中
		allProductLists = append(allProductLists, model.AllProductList{
			AllProductID:    allProductID,
			ProductListName: productListName,
			ProductID:       productID,
			ProductName:     productName,
			ProductNumber:   productNumber,
			ProductPrice:    productPrice,
		})
	}

	return allProductLists, nil
}
