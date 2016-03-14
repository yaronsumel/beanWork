package beanWork

import (
	"testing"
	"github.com/kr/beanstalk"
	"reflect"
)

func getTestBeanTube(conn *beanstalk.Conn,tubeName string)*beanTube{
	return &beanTube{
		conn:conn,
		tubeName:tubeName,
	}
}

func TestGetTubeSet(t *testing.T){

	c:=getTestConnection(t)
	defer c.Close()
	tube := getTestTube(c,t)

	bt := getTestBeanTube(c,tube.Name)

	tsA := bt.getTubeSet()
	tsB := getTestTubeSet(c,tube.Name)

	if !reflect.DeepEqual(tsA,tsB){
		t.Fatalf("FAILED AT %s","DeepEqual")
	}


}