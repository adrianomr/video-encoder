package services_test

import (
	"log"
	"testing"
	"time"
	"video-encoder/application/repositories"
	"video-encoder/application/services"
	"video-encoder/domain"
	"video-encoder/framework/database"

	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func prepare() (domain.Video, repositories.VideoRepositoryDb) {

	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "teste.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.NewVideoRepository(db)

	return *video, *repo
}

func TestDownload(t *testing.T) {
	video, repo := prepare()

	videoService := services.NewVideoService()
	videoService.Video = &video
	videoService.VideoRepository = repo

	download := videoService.Download("codeeducationtest")

	require.Nil(t, download)
}
