package query_bus

import (
	"fmt"
)

type QueryBus struct {
	handlers map[string]Handler // Mapping between query types and handlers.
}

func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[string]Handler),
	}
}

// Register a handler for a certain type of Query.
func (b *QueryBus) Register(query Query, handler Handler) {
	name := fmt.Sprintf("%T", query) // use the type of the query as its name
	b.handlers[name] = handler
}

// Dispatch a query to its respective handler.
func (b *QueryBus) Dispatch(query Query) (interface{}, error) {
	name := fmt.Sprintf("%T", query) // use the type of the query as its name
	handler, ok := b.handlers[name]
	if !ok {
		return nil, fmt.Errorf("no handlers registered for query of type %s", name)
	}

	return handler.Handle(query)
}

// Query interface
type Query interface {
	Data() map[string]interface{}
}

// Query Handler interface
type Handler interface {
	Handle(query Query) (interface{}, error)
}
