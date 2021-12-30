package domain_test

import (
	"github.com/stretchr/testify/require"
	"testing"
	"video-encoder/domain"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}
