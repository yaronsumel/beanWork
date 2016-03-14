package beanWork

import (
	"testing"
)

func getTestBeanWorker()*BeanWorker{
	return &BeanWorker{"tcp", "127.0.0.1:11300"}
}

// create number of worker as go routines
func TestWorker(t *testing.T) {

	c:=getTestConnection(t)
	tube := getTestTube(c,t)
	jobId := putTestJob([]byte("job"),tube,t)
	bw := getTestBeanWorker()

	bw.Worker(tube.Name,1,func(job *BeanJob) {
		defer job.Delete()
		if jobId != job.Id{
			t.Fatalf("FAILED AT %s","jobId != job.Id")
		}
	})

}

func TestGetNewConnection(t *testing.T){

	bw := getTestBeanWorker()

	defer func() {
		if r := recover(); r != nil {
			t.FailNow()
		}
	}()

	c := bw.getNewConnection()

	if _,err := c.Stats();err!=nil{
		t.Fatalf("FAILED AT %s",err.Error())
	}

}

func TestWork(t *testing.T) {
}

func TestRun(t *testing.T) {
}