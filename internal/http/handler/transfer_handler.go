package handler

import (
	"net/http"

	"github.com/Tsugami/ftransfer/internal/domain/repository"
	"github.com/Tsugami/ftransfer/internal/http/dto/request"
	"github.com/Tsugami/ftransfer/internal/http/dto/response"
	"github.com/gin-gonic/gin"
)

type TransferHandler struct {
	transferRepo repository.TransferRepository
}

func NewTransferHandler(transferRepo repository.TransferRepository) *TransferHandler {
	return &TransferHandler{
		transferRepo: transferRepo,
	}
}

func (h *TransferHandler) Create(c *gin.Context) {
	var createReq request.CreateTransferRequest
	if err := c.ShouldBindJSON(&createReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transfer, err := createReq.ToDomain()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.transferRepo.Create(c.Request.Context(), transfer); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, response.NewTransferResponse(transfer))
}

func (h *TransferHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Transfer ID is required"})
		return
	}

	transfer, err := h.transferRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response.NewTransferResponse(transfer))
}

func (h *TransferHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Transfer ID is required"})
		return
	}

	if err := h.transferRepo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}
