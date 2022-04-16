package main

import (
	"github.com/dikuropiatnyk/image-collector/services"
	u "github.com/dikuropiatnyk/image-collector/utils"
	"strconv"
)

func main() {
	// Initialize environment
	services.InitEnv()

	println("Full URL is: " + u.RootImageURL + u.PhotosEndpoint + strconv.Itoa(u.DefaultImagesCount))
	services.GetPhotos()
}
