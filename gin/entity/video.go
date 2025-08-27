package entity

type Video struct {
	Title       string `json:"name"`
	Description string `json:"overview"`
	URL         string `json:"link"`
}
