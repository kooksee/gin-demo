package app

import (
	"sync"
)

type singleton struct {
	services map[string]interface{}
}

func (s *singleton) Set(name string, service interface{}) {
	s.services[name] = service
}

func (s *singleton) Get(name string) interface{} {
	return s.services[name]
}

var (
	once sync.Once
	instance *singleton
)

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{services: make(map[string]interface{})}
	})

	return instance
}
