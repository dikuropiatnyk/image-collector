package models

type Image struct {
	Id   string `json:"id"`
	Urls struct {
		Regular string `json:"regular"`
		Small   string `json:"small"`
	} `json:"urls"`
	User struct {
		Username string `json:"username"`
	} `json:"user"`
}

type EncodedImages struct {
	Images []EncodedImage
}

type EncodedImage struct {
	Name    string
	Encoded string
}

type ItemResult struct {
	ObjectName string `json:"object_name"`
	Etag       string `json:"etag"`
}
