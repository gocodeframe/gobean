package bean

import (
	"fmt"
	"sync"
)

func NewFactory() *beanFactory {
	return &beanFactory{Beans: make(map[string]interface{}, 0)}
}

type beanFactory struct {
	Beans map[string]interface{}
	mutex sync.RWMutex
}

func (this *beanFactory) AddBean(beanName string, bean interface{}) {
	this.mutex.Lock()
	if this.Beans[beanName] != nil {
		panic(fmt.Errorf("%s has exist", beanName))
		return
	}

	this.Beans[beanName] = bean
	this.mutex.Unlock()
}

func (this *beanFactory) GetBean(beanName string) interface{} {
	this.mutex.RLock()
	bean := this.Beans[beanName]
	if bean == nil {
		panic(fmt.Errorf("%s not exist", beanName))
	}
	this.mutex.RUnlock()
	return bean
}
