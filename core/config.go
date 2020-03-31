package core

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

type (
	NodeConfig    []map[string]interface{}
	ServiceConfig struct {
		raw          map[string]interface{}
		wereRead     map[string]bool
		errorBuilder strings.Builder
	}
)

func ParseConfig(data []byte) *NodeConfig {
	c := &NodeConfig{}
	err := yaml.Unmarshal(data, &c)
	if err != nil {
		log.WithError(err).WithField("data", string(data)).Error("config unmarshal error")
		return nil
	}
	return c
}

func NewServiceConfig(raw map[string]interface{}) *ServiceConfig {
	if raw == nil {
		log.Error("raw config not initialized")
	}
	return &ServiceConfig{
		raw:      raw,
		wereRead: make(map[string]bool),
	}
}

func (sc *ServiceConfig) GetInterface(name string) interface{} {
	i, ok := sc.raw[name]
	if !ok {
		sc.errorBuilder.WriteString(fmt.Sprintf("The value of %s is not defined in configuration\n", name))
	}
	sc.wereRead[name] = true
	return i
}

func (sc *ServiceConfig) GetUint(name string) uint {
	i := sc.GetInterface(name)
	result, ok := i.(uint)
	if !ok {
		sc.errorBuilder.WriteString(fmt.Sprintf("Can't convert %s(%T) to uint\n", name, i))
	}
	return result
}

func (sc *ServiceConfig) GetString(name string) string {
	i := sc.GetInterface(name)
	result, ok := i.(string)
	if !ok {
		sc.errorBuilder.WriteString(fmt.Sprintf("Can't convert %s(%T) to string\n", name, i))
	}
	return result
}

func (sc *ServiceConfig) GetStrings(name string) []string {
	i := sc.GetInterface(name)
	result, ok := i.([]string)
	if !ok {
		sc.errorBuilder.WriteString(fmt.Sprintf("Can't convert %s(%T) to []string\n", name, i))
	}
	return result
}

func (sc *ServiceConfig) GetErrors() error {
	for key, _ := range sc.raw {
		if !sc.wereRead[key] {
			sc.errorBuilder.WriteString(fmt.Sprintf("key %v was not read\n", key))
		}
	}
	if sc.errorBuilder.Len() > 0 {
		return errors.New(sc.errorBuilder.String())
	}
	return nil
}
