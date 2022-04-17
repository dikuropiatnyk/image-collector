package services

import (
	"fmt"
	"github.com/dikuropiatnyk/image-collector/models"
	"github.com/dikuropiatnyk/image-collector/utils"
)

func CollectImagesInParallel() models.EncodedImages {
	// Get images data from Unsplash
	images := GetImages()

	// Create a variable for the final results
	var encodedImages models.EncodedImages
	encodedImages.Images = make([]models.EncodedImage, 0, utils.DefaultImagesCount)

	// Create channels to gather data from goroutines
	var channels [utils.DefaultImagesCount]chan models.EncodedImage

	// Initialize routines and specific channels for them
	for i := range channels {
		channels[i] = make(chan models.EncodedImage)
		go getEncodedImageGoroutine(images[i], channels[i])
	}

	// Wait and collect result of all routines
	for i := range channels {
		encodedImages.Images = append(encodedImages.Images, <-channels[i])
	}

	return encodedImages
}

func getEncodedImageGoroutine(image models.Image, channel chan models.EncodedImage) {
	imageName := fmt.Sprintf("%s-%s.jpeg", image.Id, image.User.Username)
	encodedValue := GetImageBase64(image.Urls.Regular)

	// Add prepared object into the channel
	channel <- models.EncodedImage{
		Name:    imageName,
		Encoded: encodedValue,
	}
}
