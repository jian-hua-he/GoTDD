package main

type Counter struct {
	value int
}

func (c Counter) Inc() {}

func (c Counter) Value() int {
	return 0
}
