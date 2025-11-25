package api

type User struct {
	ID           string `json:"id"`
	DisplayName  string `json:"display_name"`
	EmailAddress string `json:"email_address"`
}
