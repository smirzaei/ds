package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueueToInsertElements(t *testing.T) {
	// Arrange
	sut := NewQueue[int](4)

	// Act
	sut.Enqueue(1)
	sut.Enqueue(2)
	sut.Enqueue(3)
	sut.Enqueue(4)

	// Assert
	assert.True(t, sut.isFull())
	assert.Equal(t, []int{1, 2, 3, 4, 0}, sut.buffer)
}

func TestDequeueToRemoveElements(t *testing.T) {
	// Arrange
	sut := NewQueue[int](4)

	// Act && Assert
	assert.True(t, sut.isEmpty())
	assert.False(t, sut.isFull())

	sut.Enqueue(1)
	assert.False(t, sut.isEmpty())

	sut.Enqueue(2)
	sut.Enqueue(3)

	x, ok := sut.Dequeue()
	assert.Equal(t, 1, x)
	assert.True(t, ok)

	x, ok = sut.Dequeue()
	assert.Equal(t, 2, x)
	assert.True(t, ok)

	x, ok = sut.Dequeue()
	assert.Equal(t, 3, x)
	assert.True(t, ok)

	x, ok = sut.Dequeue()
	assert.Equal(t, 0, x)
	assert.False(t, ok)

	sut.Enqueue(55)

	x, ok = sut.Dequeue()
	assert.Equal(t, 55, x)
	assert.True(t, ok)
}

func TestEnqueueToGrowIfNeeded(t *testing.T) {
	// Arrange
	sut := NewQueue[int](2)

	// Act
	sut.Enqueue(1)
	sut.Enqueue(2)
	sut.Enqueue(3)
	sut.Enqueue(4)
	sut.Enqueue(5)
	sut.Enqueue(6)

	// Assert
	assert.False(t, sut.isFull())
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 0, 0, 0}, sut.buffer)
}
