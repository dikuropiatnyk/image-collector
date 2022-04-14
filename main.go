package main

import (
	"image-collector/services"
	u "image-collector/utils"
	"strconv"
)

func main() {
	// Initialize environment
	services.InitEnv()

	println("Full URL is: " + u.RootImageURL + u.PhotosEndpoint + strconv.Itoa(u.DefaultImagesCount))
	println("Your access key is: " + services.GetEnvVar(u.AccessKeyName))
}
