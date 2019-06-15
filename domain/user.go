package domain

type User struct {
	ID       uint   `json:"id,omitempty"`
	Contact  string `json:"contact,omitempty"`
	City     string `json:"city,omitempty"`
	Country  string `json:"country,omitempty"`
	Email    string `json:"email,omitempty"`
	Name     string `json:"name,omitempty"`
	PhotoURL string `json:"photo_url,omitempty"`
	Role     string `json:"role,omitempty"`
}
