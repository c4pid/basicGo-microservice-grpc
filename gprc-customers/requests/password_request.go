package requests

type ChangePasswordRequest struct {
	Id              int64
	OldPassword     string
	NewPassword     string
	ConfirmPassword string
}
