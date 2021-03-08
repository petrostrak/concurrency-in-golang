package common

import "sync"

// Train struct
type Train struct {
	ID          int
	TrainLength int
	Front       int
}

// Intersection struct
type Intersection struct {
	ID       int
	Mutex    sync.Mutex
	LockedBy int
}

// Crossing struct
type Crossing struct {
	Position     int
	Intersection *Intersection
}
