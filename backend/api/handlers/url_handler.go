package handlers

import (
	"errors"
	"backend/internal/config"
	"backend/internal/database/models"
	Apirequest "backend/pkg/ApiRequest"
	"backend/pkg/ApiResponse"
	"backend/pkg/utils"
	"log"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type UrlHandler struct {
	Db  *gorm.DB
	Cfg *config.Config
}

func (u *UrlHandler) CreateUrlAliase(c *gin.Context) {

	var createAliaseInput Apirequest.CreateUrlAliaseInput

	if err := c.ShouldBind(&createAliaseInput); err != nil {
		ApiResponse.SendBadRequest(c, err.Error())
		return
	}

	var url models.Url

	findAliaseError := u.Db.Where(&models.Url{LongUrl: createAliaseInput.Url}).First(&url)

	if findAliaseError.Error != nil {
		if errors.Is(findAliaseError.Error, gorm.ErrRecordNotFound) {
			shortCode := utils.EncodeToString()

			newAliase := &models.Url{
				LongUrl:  createAliaseInput.Url,
				ShortUrl: "https://localhost:5173/" + shortCode,
				ShortCode: shortCode,
			}

			newAliaseError := u.Db.Create(newAliase)

			if newAliaseError.Error != nil {
				ApiResponse.SendInternalServerError(c, "Internal Server Error")
				return
			}

			ApiResponse.SendCreated(c, "URL Aliase Created Successfully", newAliase)
			return
		} else {
			log.Print(findAliaseError.Error)
			ApiResponse.SendInternalServerError(c, "Internal Server Error")
			return
		}
	}

	ApiResponse.SendSuccess(c, "URL Aliase Fetched Successfully", url)
}

func (u *UrlHandler) GetUrl(c *gin.Context) {
	shortCode := c.Param("code")
	var url models.Url

	findUrlError := u.Db.Where(&models.Url{ShortUrl: shortCode}).First(&url)

	if findUrlError.Error != nil {
		if errors.Is(findUrlError.Error, gorm.ErrRecordNotFound) {
			ApiResponse.SendNotFound(c, "URL not found")
		} else {
			log.Print(findUrlError.Error)
			ApiResponse.SendInternalServerError(c, "Internal Server Error")
			return
		}
	}

	ApiResponse.SendSuccess(c, "URL retrieved successfully", url)

}

// func (u *UserHandler) UpdateUser(c *gin.Context) {
// 	_, _ = strconv.Atoi(c.Param("id"))
// 	// Update user in the database by userID
// 	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
// }

// func (u *UserHandler) DeleteUser(c *gin.Context) {
// 	_, _ = strconv.Atoi(c.Param("id"))
// 	// Delete user from the database by userID
// 	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
// }
