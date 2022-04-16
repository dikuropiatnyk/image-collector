package models

// TODO: Prepare a format to download image
// TODO: Prepare a JSON-like object to send data to the Image Processor

type Photo struct {
	Id   string `json:"id"`
	Urls struct {
		Regular string `json:"regular"`
		Small   string `json:"small"`
	} `json:"urls"`
	User struct {
		Username string `json:"username"`
	} `json:"user"`
}
