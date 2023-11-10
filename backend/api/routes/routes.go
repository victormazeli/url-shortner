package routes

import (
	"github.com/gin-gonic/gin"
	"backend/internal/config"
	"gorm.io/gorm"
)

func SetupRoute(cfg *config.Config, db *gorm.DB, rg *gin.RouterGroup) {
	UrlRoute(cfg, db, rg)

}
