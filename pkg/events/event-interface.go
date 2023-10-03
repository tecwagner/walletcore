package events

import "time"

//Evento
type IEventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{} // Podemos ter varios payload criado
}

// Operação que executa o evento
type IEventHandlerInterface interface {
	Handle(event IEventInterface)
}

// Gerenciado dos eventos Registra, Dispatcha, Remove e Limpa
type IEventDispatcherInterface interface {
	// Regitra o nome do evento e executa o evento
	Register(eventName string, handler IEventHandlerInterface) error
	// Faz com que os Handlers sejam executados
	Dispatch(event IEventInterface) error
	// Remove o evento da fila
	Remove(eventName string, handler IEventHandlerInterface) error
	// Verifica se tem um event name com esse evento e retorna true ou false
	Has(eventName string, handler IEventHandlerInterface) bool
	// Limpa todos os eventos da fila
	Clear() error
}
