package models

type Image struct {
	Id          string `json:"id"`
	Description string `json:"alt_description"`
	Urls        struct {
		Regular string `json:"regular"`
		Small   string `json:"small"`
	} `json:"urls"`
	User struct {
		Username string `json:"username"`
	} `json:"user"`
}

type EncodedImages struct {
	Images []EncodedImage `json:"images"`
}

type EncodedImage struct {
	Name        string `json:"name"`
	Encoded     string `json:"raw_data"`
	Description string `json:"description"`
}

type ItemResult struct {
	ObjectName string `json:"name"`
	ObjectId   string `json:"id"`
}
