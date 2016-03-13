package beanWork

import (
	"github.com/kr/beanstalk"
	"sync"
)

type(
	// callback function
	JobHandler func(*BeanJob)

	//
	BeanWorker struct {
		Net     string
		Address string
	}

)


// create number of worker as go routines
func (bw *BeanWorker)Worker(tube string, numberOfWorkers int, fn JobHandler) {
	for i := 0; i < numberOfWorkers; i++ {
		go bw.work(tube, fn)
	}
}

// create new beanstalk connection and return it
// panic if got an connection error
func (bw *BeanWorker)getNewConnection() *beanstalk.Conn {
	c, err := beanstalk.Dial(bw.Net, bw.Address)
	if err != nil {
		panic(err)
	}
	return c
}

// create all resource required running an active tube connection
// each worker gets its own connection to prevent race conditions
func (bw *BeanWorker)work(tube string, fn JobHandler) {
	beanTube := beanTube{
		tubeName:tube,
		conn:bw.getNewConnection(),
		wg:sync.WaitGroup{},
	}
	beanTube.tubeSet = beanTube.getTubeSet()
	beanTube.reserveTubeJobs(fn)
}

// run waits for incoming chan messages
func (bw *BeanWorker)Run() {
	for {
		select {
		}
	}
}