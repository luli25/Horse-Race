package race

import (
	"math/rand"
	"sync"
	"time"

	horse "horse-race/internal/model"
)

type Race struct {
	Horses []*horse.Horse
	Length int
	mutex  sync.Mutex
	Winner *horse.Horse
}

func NewRace(names []string, length int) *Race {
	horses := []*horse.Horse{}
	for _, n := range names {
		horses = append(horses, &horse.Horse{Name: n})
	}

	return &Race{
		Horses: horses,
		Length: length,
	}
}

func (r *Race) RunWithUpdates(updates chan Update) {
	var wait sync.WaitGroup

	for _, h := range r.Horses {
		wait.Add(1)

		go func(h *horse.Horse) {
			defer wait.Done()

			for {
				time.Sleep(time.Millisecond * time.Duration(100+rand.Intn(200)))

				r.mutex.Lock()

				if r.Winner != nil {
					r.mutex.Unlock()
					return
				}

				h.Position += rand.Intn(3) + 1

				if h.Position >= r.Length && r.Winner == nil {
					r.Winner = h
				}

				// snapshot
				snapshot := make([]horse.Horse, len(r.Horses))
				for i, hh := range r.Horses {
					snapshot[i] = *hh
				}

				update := Update{
					Horses: snapshot,
				}

				if r.Winner != nil {
					update.Winner = r.Winner.Name
				}

				r.mutex.Unlock()

				updates <- update
			}
		}(h)
	}

	wait.Wait()
	close(updates)
}
