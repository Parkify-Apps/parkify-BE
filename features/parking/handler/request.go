package handler

type ParkingRequest struct {
	ImageLoc     string `form:"imageloc"`
	LocationName string	`form:"locname"`
	Location     string	`form:"location"`
}