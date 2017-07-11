package framework

import "sync"

type Runnable interface {
	Id() string
	Description() string
	Run()
}

type Object struct {
	Id          string
	Description string
}

type store struct {
	mStore map[string]Runnable
	sync.Mutex
}

var Store = store{mStore: make(map[string]Runnable)}
