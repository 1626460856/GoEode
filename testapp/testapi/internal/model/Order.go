package model

type Order struct {
	OrderID     int     `json:"orderId"`     // 订单id  mysql自增键
	ProductID   int     `json:"productId"`   // 商品id
	ProductName string  `json:"productName"` // 商品名称
	Price       float64 `json:"price"`       // 商品价格
	Boss        string  `json:"boss"`        // 商家
	BuyQuantity int     `json:"buyQuantity"` // 购买商品数量
	UserName    string  `json:"userName"`    // 购买者这个通过传入的token解析获得
	Coupon      float64 `json:"coupon"`      // 优惠券,表格初始值默认为1
	OrderStatus string  `json:"orderStatus"` // 订单状态 有三种状态，“unpaid”为未支付，“paying”为支付中，“paid”为已支付，创建的时候默认未支付
	CreatedAt   string  `json:"createdAt"`   // 创建时间
	UpdatedAt   string  `json:"updatedAt"`   // 更新时间
}