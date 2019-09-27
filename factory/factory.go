// Credit: https://github.com/ipsusila/so

package factory

import (
	"errors"
	"strings"
	"sync"
)

type Factory interface {
	New(name string) interface{}
}

var (
	mu sync.RWMutex
	factories = make(map[string]Factory)
)

func Register(name string, f Factory) error {
	mu.Lock()
	defer mu.Unlock()

	if f == nil {
		return errors.New("Factory is nil")
	}
	if _, exist := factories[name]; exist {
		return errors.New("Factory is already registered")
	}

	factories[name] = f
	return nil
}

func New(typeName string) interface{} {
	items := strings.Split(typeName, ".")
	if len(items) >= 2 {
		mu.RLock()
		defer mu.RUnlock()
		if f, exist := factories[items[0]]; exist {
			return f.New(items[1])
		}
	}
	return nil
}
