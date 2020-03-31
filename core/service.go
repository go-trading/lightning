package core

import (
	"reflect"

	"github.com/sirupsen/logrus"
)

type ServiceStatus int

const (
	Running ServiceStatus = iota
	Stopped
)

type Service interface {
	Name() string
	Init(*Node, *ServiceConfig) error
	Start() error
	Stop() error
	Status() ServiceStatus
}

type AvailableServices map[string]reflect.Type

func NewAvailableServices() AvailableServices {
	return make(AvailableServices)
}

func (typesOfService AvailableServices) Register(name string, service Service) {
	t := reflect.TypeOf(service)
	log.WithFields(logrus.Fields{
		"name": name,
		"type": t,
	}).Debug("register service")
	typesOfService[name] = t
}
