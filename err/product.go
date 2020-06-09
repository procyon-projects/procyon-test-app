package err

import "fmt"

type ProductNotFoundError struct {
	message string
}

func NewProductNotFoundError(productId int) ProductNotFoundError {
	return ProductNotFoundError{
		message: fmt.Sprintln("Product not found with id : ", productId),
	}
}

func (err ProductNotFoundError) Error() string {
	return err.message
}
