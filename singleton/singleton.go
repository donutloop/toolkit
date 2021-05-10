// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package singleton

import (
	"sync"
	"sync/atomic"
)

type ConstructorFunc func() (interface{}, error)

type Singleton interface {
	Get() (interface{}, error)
	Reset()
}

// Call to create a new singleton that is instantiated with the given constructor function.
// Constructor is not called until the first call of Get(). If constructor returns a error, it will be called again
// on the next call of Get().
func NewSingleton(constructor ConstructorFunc) Singleton {
	return &singleton{
		Constructor: constructor,
	}
}

type singleton struct {
	object interface{}
	// Constructor of object
	Constructor ConstructorFunc

	m    sync.Mutex
	done uint32
}

// todo(marcel): Rename to GetOrCreate (major break).
func (s *singleton) Get() (interface{}, error) {
	if atomic.LoadUint32(&s.done) == 1 {
		return s.object, nil
	}

	s.m.Lock()
	defer s.m.Unlock()

	if s.done == 0 {
		var err error

		s.object, err = s.Constructor()
		if err != nil {
			return nil, err
		}

		defer atomic.StoreUint32(&s.done, 1)
	}

	return s.object, nil
}

// Reset indicates that the next call of Get should actually a create instance.
func (s *singleton) Reset() {
	s.m.Lock()
	defer s.m.Unlock()
	atomic.StoreUint32(&s.done, 0)
}
