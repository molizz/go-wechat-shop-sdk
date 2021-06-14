package shop

type Order struct {
	accessToken string
}

func NewOrder(accessToken string) *Order {
	return &Order{
		accessToken: accessToken,
	}
}

type ProductInfo struct {
	OutProductID string `json:"out_product_id"`
	OutSkuID     string `json:"out_sku_id"`
	ProductCnt   int    `json:"product_cnt"`
	SalePrice    int    `json:"sale_price"` // 生成这次订单时商品的售卖价，可以跟上传商品接口的价格不一致
	Path         string `json:"path"`
	Title        string `json:"title"`
	HeadImg      string `json:"head_img"`
}

type PayInfo struct {
	PayMethodType int    `json:"pay_method_type"` // 0: 微信支付, 1: 货到付款, 2: 商家会员储蓄卡（默认0）
	PrepayID      string `json:"prepay_id"`       // pay_method_type = 0时必填
	PrepayTime    string `json:"prepay_time"`
}

// 注意价格字段的单价是分，不是元
type PriceInfo struct {
	OrderPrice        int    `json:"order_price"`
	Freight           int    `json:"freight"`
	DiscountedPrice   int    `json:"discounted_price"`
	AdditionalPrice   int    `json:"additional_price"`
	AdditionalRemarks string `json:"additional_remarks"`
}

type OrderDetail struct {
	ProductInfos []*ProductInfo `json:"product_infos"`
	PayInfo      *PayInfo       `json:"pay_info"`
	PriceInfo    *PriceInfo     `json:"price_info"`
}

type AddressInfo struct {
	ReceiverName    string `json:"receiver_name"`
	DetailedAddress string `json:"detailed_address"`
	TelNumber       string `json:"tel_number"`
	Country         string `json:"country"`
	Province        string `json:"province"`
	City            string `json:"city"`
	Town            string `json:"town"`
}

type DeliveryDetail struct {
	DeliveryType int `json:"delivery_type"` // 1: 正常快递, 2: 无需快递, 3: 线下配送, 4: 用户自提
}

type AddOrder struct {
	CreateTime     string          `json:"create_time"`
	Type           int             `json:"type"`         // 非必填，默认为0。0:普通场景, 1:合单支付
	OutOrderID     string          `json:"out_order_id"` // 必填，普通场景下的外部订单ID；合单支付（多订单合并支付一次）场景下是主外部订单ID
	Openid         string          `json:"openid"`       //
	Path           string          `json:"path"`         // 这里的path中的最好有一个参数的值能和out_order_id的值匹配上
	Scene          int             `json:"scene"`        // 下单时小程序的场景值，可通过[getLaunchOptionsSync](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/life-cycle/wx.getLaunchOptionsSync.html)或[onLaunch/onShow](https://developers.weixin.qq.com/miniprogram/dev/reference/api/App.html#onLaunch-Object-object)拿到
	OutUserID      string          `json:"out_user_id"`
	OrderDetail    *OrderDetail    `json:"order_detail"`
	DeliveryDetail *DeliveryDetail `json:"delivery_detail"`
	AddressInfo    *AddressInfo    `json:"address_info"`
}

type AddOrderResult struct {
	Result
	Data struct {
		OrderID          int    `json:"order_id"`
		OutOrderID       string `json:"out_order_id"`
		Ticket           string `json:"ticket"`
		TicketExpireTime string `json:"ticket_expire_time"`
		FinalPrice       int    `json:"final_price"`
	} `json:"data"`
}

// Add
// 生成订单
func (o *Order) Add(order *AddOrder) error {
	result := new(AddOrderResult)

	_, err := POST(o.accessToken, "shop/order/add", order, result)
	if err != nil {
		return err
	}
	if !result.OK() {
		return result
	}
	return nil
}

type PayOrder struct {
	OrderID       int    `json:"order_id"`
	OutOrderID    string `json:"out_order_id"`
	Openid        string `json:"openid"`
	ActionType    int    `json:"action_type"`
	ActionRemark  string `json:"action_remark"`
	TransactionID string `json:"transaction_id"`
	PayTime       string `json:"pay_time"`
}

// Pay
// 同步订单支付结果
// https://developers.weixin.qq.com/miniprogram/dev/framework/ministore/minishopopencomponent2/API/order/pay_order.html
func (o *Order) Pay(pay *PayOrder) error {
	result := new(Result)
	_, err := POST(o.accessToken, "shop/order/pay", pay, result)
	if err != nil {
		return err
	}
	if !result.OK() {
		return result
	}
	return nil
}

type GetOrderDetail struct {
	OrderDetail
	MultiPayInfo   []*PayInfo `json:"multi_pay_info"`
	DeliveryDetail struct {
		DeliveryType      int `json:"delivery_type"`
		FinishAllDelivery int `json:"finish_all_delivery"`
		DeliveryList      []struct {
			WaybillID  string `json:"waybill_id"`
			DeliveryID string `json:"delivery_id"`
		} `json:"delivery_list"`
	} `json:"delivery_detail"`
}

type GetOrderResult struct {
	Result
	Order struct {
		OrderID     int             `json:"order_id"`
		OutOrderID  string          `json:"out_order_id"`
		Status      int             `json:"status"`
		Path        string          `json:"path"`
		OrderDetail *GetOrderDetail `json:"order_detail"`
	} `json:"order"`
}

type ClientOrder struct {
	OutOrderID string `json:"out_order_id"`
	OpenID     string `json:"openid"`
}

// Get
// 获取订单详情
// https://developers.weixin.qq.com/miniprogram/dev/framework/ministore/minishopopencomponent2/API/order/get_order.html
func (o *Order) Get(get *ClientOrder) (*GetOrderResult, error) {
	result := new(GetOrderResult)
	_, err := POST(o.accessToken, "shop/order/get", get, result)
	if err != nil {
		return nil, err
	}
	if !result.OK() {
		return nil, result
	}
	return result, nil
}
