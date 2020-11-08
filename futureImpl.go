package main

import "time"

func newFuture(f func() Result) *FutureTask {
	channel := make(chan Result)

	futureTask := FutureTask{
		success:          false,
		done:             false,
		canceled:         false,
		result:           Result{},
		interfaceChannel: channel,
	}

	go func() {
		result := f()
		channel <- result
	}()

	return &futureTask
}

func (futureTask *FutureTask) get() Result {
	if futureTask.done {
		return futureTask.result
	}
	futureTask.result = <-futureTask.interfaceChannel
	futureTask.success = true
	futureTask.done = true
	return futureTask.result
}

func (futureTask *FutureTask) getWithTimeout(timeout time.Duration) Result {
	if futureTask.done {
		return futureTask.result
	}
	timeoutChannel := time.After(timeout)
	select {
	case res := <-futureTask.interfaceChannel:
		futureTask.result = res
		futureTask.success = true
		futureTask.done = true
	case <-timeoutChannel:
		futureTask.done = true
		futureTask.success = false
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
	if futureTask.done {
		if futureTask.canceled {
			return true
		}
	}

	return false
}

func (futureTask *FutureTask) cancel() {
	if futureTask.isComplete() || futureTask.isCancelled() {
		return
	}

	futureTask.done = true
	futureTask.success = false
	futureTask.canceled = true
	futureTask.result = Result{resultValue: nil}
	futureTask.interfaceChannel <- futureTask.result
}
