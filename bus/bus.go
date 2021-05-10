package bus

import (
	"fmt"
	"reflect"
	"sync"
)

// The type of the function's first and only argument.
// declares the msg to listen for.
type HandlerFunc interface{}

type Msg interface{}

// It is a simple but powerful publish-subscribe event system. It requires object to
// register themselves with the event bus to receive events.
type Bus interface {
	Dispatch(msg Msg) error
	AddHandler(handler HandlerFunc) error
	AddEventListener(handler HandlerFunc)
	Publish(msg Msg) error
}

type InProcBus struct {
	sync.Mutex
	handlers  map[string]reflect.Value
	listeners map[string][]reflect.Value
}

func New() Bus {
	return &InProcBus{
		handlers:  make(map[string]reflect.Value),
		listeners: make(map[string][]reflect.Value),
	}
}

// Dispatch sends an msg to registered handler that were declared
// to accept values of a msg.
func (b *InProcBus) Dispatch(msg Msg) error {
	nameOfMsg := reflect.TypeOf(msg)

	handler, ok := b.handlers[nameOfMsg.String()]
	if !ok {
		return &ErrHandlerNotFound{Name: nameOfMsg.Name()}
	}

	params := make([]reflect.Value, 0, 1)
	params = append(params, reflect.ValueOf(msg))

	ret := handler.Call(params)

	v := ret[0].Interface()
	if err, ok := v.(error); ok && err != nil {
		return err
	}

	return nil
}

// Publish sends an msg to all registered listeners that were declared
// to accept values of a msg.
func (b *InProcBus) Publish(msg Msg) error {
	nameOfMsg := reflect.TypeOf(msg)
	listeners := b.listeners[nameOfMsg.String()]

	params := make([]reflect.Value, 0, 1)
	params = append(params, reflect.ValueOf(msg))

	for _, listenerHandler := range listeners {
		ret := listenerHandler.Call(params)

		v := ret[0].Interface()
		if err, ok := v.(error); ok && err != nil {
			return err
		}
	}

	return nil
}

// AddHandler registers a handler function that will be called when a matching
// msg is dispatched.
func (b *InProcBus) AddHandler(handler HandlerFunc) error {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	handlerType := reflect.TypeOf(handler)
	validateHandlerFunc(handlerType)

	typeOfMsg := handlerType.In(0)
	if _, ok := b.handlers[typeOfMsg.String()]; ok {
		return &ErrOverwrite{Name: typeOfMsg.Name()}
	}

	b.handlers[typeOfMsg.String()] = reflect.ValueOf(handler)

	return nil
}

// AddListener registers a listener function that will be called when a matching
// msg is dispatched.
func (b *InProcBus) AddEventListener(handler HandlerFunc) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	handlerType := reflect.TypeOf(handler)
	validateHandlerFunc(handlerType)
	// the first input parameter is the msg
	typOfMsg := handlerType.In(0)

	_, ok := b.listeners[typOfMsg.String()]
	if !ok {
		b.listeners[typOfMsg.String()] = make([]reflect.Value, 0)
	}

	b.listeners[typOfMsg.String()] = append(b.listeners[typOfMsg.String()], reflect.ValueOf(handler))
}

// panic if conditions not met (this is a programming error).
func validateHandlerFunc(handlerType reflect.Type) {
	switch {
	case handlerType.Kind() != reflect.Func:
		panic(ErrBadFuncError("handler func must be a function"))
	case handlerType.NumIn() != 1:
		panic(ErrBadFuncError("handler func must take exactly one input argument"))
	case handlerType.NumOut() != 1:
		panic(ErrBadFuncError("handler func must take exactly one output argument"))
	}
}

// ErrBadFuncError is raised via panic() when AddEventListener or AddHandler is called with an
// invalid listener function.
type ErrBadFuncError string

func (bhf ErrBadFuncError) Error() string {
	return fmt.Sprintf("bad handler func: %s", string(bhf))
}

type ErrHandlerNotFound struct {
	Name string
}

func (e *ErrHandlerNotFound) Error() string { return e.Name + ": not found" }

type ErrOverwrite struct {
	Name string
}

func (e *ErrOverwrite) Error() string { return e.Name + ": handler exists" }
