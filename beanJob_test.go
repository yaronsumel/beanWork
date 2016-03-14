package beanWork

import (
	"testing"
	"github.com/kr/beanstalk"
	"time"
	"strconv"
	"reflect"
	"math/rand"
)

func getTestConnection(t *testing.T)*beanstalk.Conn{
	c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
	if err != nil{
		t.Fatalf("getTestConnection %s",err)
	}
	return c
}

func getTestTube(c *beanstalk.Conn,t *testing.T)*beanstalk.Tube{
	randTube := strconv.Itoa(int(time.Now().Unix())+rand.Int())
	return 	&beanstalk.Tube{c, randTube}
}

func putTestJob(jobBody []byte,tube *beanstalk.Tube,t *testing.T)uint64{
	jobId,err := tube.Put(jobBody, 1, 0, 120*time.Second)
	if err != nil{
		t.Fatalf("putTestJob %s",err)
	}
	return jobId
}

func getTestBeanJob(c *beanstalk.Conn,jobId uint64,t *testing.T)*BeanJob{
	return &BeanJob{conn:c,Id:jobId}
}

func getTestTubeSet(c *beanstalk.Conn,tubeName string) *beanstalk.TubeSet {
	return beanstalk.NewTubeSet(c, tubeName)
}

func TestDelete(t *testing.T){

	c:=getTestConnection(t)
	defer c.Close()

	tube := getTestTube(c,t)
	jobId := putTestJob([]byte("job-body"),tube,t)
	BeanJob := getTestBeanJob(c,jobId,t)

	if BeanJob.Delete()!=nil{
		t.Fatalf("FAILED AT %s","BeanJob.Delete()")
	}

	if _,err := c.Peek(jobId);err==nil{
		t.Fatalf("FAILED AT %s","err is nil")
	}

}

func TestRelease(t *testing.T) {

	c:=getTestConnection(t)
	defer c.Close()

	tube := getTestTube(c,t)
	jobId := putTestJob([]byte("job-body"),tube,t)
	BeanJob := getTestBeanJob(c,jobId,t)

	getTestTubeSet(c,tube.Name).Reserve(time.Second * 0)

	if err:=BeanJob.Release(1,0);err!=nil{
		t.Fatalf("FAILED AT %s",err.Error())
	}

	jobStats,err:=BeanJob.StatsJob()

	if err != nil{
		t.Fatalf("FAILED AT %s",err.Error())
	}

	if jobStats["releases"] != "1"  {
		t.Fatalf("FAILED AT %s","releases != 1")
	}

}

func TestBury(t *testing.T) {

	c:=getTestConnection(t)
	defer c.Close()

	tube := getTestTube(c,t)
	jobId := putTestJob([]byte("job-body"),tube,t)
	BeanJob := getTestBeanJob(c,jobId,t)

	getTestTubeSet(c,tube.Name).Reserve(time.Second * 0)

	if err:=BeanJob.Bury(1);err!=nil{
		t.Fatalf("FAILED AT %s",err.Error())
	}

	jobStats,err:=BeanJob.StatsJob()

	if err != nil{
		t.Fatalf("FAILED AT %s",err.Error())
	}

	if jobStats["buries"] != "1"  {
		t.Fatalf("FAILED AT %s","buries != 1")
	}

}

func TestTouch(t *testing.T) {

	c:=getTestConnection(t)
	defer c.Close()

	tube := getTestTube(c,t)
	jobId := putTestJob([]byte("job-body"),tube,t)
	BeanJob := getTestBeanJob(c,jobId,t)

	if err:=BeanJob.Touch();err==nil{
		t.Fatalf("FAILED AT %s",err.Error())
	}

	getTestTubeSet(c,tube.Name).Reserve(time.Second * 0)

	if err:=BeanJob.Touch();err!=nil{
		t.Fatalf("FAILED AT %s",err.Error())
	}

}

func TestPeek(t *testing.T) {

	c:=getTestConnection(t)
	defer c.Close()

	jobBody := []byte("job-body")
	tube := getTestTube(c,t)
	jobId := putTestJob(jobBody,tube,t)
	BeanJob := getTestBeanJob(c,jobId,t)

	cBody,cErr := c.Peek(jobId)

	if cErr != nil{
		t.Fatalf("FAILED AT %s",cErr.Error())
	}

	if !reflect.DeepEqual(cBody, jobBody)  {
		t.Fatalf("FAILED AT %s","DeepEqual(cBody, jobBody)")
	}

	body,err := BeanJob.conn.Peek(jobId)

	if err != nil{
		t.Fatalf("FAILED AT %s",err.Error())
	}

	if !reflect.DeepEqual(body, jobBody) || !reflect.DeepEqual(body, cBody) {
		t.Fatalf("FAILED AT %s","reflect.DeepEqual")
	}

}

func TestStatsJob(t *testing.T) {

	c:=getTestConnection(t)
	defer c.Close()

	tube := getTestTube(c,t)
	jobId := putTestJob([]byte("job-body"),tube,t)
	BeanJob := getTestBeanJob(c,jobId,t)

	jobStats,err:=BeanJob.StatsJob()

	if err != nil {
		t.Fatalf("FAILED AT %s",err.Error())
	}

	if reflect.TypeOf(jobStats["id"]).Kind() != reflect.String{
		t.Fatalf("FAILED AT %s","reflect.TypeOf")
	}

	i,err := strconv.Atoi(jobStats["id"])

	if err != nil{
		t.Fatalf("FAILED AT %s",err.Error())
	}

	if uint64(i) != jobId{
		t.Fatalf("FAILED AT %s","uint64(i) != jobId")
	}

}
