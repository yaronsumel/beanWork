package beanWork

import (
	"github.com/kr/beanstalk"
	"time"
	"sync"
)

type beanTube struct {
	tubeSet  *beanstalk.TubeSet
	tubeName string
	conn     *beanstalk.Conn
	wg       sync.WaitGroup
}

// reserve jobs in tube
// wait for incoming jobs and call pre-defined function as result
func (bt *beanTube)reserveTubeJobs(fn JobHandler) {

	defer bt.wg.Wait()

	for {
		id, body, err := bt.tubeSet.Reserve(0 * time.Second)

		if err != nil {
			if &bt.conn == nil{
				panic("connection is lost")
			}
			continue;
		}

		bt.wg.Add(1)

		bt.run(fn,&BeanJob{
			Id:id,
			Body:body,
			conn:bt.conn,
		})
	}
}

// run the user function
func (bt *beanTube)run(fn JobHandler,job *BeanJob){
	defer bt.wg.Done()
	fn(job)
}

// return tubeSet resource based on connection and tubeName
func (bt *beanTube)getTubeSet() *beanstalk.TubeSet {
	return beanstalk.NewTubeSet(bt.conn, bt.tubeName)
}