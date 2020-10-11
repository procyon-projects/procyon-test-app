package errors

import (
	web "github.com/procyon-projects/procyon-web"
	"net/http"
)

type ErrorHandlerAdviser struct {
}

func NewErrorHandlerAdviser() ErrorHandlerAdviser {
	return ErrorHandlerAdviser{}
}

func (adviser ErrorHandlerAdviser) RegisterErrorHandlers(registry web.ErrorHandlerRegistry) {
	registry.Register(
		web.NewErrorHandler(adviser.HandleProductNotFoundException, ProductNotFoundError{}),
	)
}

func (adviser ErrorHandlerAdviser) HandleProductNotFoundException(notFoundError ProductNotFoundError) *web.ResponseEntity {
	return web.NewResponseEntity(
		web.WithStatus(http.StatusNotFound),
	)
}
