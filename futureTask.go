package main

type FutureTask struct {
	success       bool
	done          bool
	canceled      bool
	result        Result
	taskChannel   chan Result
	errorChannel  chan Result
	resultChannel chan Result
}
