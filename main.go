package main

import (
	"github.com/dikuropiatnyk/image-collector/services"
	"log"
	"time"
)

func main() {
	start := time.Now()
	// Initialize environment
	services.InitEnv()
	encodedImages := services.CollectImagesInParallel()
	services.SendImagesToProcessor(encodedImages)
	elapsed := time.Since(start)
	log.Printf("Total execution time: %s\n", elapsed)
}
