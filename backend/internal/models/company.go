package models

type Companies struct {
	CompanyID       int    `json:"company_id"`
	CompanyName     string `json:"company_name"`
	CompanyAddress  string `json:"company_address"`
	CompanyPhone    string `json:"company_phone"`
	CompanyEmail    string `json:"company_email"`
	CompanyWebsite  string `json:"company_website"`
	CompanyPassword string `json:"company_password"`
}
