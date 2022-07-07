package ds

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapGetToReturnZeroValAndFalseIfKeyDoesNotExist(t *testing.T) {
	// Arrange
	sut := NewMap[string, int]()

	// Act
	result, ok := sut.Get("test")

	// Assert
	assert.Equal(t, 0, result)
	assert.False(t, ok)
}

func TestMapGetToReturnValAndTrueIfKeyExists(t *testing.T) {
	// Arrange
	sut := NewMap[string, int]()
	sut.Set("test", 1)

	// Act
	result, ok := sut.Get("test")

	// Assert
	assert.Equal(t, 1, result)
	assert.True(t, ok)
}

func TestMapGetToWorkFineWithConcurrentReadAndWrite(t *testing.T) {
	// Arrange
	sut := NewMap[string, int]()
	count := 10000000
	wg := &sync.WaitGroup{}
	wg.Add(count)

	// Act & Assert
	for i := 0; i < count; i++ {
		go func(i int) {
			sut.Set("test", i)
			sut.Get("test")

			wg.Done()
		}(i)
	}

	wg.Wait()

	// Didn't receive concurrent access panic 👍
}
