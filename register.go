package shop

type Register struct {
	accessToken string
}

func NewRegister(accessToken string) *Register {
	return &Register{accessToken: accessToken}
}

/*
Apply
通过此接口开通自定义版交易组件，将同步返回接入结果，不再有异步事件回调。
如果账户已接入标准版组件，则无法开通，请到微信公众平台取消标准组件的开通。
*/
func (r *Register) Apply() (bool, error) {
	var result = new(Result)
	_, err := POST(r.accessToken, "shop/register/apply", nil, result)
	if err != nil {
		return false, err
	}
	return result.OK(), nil
}

type RegisterCheckResult struct {
	Result
	Data *RegisterCheckData `json:"data"`
}

type RegisterCheckData struct {
	Status     int `json:"status"`
	AccessInfo struct {
		SpuAuditSuccess int `json:"spu_audit_success"`
		PayOrderSuccess int `json:"pay_order_success"`
	} `json:"access_info"`
}

/*
Check
如果账户未接入，将返回错误码1040003。
*/
func (r *Register) Check() (*RegisterCheckResult, error) {
	var result = &RegisterCheckResult{
		Data: &RegisterCheckData{
			// AccessInfo: &AccessInfo{},
		},
	}
	_, err := POST(r.accessToken, "shop/register/check", nil, result)
	if err != nil {
		return nil, err
	}
	if !result.OK() {
		return nil, result
	}
	return result, nil
}
