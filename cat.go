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
