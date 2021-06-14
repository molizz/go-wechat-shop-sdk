package shop

type Aftersale struct {
	accessToken string
}

func NewAftersale(accessToken string) *Aftersale {
	return &Aftersale{
		accessToken: accessToken,
	}
}

type AftersaleInfo struct {
	OutOrderID         string `json:"out_order_id"`
	OutAftersaleID     string `json:"out_aftersale_id"`
	Openid             string `json:"openid"`
	Type               int    `json:"type"`
	CreateTime         string `json:"create_time"`
	Status             int    `json:"status"`
	FinishAllAftersale int    `json:"finish_all_aftersale,omitempty"`
	Path               string `json:"path"`
	ProductInfos       []struct {
		OutProductID string `json:"out_product_id"`
		OutSkuID     string `json:"out_sku_id"`
		ProductCnt   int    `json:"product_cnt"`
	} `json:"product_infos"`
}

/*
Add
创建售后

订单原始状态为10, 200, 250时会返回错误码100000
finish_all_aftersale = 1时订单状态会流转到200（全部售后结束，不可继续售后）
*/
func (a *Aftersale) Add(aftersale *AftersaleInfo) error {
	result := new(Result)
	_, err := POST(a.accessToken, "shop/aftersale/add", aftersale, result)
	if err != nil {
		return err
	}
	if !result.OK() {
		return result
	}
	return nil
}

type AftersaleResult struct {
	Result
	Aftersaleinfos []*AftersaleInfo `json:"aftersale_infos"`
}

/*
Get
获取订单下售后单

订单原始状态为10, 200, 250时会返回错误码100000
finish_all_aftersale = 1时订单状态会流转到200（全部售后结束，不可继续售后）
*/
func (a *Aftersale) Get(order *ClientOrder) error {
	result := new(AftersaleResult)
	_, err := POST(a.accessToken, "shop/aftersale/get", order, result)
	if err != nil {
		return err
	}
	if !result.OK() {
		return result
	}
	return nil
}

type AftersaleUpdate struct {
	ClientOrder
	OutAftersaleID     string `json:"out_aftersale_id"`
	Status             int    `json:"status"`               // 0:未受理,1:用户取消,2:商家受理中,3:商家逾期未处理,4:商家拒绝退款,5:商家拒绝退货退款,6:待买家退货,7:退货退款关闭,8:待商家收货,11:商家退款中,12:商家逾期未退款,13:退款完成,14:退货退款完成,15:换货完成,16:待商家发货,17:待用户确认收货,18:商家拒绝换货,19:商家已收到货
	FinishAllAftersale int    `json:"finish_all_aftersale"` // 0:售后未结束, 1:售后结束且订单状态流转
}

func (a *Aftersale) Update(update *AftersaleUpdate) error {
	result := new(Result)
	_, err := POST(a.accessToken, "shop/aftersale/get", update, result)
	if err != nil {
		return err
	}
	if !result.OK() {
		return result
	}
	return nil
}
