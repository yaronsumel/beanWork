package beanWork

import (
	"github.com/kr/beanstalk"
	"time"
)

type BeanJob struct {
	Id uint64
	Body []byte
	conn *beanstalk.Conn
}

// Delete deletes the given job.
func (j *BeanJob) Delete() error {
	return j.conn.Delete(j.Id)
}

// Release tells the server to perform the following actions:
// set the priority of the given job to pri, remove it from the list of
// jobs reserved by c, wait delay seconds, then place the job in the
// ready queue, which makes it available for reservation by any client.
func (j *BeanJob) Release(pri uint32, delay time.Duration) error {
	return j.conn.Release(j.Id,pri,delay)
}

// Bury places the given job in a holding area in the job's tube and
// sets its priority to pri. The job will not be scheduled again until it
// has been kicked; see also the documentation of Kick.
func (j *BeanJob) Bury(pri uint32) error {
	return j.conn.Bury(j.Id,pri)
}

// Touch resets the reservation timer for the given job.
// It is an error if the job isn't currently reserved by c.
// See the documentation of Reserve for more details.
func (j *BeanJob) Touch(id uint64) error {
	return j.conn.Touch(j.Id)
}

// Peek gets a copy of the specified job from the server.
func (j *BeanJob) Peek() (body []byte, err error) {
	return j.conn.Peek(j.Id)
}

// StatsJob retrieves statistics about the given job.
func (j *BeanJob) StatsJob() (map[string]string, error) {
	return j.conn.StatsJob(j.Id)
}
