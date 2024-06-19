package controllers

import "tokoku/internal/models"

type AdminController struct {
	model *models.AdminModel
}

func NewAdminController(m *models.AdminModel) *AdminController {
	return &AdminController{
		model: m,
	}
}
