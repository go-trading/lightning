package core

import (
	"testing"

	"github.com/sirupsen/logrus"
)

type TestService struct{ config *ServiceConfig }

func (ts *TestService) Name() string { return "Test-" + ts.config.GetString("key") }
func (ts *TestService) Init(_ *Node, config *ServiceConfig) error {
	ts.config = config
	return nil
}
func (ts *TestService) Start() error                          { return nil }
func (ts *TestService) Stop() error                           { return nil }
func (ts *TestService) Status() ServiceStatus                 { return 0 }
func (ts *TestService) SubscribeStatus(func(ServiceStatus))   {}
func (ts *TestService) UnsubscribeStatus(func(ServiceStatus)) {}

func TestNewNode(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	services := NewAvailableServices()
	services.Register("TestService1", &TestService{})
	services.Register("TestService2", &TestService{})
	services.Register("TestService3", &TestService{})

	config := NodeConfig{
		{
			"service": "TestService2",
			"key":     "key2",
		},
		{
			"service": "TestService3",
			"key":     "key3",
		},
	}

	node, err := NewNode(services, config)
	if err != nil {
		t.Error(err)
	}
	if len(node.Services) != 2 {
		t.Errorf("len(node.Services)=%v", len(node.Services))
	}
	if node.GetService("Test-key2") == nil {
		t.Error("service Test-key2 not found")
	}
	if node.GetService("Test-key3") == nil {
		t.Error("service Test-key3 not found")
	}
}
