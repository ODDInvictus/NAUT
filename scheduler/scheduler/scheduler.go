package scheduler

import "github.com/zmb3/spotify/v2"

type Scheduler struct {
	client *spotify.Client
}

// func (s *Scheduler) Start() {
// 	go s.start()
// }