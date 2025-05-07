package main

import (
	"net/http"

	"github.com/Tsugami/ftransfer/internal/protocol"
	"github.com/Tsugami/ftransfer/internal/storage_provider"
	error_middleware "github.com/Tsugami/ftransfer/pkg/errormiddleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupMiddleware(
	app *gin.Engine,
) {
	app.Use(CORSMiddleware())
	app.Use(ErrorMiddleware())
}

func ErrorMiddleware() gin.HandlerFunc {
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
		error_middleware.Map(storage_provider.ErrInvalidStorageProvider).ToStatusCode(http.StatusBadGateway),
		error_middleware.Map(storage_provider.ErrCreateStorageProvider).ToStatusCode(http.StatusBadGateway),
		error_middleware.Map(storage_provider.ErrUpdateStorageProvider).ToStatusCode(http.StatusBadGateway),
		error_middleware.Map(storage_provider.ErrDeleteStorageProvider).ToStatusCode(http.StatusBadGateway),
		error_middleware.Map(storage_provider.ErrGetStorageProvider).ToStatusCode(http.StatusBadGateway),
		error_middleware.Map(storage_provider.ErrListStorageProviders).ToStatusCode(http.StatusBadGateway),
		error_middleware.Map(storage_provider.ErrStorageProviderNotFound).ToStatusCode(http.StatusConflict),
		error_middleware.Map(storage_provider.ErrStorageProviderExists).ToStatusCode(http.StatusNotFound),
		error_middleware.Map(storage_provider.ErrEmptyName).ToStatusCode(http.StatusBadRequest),
		error_middleware.Map(storage_provider.ErrEmptyProtocolConnection).ToStatusCode(http.StatusBadRequest),
		error_middleware.Map(storage_provider.ErrEmptyFileSystem).ToStatusCode(http.StatusBadRequest),
		error_middleware.Map(storage_provider.ErrInvalidFileSystem).ToStatusCode(http.StatusBadRequest),
	)
}

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
	})
}
