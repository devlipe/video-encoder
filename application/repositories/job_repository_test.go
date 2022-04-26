package repositories_test

import (
	"fmt"
	"github.com/devlipe/encoder/application/repositories"
	"github.com/devlipe/encoder/domain"
	"github.com/devlipe/encoder/framework/database"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	_, err := repo.Insert(video)

	require.Nil(t, err)

	job, err := domain.NewJob("output_path", "pending", video)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	_, err = repoJob.Insert(job)
	require.Nil(t, err)

	j, err := repoJob.Find(job.ID)

	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.VideoID, video.ID)

}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	_, err := repo.Insert(video)
	require.Nil(t, err)

	job, err := domain.NewJob("output_path", "pending", video)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	_, err = repoJob.Insert(job)
	require.Nil(t, err)

	job.Status = "Complete"

	_, err = repoJob.Update(job)
	require.Nil(t, err)

	j, err := repoJob.Find(job.ID)

	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.Status, job.Status)
}
