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
	Params struct {
		Query int `json:"query" yaml:"query"`
	} `request:"param"`
	PathVariables struct {
		ProductId int `json:"productId" yaml:"productId"`
	} `request:"path"`
	Header struct {
		ContentType string `json:"Content-Type" yaml:"Content-Type"`
	} `request:"header"`
}

type ProductDeleteRequest struct {
	PathVariables struct {
		ProductId int `json:"productId" yaml:"productId"`
	} `request:"path"`
}
