package main

import "time"

type Future interface {
	get() Result
	getWithTimeout(duration time.Duration) Result
	isComplete() bool
	isCancelled() bool
	cancel()
}
