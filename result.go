package shop

type Result struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (r *Result) OK() bool {
	return r.ErrCode == 0
}
