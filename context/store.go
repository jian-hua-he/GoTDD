package main

type Store interface {
	Fetch() string
	Cancel()
}
