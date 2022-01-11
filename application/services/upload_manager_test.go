package services_test

import (
	"log"
	"os"
	"testing"
	"video-encoder/application/services"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func TestVideoServiceUpload(t *testing.T) {
	video, repo := prepare()

	videoService := services.NewVideoService()
	videoService.Video = &video
	videoService.VideoRepository = repo

	download := videoService.Download("codeeducationtest-amr")

	log.Printf("Error: %v", download)

	require.Nil(t, download)

	err := videoService.Fragment()

	log.Printf("Error: %v", err)

	require.Nil(t, err)

	err = videoService.Encode()

	log.Printf("Error: %v", err)

	require.Nil(t, err)

	videoUpload := services.NewVideoUpload()
	videoUpload.OutputBucket = "codeeducationtest-amr"
	videoUpload.VideoPath = os.Getenv("localStoragePath") + "/" + video.ID

	doneUpload := make(chan string)
	go videoUpload.ProcessUpload(50, doneUpload)

	result := <-doneUpload

	require.Equal(t, result, "upload completed")

	// err = videoService.Finish()

	// require.Nil(t, err)
}
