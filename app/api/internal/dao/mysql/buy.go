package mysql

import (
	"database/sql"
	"dianshang/app/api/global"
	"dianshang/app/api/internal/model"
	"fmt"
	"go.uber.org/zap"
)

func BuyMakeCar(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS Car (
		    ID INT AUTO_INCREMENT PRIMARY KEY,
			AllProductID INT,
			ProductListName VARCHAR(50) NOT NULL,
			ProductID INT,
			ProductName VARCHAR(100) NOT NULL,
			ProductNumber INT NOT NULL,
			ProductPrice DOUBLE NOT NULL
		);
	`

	_, err := db.Exec(query)
	if err != nil {
		global.Logger.Fatal("创建购物车失败," + err.Error())
		return err
	}
	global.Logger.Info("已成功创建购物车")
	return nil
}
func ShowCarList(db *sql.DB) error {
	rows, err := db.Query("SELECT ID, AllProductID, ProductListName, ProductID, ProductName, ProductNumber, ProductPrice FROM Car")
	if err != nil {
		return fmt.Errorf("无法从购物车中获取数据: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, allProductID, productID, productNumber int
		var productListName, productName string
		var productPrice float64
		err := rows.Scan(&id, &allProductID, &productListName, &productID, &productName, &productNumber, &productPrice)
		if err != nil {
			return fmt.Errorf("无法扫描行: %v", err)
		}

		fmt.Printf("ID: %d, AllProductID: %d, ProductListName: %s, ProductID: %d, ProductName: %s, ProductNumber: %d, ProductPrice: %f\n", id, allProductID, productListName, productID, productName, productNumber, productPrice)
	}

	return nil
}
func BuyAddCar(db *sql.DB, BuyAllProductID int, BuyProductNumber int, BuyUserAccount string) error {
	// 查询商品信息
	var (
		AllProductID    int
		ProductListName string
		ProductID       int
		ProductName     string
		ProductNumber   int
		ProductPrice    float64
	)
	err := db.QueryRow("SELECT AllProductID, ProductListName, ProductID, ProductName, ProductNumber, ProductPrice FROM AllProductList WHERE AllProductID = ?", BuyAllProductID).
		Scan(&AllProductID, &ProductListName, &ProductID, &ProductName, &ProductNumber, &ProductPrice)
	if err != nil {
		global.Logger.Fatal("查询商品数据失败," + err.Error())
		return err
	}

	// 查询购买用户账户信息
	var (
		UserAccount string
		Password    string
		ID          int
		Nickname    string
		Identity    string
		Balance     float64
	)
	err = db.QueryRow("SELECT UserAccount, Password, ID, Nickname, Identity, Balance FROM AccountList WHERE UserAccount = ?", BuyUserAccount).
		Scan(&UserAccount, &Password, &ID, &Nickname, &Identity, &Balance)
	if err != nil {
		global.Logger.Fatal("查询用户账户数据失败," + err.Error())
		return err
	}
	if BuyProductNumber > ProductNumber {
		global.Logger.Fatal("购买量需求过大")
		return err
	}
	if Identity != "customer" {
		global.Logger.Fatal("非顾客无购买权限")
		return err
	}
	if Balance < (ProductPrice * float64(BuyProductNumber)) {
		global.Logger.Fatal("账户金额不足")
		return err
	}

	// 执行更新操作
	_, err = db.Exec("UPDATE AccountList SET Balance = ? WHERE UserAccount = ?", Balance-(ProductPrice*float64(BuyProductNumber)), BuyUserAccount)
	if err != nil {
		global.Logger.Fatal("更新买家账户余额失败," + err.Error())
		return err
	}

	// 执行更新卖家用户商品信息
	_, err = db.Exec("UPDATE "+ProductListName+" SET ProductNumber = ? WHERE ProductID = ?", ProductNumber-BuyProductNumber, ProductID)
	if err != nil {
		global.Logger.Fatal("更新卖家商品信息失败," + err.Error())
		return err
	}
	err = UpdateAllProductList(global.MysqlDB)
	if err != nil {
		global.Logger.Fatal("更新总商品信息失败," + err.Error())
		return err
	}

	// 将购买的商品插入到购物车列表
	_, err = db.Exec("INSERT INTO Car (AllProductID, ProductListName, ProductID, ProductName, ProductNumber, ProductPrice) VALUES (?, ?, ?, ?, ?, ?)",
		AllProductID, ProductListName, ProductID, ProductName, BuyProductNumber, ProductPrice)
	if err != nil {
		global.Logger.Fatal("插入购物车数据失败," + err.Error())
		return err
	}

	global.Logger.Info("购物车添加成功")
	return nil
}

func MakeHostList(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS HostList (
			ProductName VARCHAR(100) NOT NULL,
			SoldNumber INT NOT NULL,
			Price DOUBLE NOT NULL
		);
	`
	_, err := db.Exec(query)
	if err != nil {
		global.Logger.Fatal("create table failed," + err.Error())
		return fmt.Errorf("无法创建HostList表")
	}
	global.Logger.Info("HostList表创建成功")
	return nil
}

