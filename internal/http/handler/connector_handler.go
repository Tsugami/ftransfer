package handler

import (
	"fmt"
	"net/http"

	"github.com/Tsugami/ftransfer/internal/domain/repository"
	"github.com/Tsugami/ftransfer/internal/http/dto/request"
	"github.com/Tsugami/ftransfer/internal/http/dto/response"
	"github.com/gin-gonic/gin"
)

type ConnectorHandler struct {
	connectorRepo repository.ConnectorRepository
}

func NewConnectorHandler(connectorRepo repository.ConnectorRepository) *ConnectorHandler {
	return &ConnectorHandler{
		connectorRepo: connectorRepo,
	}
}

func (h *ConnectorHandler) Create(c *gin.Context) {
	var createReq request.CreateConnectorRequest
	if err := c.ShouldBindJSON(&createReq); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	connector, err := createReq.ToDomain()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.connectorRepo.Create(c.Request.Context(), connector); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, response.NewConnectorResponse(connector))
}

func (h *ConnectorHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Connector ID is required"})
		return
	}

	connector, err := h.connectorRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response.NewConnectorResponse(connector))
}

func (h *ConnectorHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Connector ID is required"})
		return
	}

	if err := h.connectorRepo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func (h *ConnectorHandler) List(c *gin.Context) {
	connectors, err := h.connectorRepo.List(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	connectorResponses := make([]*response.ConnectorResponse, len(connectors))
	for i, connector := range connectors {
		connectorResponses[i] = response.NewConnectorResponse(connector)
	}

	c.JSON(200, connectorResponses)
}
