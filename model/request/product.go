package request

type ProductRequest struct {
	Body struct {
		Name     string `json:"productName" yaml:"productName"`
		Category int    `json:"categoryId" yaml:"parentCategoryId"`
	} `request:"body"`
}

type ProductCreateRequest struct {
	ProductRequest
}

type ProductUpdateRequest struct {
	ProductRequest
	PathVariables struct {
		ProductId int `json:"id" yaml:"id"`
	} `request:"path"`
	RequestParams struct {
	} `request:"param"`
}

type ProductGetRequest struct {
	PathVariables struct {
		ProductId int `json:"id" yaml:"id"`
	} `request:"path"`
	RequestParams struct {
	} `request:"param"`
}

type ProductDeleteRequest struct {
	PathVariables struct {
		ProductId int `json:"id" yaml:"id"`
	} `request:"path"`
	RequestParams struct {
	} `request:"param"`
}
