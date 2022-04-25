package domain_test

import (
	"github.com/devlipe/encoder/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIdIsNotUuid(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "this-is-not-an-uuid"
	video.ResourceID = "a"
	video.FilePath = "path/path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Error(t, err)
}
func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.ResourceID = "a"
	video.FilePath = "path/path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)
}
