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

// Percorre a lista de eventos registrado e executa um por um pelo nome 
func (evd *EventDispatcher) Dispatch(event IEventInterface) error {

	if handlers, ok := evd.handlers[event.GetName()]; ok {
		for _, handler := range handlers {
			handler.Handle(event)
		}
	}

	return nil
}

// Registra os eventos por nome
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

// Remove eventos por nome
func (evd *EventDispatcher) Remove(eventName string, handler IEventHandlerInterface) error {
	if _, ok := evd.handlers[eventName]; ok {
		for i, h := range evd.handlers[eventName] {
			if h == handler {
				evd.handlers[eventName] = append(evd.handlers[eventName][:i], evd.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}

// Limpa todos os eventos existente no handler
func (evd *EventDispatcher) Clear() {

	// Refaz o map de handler
	evd.handlers = make(map[string][]IEventHandlerInterface)
}

// Verifica todos os eventos que foram executados
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
