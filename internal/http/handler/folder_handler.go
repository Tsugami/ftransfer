package handler

import (
	"net/http"

	"github.com/Tsugami/ftransfer/internal/domain/repository"
	"github.com/Tsugami/ftransfer/internal/http/dto/request"
	"github.com/Tsugami/ftransfer/internal/http/dto/response"
	"github.com/gin-gonic/gin"
)

type FolderHandler struct {
	folderRepo repository.FolderRepository
}

func NewFolderHandler(folderRepo repository.FolderRepository) *FolderHandler {
	return &FolderHandler{
		folderRepo: folderRepo,
	}
}

func (h *FolderHandler) Create(c *gin.Context) {
	var createReq request.CreateFolderRequest
	if err := c.ShouldBindJSON(&createReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	folder, err := createReq.ToDomain()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.folderRepo.Create(c.Request.Context(), folder); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, response.NewFolderResponse(folder))
}

func (h *FolderHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Folder ID is required"})
		return
	}

	folder, err := h.folderRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response.NewFolderResponse(folder))
}

func (h *FolderHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Folder ID is required"})
		return
	}

	if err := h.folderRepo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}
