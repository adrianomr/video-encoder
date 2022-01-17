package utils_test

import (
	"testing"
	"video-encoder/utils"

	"github.com/stretchr/testify/require"
)

func TestIsJson(t *testing.T) {
	json := `{"test": 1234}`

	err := utils.IsJson(json)

	require.Nil(t, err)
}

func TestIsNotJson(t *testing.T) {
	json := `asd`

	err := utils.IsJson(json)

	require.Error(t, err)
}