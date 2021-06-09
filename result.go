package shop

import "fmt"

type Result struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (r *Result) OK() bool {
	return r.ErrCode == 0
}

func (r *Result) Error() string {
	return fmt.Sprintf("%d:%s", r.ErrCode, r.ErrMsg)
}
