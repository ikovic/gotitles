package queue

import (
	"os"
)

type queue interface {
	IsEmpty() bool
	Enqueue(os.FileInfo)
	Dequeue() *os.FileInfo
}

// FileQueue describes a node in FS tree
type FileQueue struct {
	files []os.FileInfo
}

// IsEmpty tells us if queue has any elements
func (q FileQueue) IsEmpty() bool {
	return len(q.files) == 0
}

// Enqueue pushes a new item to the queue
func (q *FileQueue) Enqueue(f os.FileInfo) {
	q.files = append(q.files, f)
}

// Dequeue removes the first item from the queue and returns it
func (q FileQueue) Dequeue() *os.FileInfo {
	if q.IsEmpty() {
		return nil
	}
	firstItem := q.files[0]
	q.files = q.files[1:]
	return &firstItem
}
