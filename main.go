package main

import "fmt"

func main() {
	d := NewDequeue(3, 3)

	d.PushFront(2)
	d.PushFront(5)
	d.PushFront(3)
	fmt.Println(d.Storage[1])
}

type Dequeue struct {
	DSize       int
	BSize       int
	Storage     []*Bucket
	FrontBucket Pointer
	BackBucket  Pointer
}

type Bucket struct {
	Capacity  int
	FreeSpace int
	Storage   []OptionalInt
}

type Pointer struct {
	currBucket   int
	lastValIndex int
}

type OptionalInt struct {
	Value int
	IsSet bool
}

func NewBucket(capacity int) *Bucket {
	return &Bucket{
		Capacity:  capacity,
		FreeSpace: capacity,
		Storage:   make([]OptionalInt, capacity),
	}
}

func NewDequeue(dsize int, bsize int) *Dequeue {
	storage := make([]*Bucket, dsize)
	isEven := dsize%2 == 0
	frontBucketIndex := 0
	backBucketIndex := 0
	if isEven {
		frontBucketIndex = dsize/2 - 1
		backBucketIndex = dsize / 2
	} else {
		frontBucketIndex = dsize / 2
		backBucketIndex = dsize / 2
	}
	// TODO: Pointers on value
	return &Dequeue{
		BSize:   bsize,
		DSize:   dsize,
		Storage: storage,
		FrontBucket: Pointer{
			currBucket: frontBucketIndex,
		},
		BackBucket: Pointer{
			currBucket: backBucketIndex,
		},
	}
}

func (d *Dequeue) PushFront(val int) {
	frontBucket := d.Storage[d.FrontBucket.currBucket]
	value := OptionalInt{
		Value: val,
		IsSet: true,
	}
	if frontBucket == nil {
		d.Storage[d.FrontBucket.currBucket] = NewBucket(d.BSize)
		if d.FrontBucket.currBucket == d.DSize/2 {
			d.FrontBucket.lastValIndex = d.BSize / 2
		} else {
			d.FrontBucket.lastValIndex = 0
		}
		d.PushFront(val)
	} else {
		currIndex := d.FrontBucket.lastValIndex
		if frontBucket.Storage[currIndex].IsSet {
			if currIndex+1 < len(frontBucket.Storage) {
				currIndex++
				d.FrontBucket.lastValIndex++
				frontBucket.Storage[currIndex] = value
			} else {
				if d.FrontBucket.currBucket+1 > len(d.Storage) {
					fmt.Printf("No more buckets") // TODO: Add migration
				} else {
					d.FrontBucket.currBucket++
					d.FrontBucket.lastValIndex = 0
					d.PushFront(val)
				}
			}
		} else {
			frontBucket.Storage[currIndex] = value
		}
	}
}
