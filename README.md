# beanWork [![Build Status](https://travis-ci.org/yaronsumel/beanWork.svg?branch=master)](https://travis-ci.org/yaronsumel/beanWork)

Worker client for [beanstalkd](http://kr.github.com/beanstalkd/). Written in Go by [Yaron Sumel](http://sumel.me).

PR are welcome!

## Install

    $ go get github.com/yaronsumel/beanWork

## Usage

```go

    import "github.com/yaronsumel/beanWork"

	bw := &beanWork.BeanWorker{"tcp4","127.0.0.1:11300"}

	amountOfWorkers := 5
	bw.Worker("mytube",amountOfWorkers,func(job *beanWork.BeanJob) {
		//do some work here
		job.Delete()
	})
	
	bw.Run()
```


## TBD

    * Unit-Testing
