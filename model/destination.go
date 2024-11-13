package model

import "time"

type Destination struct {
	Id           int
	Name         string
	Description  string
	ImageUrl     string
	Date         time.Time
	Price        int
	TotalBooking int
	Rating       float32
}
type QueryParams struct {
	Page        int
	SortDate    bool
	SortPrice   string
	SortName    bool
	SearchPlace string
	SearchDate  string
	SearchPrice int
}
