package repositories_test

import (
	"testing"
	"time"
	"video-encoder/application/repositories"
	"video-encoder/domain"
	"video-encoder/framework/database"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.NewVideoRepository(db)
	repo.Insert(video)

	v, err := repo.Find(video.ID)
	require.Nil(t, err)

	job, err := domain.NewJob("output", "Pending", v)
	require.Nil(t, err)

	repoJob := repositories.NewJobRepository(db)
	repoJob.Insert(job)

	j, err := repoJob.Find(job.ID)

	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.NewVideoRepository(db)
	repo.Insert(video)

	v, err := repo.Find(video.ID)
	require.Nil(t, err)

	job, err := domain.NewJob("output", "Pending", v)
	require.Nil(t, err)

	repoJob := repositories.NewJobRepository(db)
	repoJob.Insert(job)

	job.Status = "Completed"

	repoJob.Update(job)

	updatedJob, err := repoJob.Find(job.ID)

	require.NotEmpty(t, updatedJob.ID)
	require.Nil(t, err)
	require.Equal(t, updatedJob.ID, job.ID)
	require.Equal(t, updatedJob.Status, "Completed")
}
