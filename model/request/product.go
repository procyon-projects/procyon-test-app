package request

type ProductCreateRequest struct {
	Body struct {
		Name     string `json:"productName" yaml:"productName"`
		Category int    `json:"categoryId" yaml:"categoryId"`
	} `request:"body"`
}

type ProductUpdateRequest struct {
	Body struct {
		Name     string `json:"productName" yaml:"productName"`
		Category int    `json:"categoryId" yaml:"categoryId"`
	} `request:"body"`
	PathVariables struct {
		ProductId int `json:"productId" yaml:"productId"`
	} `request:"path"`
}

type ProductGetRequest struct {
	PathVariables struct {
		ProductId int `json:"productId" yaml:"productId"`
	} `request:"path"`
}

type ProductDeleteRequest struct {
	PathVariables struct {
		ProductId int `json:"productId" yaml:"productId"`
	} `request:"path"`
}
