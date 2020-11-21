package request

type ProductCreateRequest struct {
	Body struct {
		Name     string `json:"productName" yaml:"productName"`
		Category int    `json:"categoryId" yaml:"parentCategoryId"`
	} `request:"body"`
}

type Category struct {
	Name  string `json:"name" yaml:"name"`
	Index int    `json:"index" yaml:"index"`
}

type ProductUpdateRequest struct {
	Body struct {
		Name       string     `json:"productName" yaml:"productName"`
		Category   int        `json:"categoryId" yaml:"categoryId"`
		Categories []Category `json:"categories" yaml:"categories"`
	} `request:"body"`
	PathVariables struct {
		ProductId int `json:"productId" yaml:"productId"`
	} `request:"path"`
	RequestParams struct {
		Order string `json:"order" yaml:"order"`
	} `request:"param"`
	Header struct {
		ContentType string `json:"Content-Type" yaml:"Content-Type"`
	} `request:"header"`
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
