package shop

type DeliveryItem struct {
	DeliveryID   string `json:"delivery_id"`
	DeliveryName string `json:"delivery_name"`
}

type DeliveryCompanyListResult struct {
	Result
	CompanyList []*DeliveryItem `json:"company_list"`
}

type Delivery struct {
	accessToken string
}

func NewDelivery(accessToken string) *Delivery {
	return &Delivery{accessToken: accessToken}
}

/*
GetCompanyList
获取快递公司列表
*/
func (d *Delivery) GetCompanyList() (*DeliveryCompanyListResult, error) {
	result := new(DeliveryCompanyListResult)
	_, err := POST(d.accessToken, "shop/delivery/get_company_list", nil, result)
	if err != nil {
		return nil, err
	}
	if !result.OK() {
		return nil, result
	}
	return result, nil
}

type DeliverySend struct {
	OutOrderID        string `json:"out_order_id"`
	Openid            string `json:"openid"`
	FinishAllDelivery int    `json:"finish_all_delivery"`
	DeliveryList      []struct {
		DeliveryID string `json:"delivery_id"`
		WaybillID  string `json:"waybill_id"`
	} `json:"delivery_list"`
}

/*
Send
订单发货
*/
func (d *Delivery) Send(send *DeliverySend) error {
	result := new(Result)
	_, err := POST(d.accessToken, "shop/delivery/send", send, result)
	if err != nil {
		return err
	}
	if !result.OK() {
		return result
	}
	return nil
}

/*
Recieve
订单确认收货
把订单状态从30（待收货）流转到100（完成）
*/
func (d *Delivery) Recieve(order *ClientOrder) error {
	result := new(Result)
	_, err := POST(d.accessToken, "shop/delivery/recieve", order, result)
	if err != nil {
		return err
	}
	if !result.OK() {
		return result
	}
	return nil
}
