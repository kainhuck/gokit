//
// Sample usage:
//
//		package main
//
//		import (
//			"fmt"
//			"gitlab.com/tedge/go-mod-bootstrap/v2/di"
//		)
//
//		type foo struct {
//			FooMessage string
//		}
//
//		func NewFoo(m string) *foo {
//			return &foo{
//				FooMessage: m,
//			}
//		}
//
//		type bar struct {
//			BarMessage string
//			Foo        *foo
//		}
//
//		func NewBar(m string, foo *foo) *bar {
//			return &bar{
//				BarMessage: m,
//				Foo:        foo,
//			}
//		}
//
//		func main() {
//			container := di.NewContainer(
//				di.ServiceConstructorMap{
//					"foo": func(get di.Get) interface{} {
//						return NewFoo("fooMessage")
//					},
//					"bar": func(get di.Get) interface{} {
//						return NewBar("barMessage", get("foo").(*foo))
//					},
//				})
//
//			b := container.Get("bar").(*bar)
//			fmt.Println(b.BarMessage)
//			fmt.Println(b.Foo.FooMessage)
//		}
//
package di

import "sync"

type Get func(serviceName string) interface{}

type ServiceConstructor func(get Get) interface{}

type ServiceConstructorMap map[string]ServiceConstructor

type service struct {
	constructor ServiceConstructor
	instance    interface{}
}

type Container struct {
	serviceMap map[string]service
	mutex      sync.RWMutex
}

func NewContainer(serviceConstructors ServiceConstructorMap) *Container {
	c := Container{
		serviceMap: map[string]service{},
		mutex:      sync.RWMutex{},
	}
	if serviceConstructors != nil {
		c.Update(serviceConstructors)
	}
	return &c
}

func (c *Container) Update(serviceConstructors ServiceConstructorMap) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for serviceName, constructor := range serviceConstructors {
		c.serviceMap[serviceName] = service{
			constructor: constructor,
			instance:    nil,
		}
	}
}

func (c *Container) get(serviceName string) interface{} {
	service, ok := c.serviceMap[serviceName]
	if !ok {
		// Returning nil allows the DIC to be queried for a object and not panic if it doesn't exist.
		return nil
	}
	if service.instance == nil {
		service.instance = service.constructor(c.get)
		c.serviceMap[serviceName] = service
	}
	return service.instance
}

func (c *Container) Get(serviceName string) interface{} {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.get(serviceName)
}
