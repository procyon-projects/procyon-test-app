package resource

type EmployeeResource struct {
	Id   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	Role string `json:"role" xml:"role"`
}
