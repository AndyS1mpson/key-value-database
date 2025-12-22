package filesystem

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSegmentsDirectory_ForEach(t *testing.T) {
	t.Parallel()

	segmentsCount := 0
	expectedSegmentsCount := 3

	directory := NewSegmentsDirectory("test_data")
	err := directory.ForEach(func(data []byte) error {
		assert.True(t, len(data) != 0)
		segmentsCount++
		return nil
	})

	require.NoError(t, err)
	assert.Equal(t, expectedSegmentsCount, segmentsCount)
}

func TestSegmentsDirectory_ForEach_WithBreak(t *testing.T) {
	t.Parallel()

	directory := NewSegmentsDirectory("test_data")
	err := directory.ForEach(func([]byte) error {
		return errors.New("error")
	})

	assert.Error(t, err, "error")
}