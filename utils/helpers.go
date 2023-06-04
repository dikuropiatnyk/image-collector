package utils

import "encoding/base64"

func ConvertToBase64(bytes []byte, result *string) {
	*result = base64.StdEncoding.EncodeToString(bytes)
}

func GenerateCounts() (counts []int) {
	// Get all needed full pages
	fullPages := DefaultImagesCount / MaxImagesPerPage

	for i := 0; i < fullPages; i++ {
		counts = append(counts, MaxImagesPerPage)
	}
	finalPage := DefaultImagesCount % MaxImagesPerPage

	if finalPage > 0 {
		counts = append(counts, finalPage)
	}

	return
}
