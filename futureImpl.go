package main

import (
	"time"
)

func newFuture(f func() Result) *FutureTask {
	taskChannel := make(chan Result)
	errorChannel := make(chan Result)
	resultChannel := make(chan Result)

	futureTask := FutureTask{
		success:       false,
		done:          false,
		canceled:      false,
		result:        Result{},
		taskChannel:   taskChannel,
		errorChannel:  errorChannel,
		resultChannel: resultChannel,
	}

	go func() {
		result := f()
		taskChannel <- result
	}()

	go func() {
		select {
		case res := <-futureTask.taskChannel:
			futureTask.result = res
			futureTask.success = true
			futureTask.done = true
			resultChannel <- futureTask.result
		case res2 := <-futureTask.errorChannel:
			futureTask.result = res2
			futureTask.success = false
			if futureTask.result.errorMessage == "cancelled" {
				futureTask.canceled = true
			}
			futureTask.done = true
			resultChannel <- futureTask.result
		}
	}()

	return &futureTask
}

func (futureTask *FutureTask) get() Result {
	if futureTask.isComplete() || futureTask.isCancelled() {
		return futureTask.result
	}

	futureTask.result = <-futureTask.resultChannel
	return futureTask.result
}

func (futureTask *FutureTask) getWithTimeout(timeout time.Duration) Result {
	if futureTask.isComplete() || futureTask.isCancelled() {
		return futureTask.result
	}

	timeoutChannel := time.After(timeout)
	select {
	case res := <-futureTask.resultChannel:
		futureTask.result = res
	case <-timeoutChannel:
		futureTask.done = true
		futureTask.success = false
		futureTask.result = Result{resultValue: nil, errorMessage: "timeout"}
	}
	return futureTask.result
}

func (futureTask *FutureTask) isComplete() bool {
	if futureTask.done {
		return true
	} else {
		return false
	}
}

func (futureTask *FutureTask) isCancelled() bool {
	if futureTask.done && futureTask.canceled {
		return true
	}

	return false
}

func (futureTask *FutureTask) cancel() {
	if futureTask.isComplete() || futureTask.isCancelled() {
		return
	}

	futureTask.errorChannel <- Result{resultValue: nil, errorMessage: "cancelled"}
}