func SortHotList(db *sql.DB) ([]model.HostList, error) {
	rows, err := db.Query("SELECT ProductName, SoldNumber, Price FROM HostList ORDER BY ProductName, Price")
	if err != nil {
		return nil, fmt.Errorf("无法从 HostList 获取排序的数据: %v", err)
	}
	defer rows.Close()

	var HotList []model.HostList
	var prevItem model.HostList
	duplicateIndex := -1

	for rows.Next() {
		var item model.HostList
		err := rows.Scan(&item.ProductName, &item.SoldNumber, &item.Price)
		if err != nil {
			return nil, fmt.Errorf("无法扫描行: %v", err)
		}

		if item.ProductName == prevItem.ProductName && item.Price == prevItem.Price {
			HotList[duplicateIndex].SoldNumber++
		} else {
			HotList = append(HotList, item)
			prevItem = item
			duplicateIndex++
		}
	}

	return HotList, nil
}

func PrintHotList(db *sql.DB) {
	hotList, err := SortHotList(db)
	if err != nil {
		fmt.Println("无法打印热门商品清单:", err)
		return
	}

	for _, item := range hotList {
		fmt.Printf("产品名称: %s, 销售数量: %d, 价格: %f\n", item.ProductName, item.SoldNumber, item.Price)
	}
}

func CountCarList(db *sql.DB) ([]model.Car, error) {
	// 声明一个用于存储购物车列表数据的切片
	var carList []model.Car

	// 查询Car表中的所有数据
	rows, err := db.Query("SELECT ID, AllProductID, ProductListName, ProductID, ProductName, ProductNumber, ProductPrice FROM Car")

	if err != nil {
		global.Logger.Error("无法查询Car表", zap.Error(err))
		return nil, fmt.Errorf("无法查询Car表: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var allProductID int
		var productListName string
		var productID int
		var productName string
		var productNumber int
		var productPrice float64
		err := rows.Scan(&id, &allProductID, &productListName, &productID, &productName, &productNumber, &productPrice)
		if err != nil {
			global.Logger.Error("扫描Car表失败", zap.Error(err))
			continue
		}

		// 创建一个新的 Car 对象并添加到切片中
		carList = append(carList, model.Car{
			Id:              id,
			AllProductID:    allProductID,
			ProductListName: productListName,
			ProductID:       productID,
			ProductName:     productName,
			ProductNumber:   productNumber,
			ProductPrice:    productPrice,
		})
	}

	return carList, nil
}
func AddProductListNameBalance(db *sql.DB, ProductPrice float64, ProductNumber int, ProductListName string) error {
	// 更新 AccountList 表中的余额
	_, err := db.Exec("UPDATE AccountList SET Balance = Balance + ? WHERE UserAccount = ?", ProductPrice*float64(ProductNumber), ProductListName)
	if err != nil {
		return fmt.Errorf("无法更新 AccountList 中商家的余额: %v", err)
	}
	return nil
}
func AddHostList(db *sql.DB, ProductPrice float64, ProductNumber int, ProductName string) error {
	_, err := db.Exec("INSERT INTO HostList (ProductName, SoldNumber, Price) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE SoldNumber = SoldNumber + ?, Price = ?", ProductName, ProductNumber, ProductPrice, ProductNumber, ProductPrice)
	if err != nil {
		return fmt.Errorf("无法更新或插入到 HotList 中: %v", err)
	}
	return nil
}

func AddBuyUserAccountList(db *sql.DB, BuyUserAccount string, ProductPrice float64, ProductNumber int, ProductName string) error {
	// 在处理购物车商品后，向 BuyUserAccount 表中插入数据
	_, err := db.Exec("INSERT INTO "+BuyUserAccount+" (ProductName, ProductNumber, ProductPrice) VALUES (?, ?, ?)", ProductName, ProductNumber, ProductPrice)
	if err != nil {
		return fmt.Errorf("无法插入到 %s 中: %v", BuyUserAccount, err)
	}
	return nil
}
func BuyEmptyCar(db *sql.DB, BuyUserAccount string) error {
	carList, err := CountCarList(db)
	if err != nil {
		return fmt.Errorf("无法从购物车中获取数据: %v", err)
	}

	for _, car := range carList {
		//Id:=car.Id
		//AllProductID:=car.AllProductID
		ProductListName := car.ProductListName
		//ProductID:=car.ProductID
		ProductName := car.ProductName
		ProductNumber := car.ProductNumber
		ProductPrice := car.ProductPrice
		err := AddProductListNameBalance(db, ProductPrice, ProductNumber, ProductListName)
		if err != nil {
			return err
		}
		err = AddHostList(db, ProductPrice, ProductNumber, ProductName)
		if err != nil {
			return err
		}
		err = AddBuyUserAccountList(db, BuyUserAccount, ProductPrice, ProductNumber, ProductName)
		if err != nil {
			return err
		}
	}

	DropList(db, "Car")
	global.Logger.Info("成功处理购物车中的商品，并清空了购物车") // 添加日志记录

	return nil
}
