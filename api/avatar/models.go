package avatar

type SaveAvatarRequest struct {
	Avatar string `json:"avatar" validate:"required"`
}
