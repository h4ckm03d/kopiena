package domain

type Community struct {
	ID          uint    `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	BannerURL   string  `json:"banner_url,omitempty"`
	CreatedBy   User    `json:"created_by,omitempty"`
	Description string  `json:"description,omitempty"`
	Events      []Event `json:"events,omitempty"`
	Members     []User  `json:"members,omitempty"`
}
