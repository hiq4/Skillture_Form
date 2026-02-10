package server

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, adminHandler *AdminHandler) {
	// Public routes
	r.GET("/health", HealthCheck)

	// Admin routes group
	admin := r.Group("/admin")
	admin.Use(AdminLoggingMiddleware()) // 🌟 هنا نضيف logging للadmin فقط
	{
		admin.POST("/create", adminHandler.Create)
		admin.GET("/list", adminHandler.List)
		admin.PUT("/update/:id", adminHandler.Update)
		admin.DELETE("/delete/:id", adminHandler.Delete)
	}
}
