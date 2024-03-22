package models

type Song struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artists string `json:"artist"`
	Album string `json:"album"`
	Cover string `json:"cover"`
	Duration int `json:"duration"`	
}