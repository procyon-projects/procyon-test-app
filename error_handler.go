package main

import (
	context "github.com/procyon-projects/procyon-context"
	web "github.com/procyon-projects/procyon-web"
)

type ErrorHandler struct {
	logger context.Logger
}

func NewErrorHandler(logger context.Logger) ErrorHandler {
	return ErrorHandler{
		logger,
	}
}

func (handler ErrorHandler) HandleError(err error, requestContext *web.WebRequestContext) {
	// do whatever you want
}
