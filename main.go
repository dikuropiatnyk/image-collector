package main

import (
	"github.com/dikuropiatnyk/image-collector/services"
	"log"
	"time"
)

func main() {
	start := time.Now()
	log.Println("Start collecting...")
	// Initialize environment
	services.InitEnv()

	encodedImages := services.CollectImagesInParallel()
	timeOfCollecting := time.Since(start)
	log.Printf("Image collecting time: %s\n", timeOfCollecting)

	log.Println("Send images to Processor")
	services.SendImagesToProcessor(encodedImages)
	elapsed := time.Since(start)
	log.Printf("Total execution time: %s\n", elapsed)
}
