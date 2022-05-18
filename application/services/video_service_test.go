package services_test

import (
	"fmt"
	"github.com/devlipe/encoder/application/repositories"
	"github.com/devlipe/encoder/application/services"
	"github.com/devlipe/encoder/domain"
	"github.com/devlipe/encoder/framework/database"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	"time"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("error loading .env file")
	}
}

func prepare() (*domain.Video, repositories.VideoRepositoryDb) {
	db := database.NewDbTest()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "nggyu.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}

	return video, repo
}

func TestVideoService_Download(t *testing.T) {
	video, repo := prepare()
	videoService := services.NewVideoService()

	videoService.Video = video
	videoService.VideoRepository = repo

	err := videoService.Download("encoder-bucket-go")
	fmt.Println(err.Error())
	require.Nil(t, err)

	err = videoService.Fragment()
	fmt.Println(err.Error())
	require.Nil(t, err)

	err = videoService.Encode()
	fmt.Println(err.Error())
	require.Nil(t, err)

	err = videoService.Finish()
	fmt.Println(err.Error())
	require.Nil(t, err)
}
