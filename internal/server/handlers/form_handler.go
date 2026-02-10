package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"Skillture_Form/internal/domain/entities"
	"Skillture_Form/internal/domain/enums"
	uc "Skillture_Form/internal/usecase/interfaces"
)

type FormHandler struct {
	formUC uc.FormUseCase
}

func NewFormHandler(formUC uc.FormUseCase) *FormHandler {
	return &FormHandler{formUC: formUC}
}

func (h *FormHandler) Create(c *gin.Context) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	form := &entities.Form{
		ID:          uuid.New(),
		Title:       req.Title,
		Description: req.Description,
		Status:      enums.FormStatusDraft,
	}

	if err := h.formUC.Create(c.Request.Context(), form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, form)
}
