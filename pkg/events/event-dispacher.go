package events

import (
	"errors"
)

// Criar midleawer de erros
var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	// O Handler pode executar varios eventos
	handlers map[string][]IEventHandlerInterface
}

// Para implementar o metodo NewEventDispatcher() é preciso implementar os handlers
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]IEventHandlerInterface),
	}
}

func (evd *EventDispatcher) Dispatch(event IEventInterface) error {

	if handlers, ok := evd.handlers[event.GetName()]; ok {
		for _, handler := range handlers {
			handler.Handle(event)
		}
	}

	return nil
}

func (evd *EventDispatcher) Register(eventName string, handler IEventHandlerInterface) error {

	if _, ok := evd.handlers[eventName]; ok {
		for _, h := range evd.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	evd.handlers[eventName] = append(evd.handlers[eventName], handler)
	return nil
}

func (evd *EventDispatcher) Clear() {

	// Refaz o map de handler
	evd.handlers = make(map[string][]IEventHandlerInterface)
}

func (evd *EventDispatcher) Has(eventName string, handler IEventHandlerInterface) bool {

	// O evento existe
	if _, ok := evd.handlers[eventName]; ok {
		// Percore todos os registros
		for _, h := range evd.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}
