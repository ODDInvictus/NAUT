package scheduler

import "github.com/zmb3/spotify/v2"


type Queue struct {
	tracks []*QueueItem
}

// A track can be a *spotify.SimpleTrack or a *spotify.EpisodePage
// We can use an interface to represent both types
type QueueItem struct {
	track *spotify.SimpleTrack
	isShow bool
	show *spotify.EpisodePage
}

func NewQueue() *Queue {
	return &Queue{}
}

/**
Get the tracks in the queue
*/
func (q *Queue) GetTracks() []*QueueItem {
	return q.tracks
}

/**
Get the length of the queue
*/
func (q *Queue) Length() int {
	return len(q.tracks)
}

/**
Return the item that is currently playing
*/
func (q *Queue) GetCurrentItem() *QueueItem {
	return q.tracks[0]
}

/**
Remove the first element in the queue
*/
func (q *Queue) Pop() {
	q.tracks = q.tracks[1:]
}

/**
Add multiple tracks to the queue
*/
func (q *Queue) AddTracks(queueItems []*QueueItem) {
	for _, queueItem := range queueItems {
		q.AddTrack(queueItem)
	}
}

/**
Add a track to the queue
This will sync with spotify
*/
func (q *Queue) AddTrack(queueItem *QueueItem) {
	q.tracks = append(q.tracks, queueItem)
	// TODO - sync with spotify
}

