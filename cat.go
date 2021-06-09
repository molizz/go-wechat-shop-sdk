package shop

type Cat struct {
	accessToken string
}

func NewCat(accessToken string) *Cat {
	return &Cat{accessToken: accessToken}
}

type GetCatResult struct {
	Result
	ThirdCatList []*ThirdCat `json:"third_cat_list"`
}

type ThirdCat struct {
	ThirdCatID               int    `json:"third_cat_id"`
	ThirdCatName             string `json:"third_cat_name"`
	Qualification            string `json:"qualification"`
	QualificationType        int    `json:"qualification_type"`
	ProductQualification     string `json:"product_qualification"`
	ProductQualificationType int    `json:"product_qualification_type"`
	FirstCatID               int    `json:"first_cat_id"`
	FirstCatName             string `json:"first_cat_name"`
	SecondCatID              int    `json:"second_cat_id"`
	SecondCatName            string `json:"second_cat_name"`
}

/*
Get
获取所有三级类目及其资质相关信息 注意：该接口拉到的是【全量】三级类目数据，数据回包大小约为2MB。 所以请商家自己做好缓存，不要频繁调用（有严格的频率限制），该类目数据不会频率变动，推荐商家每天调用一次更新商家自身缓存

若该类目资质必填，则新增商品前，必须先通过该类目资质申请接口进行资质申请; 若该类目资质不需要，则该类目自动拥有，无需申请，如依然调用，会报错1050011； 若该商品资质必填，则新增商品时，带上商品资质字段。 接入类目审核回调，才可获取审核结果。
*/
func (c *Cat) Get() (*GetCatResult, error) {
	var result = &GetCatResult{
		ThirdCatList: make([]*ThirdCat, 0),
	}
	_, err := POST(c.accessToken, "shop/cat/get", nil, result)
	if err != nil {
		return nil, err
	}
	if !result.OK() {
		return nil, result
	}
	return result, nil
}
