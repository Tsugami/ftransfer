package main

import (
	"net/http"

	"github.com/Tsugami/ftransfer/internal/protocol"
	error_middleware "github.com/Tsugami/ftransfer/pkg/errormiddleware"
	"github.com/gin-gonic/gin"
)

func SetupMiddleware() gin.HandlerFunc {
	return error_middleware.ErrorHandler(
		error_middleware.Map(protocol.ErrInvalidProtocolConnection).ToStatusCode(http.StatusBadRequest),
		error_middleware.Map(protocol.ErrInvalidProtocol).ToStatusCode(http.StatusBadRequest),
		error_middleware.Map(protocol.ErrEmptyFTPHost).ToStatusCode(http.StatusBadRequest),
		error_middleware.Map(protocol.ErrEmptyFTPPort).ToStatusCode(http.StatusBadRequest),
		error_middleware.Map(protocol.ErrEmptyFTPUsername).ToStatusCode(http.StatusBadRequest),
		error_middleware.Map(protocol.ErrEmptyFTPPassword).ToStatusCode(http.StatusBadRequest),
		error_middleware.Map(protocol.ErrEmptySFTPUsername).ToStatusCode(http.StatusBadRequest),
		error_middleware.Map(protocol.ErrEmptySFTPCredentials).ToStatusCode(http.StatusBadRequest),
		error_middleware.Map(protocol.ErrEmptyS3Bucket).ToStatusCode(http.StatusBadRequest),
		error_middleware.Map(protocol.ErrEmptyS3AccessKeyID).ToStatusCode(http.StatusBadRequest),
		error_middleware.Map(protocol.ErrEmptyS3SecretAccessKey).ToStatusCode(http.StatusBadRequest),
		error_middleware.Map(protocol.ErrEmptyS3Region).ToStatusCode(http.StatusBadRequest),
		error_middleware.Map(protocol.ErrEmptyS3Endpoint).ToStatusCode(http.StatusBadRequest),
		error_middleware.Map(protocol.ErrEmptyS3UseSSL).ToStatusCode(http.StatusBadRequest),
	)
}
