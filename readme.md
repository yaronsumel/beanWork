# beanWork

Worker client for [beanstalkd](http://kr.github.com/beanstalkd/). Written in Go by [Yaron Sumel](http://sumel.me).


## Install

    $ go get github.com/yaronsumel/beanWork

## Usage

```go
    package main

    import (
        "github.com/yaronsumel/beanWork"
    )


    bw := &beanWork.BeanWorker{
            Server:beanWork.BeanServer{
                Net:"tcp4",
                Address:"127.0.0.1:11300",
            },
        }

    amountOfWorkers := 5

	bw.Worker("mytube",amountOfWorkers,func(job *beanWorker.BeanJob) {
	//do some work here
		job.Delete()
	})
	
	bw.Run()
```


## TBD

    * Unit-Testing
