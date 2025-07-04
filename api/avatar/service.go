package avatar

import (
	"github.com/smartbot/account/database"
	"github.com/smartbot/account/pkg/dbclient"
	"github.com/smartbot/account/pkg/errors"
)

type AvatarService struct {
}

func (ac *AvatarService) SaveAvatar(userId string, request SaveAvatarRequest) *errors.ApiError {
	db := dbclient.GetCient()

	result := db.Model(&database.User{}).Where("id = ?", userId).Update("avatar", request.Avatar)

	if result.Error != nil {
		return errors.InternalServerError("Failed to update avatar")
	}

	return nil

}
