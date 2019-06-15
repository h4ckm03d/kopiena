package domain

import "time"

type Coordinate struct {
	Longitude float64
	Latitude float64
}

type Location struct {
	Name       string     `json:"name,omitempty"`
	Address    string     `json:"address,omitempty"`
	Coordinate Coordinate `json:"coordinate,omitempty"`
	Country    string     `json:"country,omitempty"`
}

type Event struct {
	ID          uint      `json:"id,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	Description string    `json:"description,omitempty"`
	Location    Location  `json:"location,omitempty"`
	CommunityID uint      `json:"community_id,omitempty"`
}

