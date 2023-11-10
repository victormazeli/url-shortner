package routes

import (
	"github.com/gin-gonic/gin"
	"backend/api/handlers"
	"backend/internal/config"
	"gorm.io/gorm"
)

func UrlRoute(cfg *config.Config, db *gorm.DB, r *gin.RouterGroup) {
	urlHandler := handlers.UrlHandler{
		Cfg: cfg,
		Db:  db,
	}

	r.GET("/:id", urlHandler.GetUrl)
	r. POST("/shorten", urlHandler.CreateUrlAliase)
}
