package core

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"
)

type ServicePlusConfig struct {
	Service Service
	Config  *ServiceConfig
}
type Node struct {
	Services []ServicePlusConfig
	Config   NodeConfig
}

func generateError(err error, msg string) (*Node, error) {
	log.WithError(err).Error(msg)
	return nil, errors.New("service name not found in config")
}

func NewNode(availableServices AvailableServices, nodeConfig NodeConfig) (*Node, error) {
	n := &Node{Config: nodeConfig}

	for _, rawServiceConfig := range nodeConfig {
		serviceConfig := NewServiceConfig(rawServiceConfig)
		serviceName := serviceConfig.GetString("service")
		if serviceName == "" {
			return generateError(serviceConfig.GetErrors(), "serviceName is empty")
		}
		serviceType, ok := availableServices[serviceName]
		if !ok {
			return generateError(nil, fmt.Sprintf("service type %s not found in AvailableServices: %v", serviceName, availableServices))
		}
		newService := reflect.New(serviceType.Elem())
		service, ok := newService.Interface().(Service)
		if !ok {
			return generateError(nil, "can't create service")
		} else {
			log.WithField("serviceType", serviceType.Elem()).Debug("create new service")
			n.Services = append(n.Services, ServicePlusConfig{
				Service: service,
				Config:  serviceConfig,
			})
		}
	}

	for _, service := range n.Services {
		if err := service.Service.Init(n, service.Config); err != nil {
			log.WithError(err).Errorf("Can't init service %T", service.Service)
			return nil, err
		}
	}
	return n, nil
}

func (n *Node) Start() error {
	for i, service := range n.Services {
		if err := service.Service.Start(); err != nil {
			log.WithError(err).Errorf("Can't start service %T", service)
			for j := 0; j < i; j++ {
				n.Services[j].Service.Stop()
			}
			return err
		}
	}
	return nil
}

func (n *Node) Stop() error {
	var wg sync.WaitGroup
	for _, service := range n.Services {
		wg.Add(1)
		go func(s Service) {
			defer wg.Done()
			if err := s.Stop(); err != nil {
				log.WithError(err).Errorf("Can't stop service %T", s)
			}
		}(service.Service)
	}
	wg.Wait()
	return nil
}

type NodeType int

const (
	TestNode NodeType = iota
	ProdNode
)

func (n *Node) Type() NodeType {
	return ProdNode
}

func (n *Node) GetService(name string) Service {
	for _, s := range n.Services {
		if s.Service.Name() == name {
			return s.Service
		}
	}
	log.WithField("name", name).Debug("Service not found")
	return nil
}

func (n *Node) Symbol(symbolName string) *Symbol {
	//try parse exchange name
	split := strings.SplitN(symbolName, ":", 2)
	if len(split) == 2 {
		exchange, ok := n.GetService(strings.TrimSpace(split[0])).(Exchange)
		if ok {
			symbol := exchange.Symbol(strings.TrimSpace(split[1]))
			if symbol != nil {
				return symbol
			}
		}
	}
	//try get from any exchange
	for _, service := range n.Services {
		exchange, ok := service.Service.(Exchange)
		if ok {
			symbol := exchange.Symbol(symbolName)
			if symbol != nil {
				return symbol
			}
		}
	}
	log.WithField("name", symbolName).Debug("Symbol not found")
	return nil
}

func (n *Node) Now() time.Time {
	return time.Now()
}
