package services

import (
	"encoding/json"
	"github.com/dikuropiatnyk/image-collector/models"
	u "github.com/dikuropiatnyk/image-collector/utils"
	"log"
	"net/http"
	"strconv"
)

func sendRequest(method, fullUrl string, headers map[string][]string, result interface{}) {
	req, err := http.NewRequest(method, fullUrl, nil)

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

	//Convert the body to JSON and get preview
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPhotos() {
	token := "Client-ID " + GetEnvVar(u.AccessKeyName)
	fullUrl := u.RootImageURL + u.PhotosEndpoint + strconv.Itoa(u.DefaultImagesCount)

	headers := map[string][]string{
		"Connection":     []string{"close"},
		"Accept-Version": []string{"v1"},
		"Authorization":  []string{token},
	}
	photos := &[]models.Photo{}
	sendRequest(http.MethodGet, fullUrl, headers, photos)

	// Sample of results processing
	for _, photo := range *photos {
		log.Printf("Image name: %s-%s.jpeg\n", photo.Id, photo.User.Username)
		log.Println("Image link:", photo.Urls.Regular)
	}
}

// TODO: To provide helper to send data to Image Processor
