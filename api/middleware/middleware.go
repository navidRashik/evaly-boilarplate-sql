package middleware

import (
	"go-mysql-boilerplate/logger"
	"net/http"
)

// Middleware represents http handler middleware
type Middleware func(http.Handler) http.Handler

var lgr logger.Logger

func SetLogger(l logger.Logger) {
	lgr = l
}
