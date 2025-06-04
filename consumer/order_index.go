package consumer

import "time"

type OrderIndex struct {
	OrderId                string      `json:"orderId"`
	ExtendOrderId          string      `json:"extend_order_id"`
	Uid                    int64       `json:"uid"`
	RealName               string      `json:"realName"`
	UserPhone              string      `json:"userPhone"`
	UserAddress            string      `json:"userAddress"`
	CartId                 string      `json:"cartId"`
	FreightPrice           float64     `json:"freightPrice"`
	TotalNum               int         `json:"totalNum"`
	TotalPrice             float64     `json:"totalPrice"`
	TotalPostage           float64     `json:"totalPostage"`
	PayPrice               float64     `json:"payPrice"`
	PayPostage             float64     `json:"payPostage"`
	DeductionPrice         float64     `json:"deductionPrice"`
	CouponId               int64       `json:"couponId"`
	CouponPrice            float64     `json:"couponPrice"`
	Paid                   int         `json:"paid"`
	PayTime                time.Time   `json:"payTime"`
	PayType                string      `json:"payType"`
	Status                 int         `json:"status"`
	RefundStatus           int         `json:"refundStatus"`
	RefundReasonWapImg     string      `json:"refundReasonWapImg"`
	RefundReasonWapExplain string      `json:"refundReasonWapExplain"`
	RefundReasonTime       time.Time   `json:"refundReasonTime"`
	RefundReasonWap        string      `json:"refundReasonWap"`
	RefundReason           string      `json:"refundReason"`
	RefundPrice            float64     `json:"refundPrice"`
	DeliverySn             string      `json:"deliverySn"`
	DeliveryName           string      `json:"deliveryName"`
	DeliveryType           string      `json:"deliveryType"`
	DeliveryId             string      `json:"deliveryId"`
	GainIntegral           int         `json:"gainIntegral"`
	UseIntegral            int         `json:"useIntegral"`
	PayIntegral            int         `json:"payIntegral"`
	BackIntegral           int         `json:"backIntegral"`
	Mark                   string      `json:"mark"`
	Unique                 string      `json:"unique"`
	Remark                 string      `json:"remark"`
	CombinationId          int64       `json:"combinationId"`
	PinkId                 int64       `json:"pinkId"`
	Cost                   float64     `json:"cost"`
	SeckillId              int64       `json:"seckillId"`
	BargainId              int64       `json:"bargainId"`
	VerifyCode             string      `json:"verifyCode"`
	StoreId                int64       `json:"storeId"`
	ShippingType           int         `json:"shippingType"`
	CartInfo               interface{} `json:"cartInfo"`
	OrderStatus            int         `json:"_status"`
	OrderStatusName        string      `json:"statusName"`
	PayTypeName            string      `json:"payTypeName"`
}
