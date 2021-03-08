package hierarchy

import (
	"sort"
	"time"

	"github.com/concurrency-in-golang/deadlock_train/common"
)

func lockIntersectionInDistance(ID, reserveStart, reserveEnd int, crossing []*common.Crossing) {
	var intersectionsToLock []*common.Intersection
	for _, crossing := range crossing {
		if reserveEnd >= crossing.Position && reserveStart <= crossing.Position && crossing.Intersection.LockedBy != ID {
			intersectionsToLock = append(intersectionsToLock, crossing.Intersection)
		}
	}

	sort.Slice(intersectionsToLock, func(i, j int) bool {
		return intersectionsToLock[i].ID < intersectionsToLock[j].ID
	})

	for _, it := range intersectionsToLock {
		it.Mutex.Lock()
		it.LockedBy = ID
		time.Sleep(10 * time.Millisecond)
	}
}

// MoveTrain func
func MoveTrain(train *common.Train, distance int, crossings []*common.Crossing) {
	for train.Front < distance {
		train.Front++
		for _, crossing := range crossings {
			if train.Front == crossing.Position {
				lockIntersectionInDistance(train.ID, crossing.Position, crossing.Position+train.TrainLength, crossings)
			}
			back := train.Front - train.TrainLength
			if back == crossing.Position {
				crossing.Intersection.LockedBy = -1
				crossing.Intersection.Mutex.Unlock()
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}
