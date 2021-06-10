package handlers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_jsonError(t *testing.T) {
	message := "Hello World"
	json := jsonError(message)
	require.Equal(t, `{"message":"Hello World"}`, json)
}
