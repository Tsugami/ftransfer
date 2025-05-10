package main

import (
	"net/http"

	"github.com/Tsugami/ftransfer/internal/storage_provider"
	"github.com/Tsugami/ftransfer/internal/transfer"
	"github.com/gin-gonic/gin"
)

type CreateTransferRequest struct {
	SourceStorageProviderID      string `json:"source_storage_provider_id" binding:"required"`
	DestinationStorageProviderID string `json:"destination_storage_provider_id" binding:"required"`
	SourceDir                    string `json:"source_dir" binding:"required"`
	DestinationDir               string `json:"destination_dir" binding:"required"`
	PostTransferSourceDir        string `json:"post_transfer_source_dir" binding:"required"`
}

type UpdateTransferRequest struct {
	SourceStorageProviderID      string `json:"source_storage_provider_id" binding:"required"`
	DestinationStorageProviderID string `json:"destination_storage_provider_id" binding:"required"`
	SourceDir                    string `json:"source_dir" binding:"required"`
	DestinationDir               string `json:"destination_dir" binding:"required"`
	PostTransferSourceDir        string `json:"post_transfer_source_dir" binding:"required"`
}

type TransferResponse struct {
	ID                           string `json:"id"`
	SourceStorageProviderID      string `json:"source_storage_provider_id"`
	DestinationStorageProviderID string `json:"destination_storage_provider_id"`
	SourceDir                    string `json:"source_dir"`
	DestinationDir               string `json:"destination_dir"`
	PostTransferSourceDir        string `json:"post_transfer_source_dir"`
}

func NewTransferResponse(transfer *transfer.Transfer) TransferResponse {
	return TransferResponse{
		ID:                           transfer.ID.String(),
		SourceStorageProviderID:      transfer.SourceStorageProviderID.String(),
		DestinationStorageProviderID: transfer.DestinationStorageProviderID.String(),
		SourceDir:                    transfer.SourceDir.String(),
		DestinationDir:               transfer.DestinationDir.String(),
		PostTransferSourceDir:        transfer.PostTransferSourceDir.String(),
	}
}

func NewTransferResponseList(transfers []*transfer.Transfer) []TransferResponse {
	response := make([]TransferResponse, len(transfers))
	for i, t := range transfers {
		response[i] = NewTransferResponse(t)
	}
	return response
}

func (s *Handler) CreateTransfer(c *gin.Context) {
	var input CreateTransferRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transfer, err := s.transferService.Create(c.Request.Context(),
		transfer.NewDirectory(input.SourceDir),
		transfer.NewDirectory(input.DestinationDir),
		transfer.NewDirectory(input.PostTransferSourceDir),
		storage_provider.ID(input.SourceStorageProviderID),
		storage_provider.ID(input.DestinationStorageProviderID),
	)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, NewTransferResponse(transfer))
}

func (s *Handler) ListTransfers(c *gin.Context) {
	transfers, err := s.transferService.List(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, NewTransferResponseList(transfers))
}

func (s *Handler) GetTransfer(c *gin.Context) {
	id := c.Param("id")
	transfer, err := s.transferService.Get(c.Request.Context(), transfer.ID(id))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, NewTransferResponse(transfer))
}

func (s *Handler) UpdateTransfer(c *gin.Context) {
	id := c.Param("id")
	var input UpdateTransferRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.transferService.Update(c.Request.Context(), transfer.ID(id),
		transfer.NewDirectory(input.SourceDir),
		transfer.NewDirectory(input.DestinationDir),
		transfer.NewDirectory(input.PostTransferSourceDir),
		storage_provider.ID(input.SourceStorageProviderID),
		storage_provider.ID(input.DestinationStorageProviderID),
	)

	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *Handler) DeleteTransfer(c *gin.Context) {
	id := c.Param("id")
	err := s.transferService.Delete(c.Request.Context(), transfer.ID(id))
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
