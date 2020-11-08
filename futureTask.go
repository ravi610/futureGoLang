package main

type FutureTask struct {
	success          bool
	done             bool
	canceled         bool
	result           Result
	interfaceChannel chan Result
}
