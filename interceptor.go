package main

import (
	context "github.com/procyon-projects/procyon-context"
	web "github.com/procyon-projects/procyon-web"
)

type CustomInterceptor struct {
	logger context.Logger
}

func NewCustomInterceptor(logger context.Logger) CustomInterceptor {
	return CustomInterceptor{
		logger,
	}
}

func (interceptor CustomInterceptor) HandleBefore(requestContext *web.WebRequestContext) {
	interceptor.logger.Info(requestContext, "HandleBefore is invoked")
}

func (interceptor CustomInterceptor) HandleAfter(requestContext *web.WebRequestContext) {
	interceptor.logger.Info(requestContext, "HandleAfter is invoked")
}

func (interceptor CustomInterceptor) AfterCompletion(requestContext *web.WebRequestContext) {
	interceptor.logger.Info(requestContext, "AfterCompletion is invoked")
}
