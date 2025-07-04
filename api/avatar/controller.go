package avatar

import (
	"github.com/gin-gonic/gin"
	"github.com/smartbot/account/pkg/validator"
)

type AvatarController struct {
	service AvatarService
}

func (ac *AvatarController) SaveAvatar(c *gin.Context) {
	var request SaveAvatarRequest
	userId := c.Param("id")

	err := validator.ValidateUUID(userId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	err = validator.ValidateBody(c, &request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	err = ac.service.SaveAvatar(userId, request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.Status(200)

}
