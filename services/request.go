package services

import (
	"bytes"
	"encoding/json"
	"github.com/dikuropiatnyk/image-collector/models"
	u "github.com/dikuropiatnyk/image-collector/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func sendRequest(method, fullUrl string, headers map[string][]string, body io.Reader, result interface{}) {
	req, err := http.NewRequest(method, fullUrl, body)

	if err != nil {
		log.Fatal(err)
	}
	req.Header = headers

	resp, err := http.DefaultClient.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Response is failed: %s!", resp.Status)
	}

	switch contentType := resp.Header.Get("content-type"); contentType {
	// Response is an image
	case "image/jpeg":
		// Process it to the base64
		imageBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		u.ConvertToBase64(imageBytes, result.(*string))
	// Response is JSON-like
	case "application/json":
		// Convert the body to JSON with an expected type
		err = json.NewDecoder(resp.Body).Decode(result)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetImages() []models.Image {
	token := "Client-ID " + GetEnvVar(u.AccessKeyName)
	fullUrl := u.RootImageURL + u.ImagesEndpoint + strconv.Itoa(u.DefaultImagesCount)

	headers := map[string][]string{
		"Connection":     []string{"close"},
		"Accept-Version": []string{"v1"},
		"Authorization":  []string{token},
	}

	var images []models.Image
	sendRequest(http.MethodGet, fullUrl, headers, nil, &images)

	if len(images) != u.DefaultImagesCount {
		log.Fatalf(
			"Count of received images is incorrect! Expected: %d, received: %d",
			u.DefaultImagesCount,
			len(images),
		)
	}

	return images
}

func GetImageBase64(imageUrl string) string {
	headers := map[string][]string{
		"Connection": []string{"close"},
	}
	var encodedImage string
	sendRequest(http.MethodGet, imageUrl, headers, nil, &encodedImage)

	return encodedImage
}

func SendImagesToProcessor(images models.EncodedImages) {
	headers := map[string][]string{
		"Connection":   []string{"close"},
		"Content-Type": []string{"application/json"},
	}

	jsonImages, err := json.Marshal(images)
	if err != nil {
		log.Fatal(err)
	}

	var items []models.ItemResult
	sendRequest(http.MethodPost, u.ImageProcessorURL+u.ProcessEndpoint, headers, bytes.NewBuffer(jsonImages), &items)

	// Display results
	for index, item := range items {
		log.Printf("Image №%d response data: name - %s, etag - %s\n", index, item.ObjectName, item.Etag)
	}
}
