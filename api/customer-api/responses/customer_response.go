package responses

type CustomerResponse struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	LicenseId   string `json:"license_id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Active      bool   `json:"active"`
}
