package requests

type CreateCustomerRequest struct {
	Name        string `json:"name" binding:"required"`
	Password    string `json:"password" binding:"required,min=8"`
	Address     string `json:"address" binding:"required"`
	LicenseId   string `json:"license_id" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Active      bool   `json:"active"`
}

type UpdateCustomerRequest struct {
	Id          int64  `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Password    string `json:"password" binding:"required,min=8"`
	Address     string `json:"address" binding:"required"`
	LicenseId   string `json:"license_id" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Active      bool   `json:"active"`
}

type ChangePasswordRequest struct {
	Id              int64  `json:"id" binding:"required"`
	OldPassword     string `json:"old_pass" binding:"required"`
	NewPassword     string `json:"new_pass" binding:"required"`
	ConfirmPassword string `json:"confirm_pass" binding:"required"`
}
