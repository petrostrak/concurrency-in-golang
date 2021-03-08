package deadlock

import (
	"time"

	"github.com/concurrency-in-golang/deadlock_train/common"
)

// MoveTrain func
func MoveTrain(train *common.Train, distance int, crossing []*common.Crossing) {
	for train.Front < distance {
		train.Front++
		for _, crossing := range crossing {
			if train.Front == crossing.Position {
				crossing.Intersection.Mutex.Lock()
				crossing.Intersection.LockedBy = train.ID
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
