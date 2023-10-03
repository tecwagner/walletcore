package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (et *TestEvent) GetName() string {
	return et.Name
}

func (et *TestEvent) GetPayload() interface{} {
	return et.Payload
}

func (et *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (t *TestEventHandler) Handle(event IEventInterface) {
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{ID: 1}
	suite.handler2 = TestEventHandler{ID: 2}
	suite.handler3 = TestEventHandler{ID: 3}
	suite.event = TestEvent{Name: "test", Payload: "test"}
	suite.event2 = TestEvent{Name: "test2", Payload: "test2"}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	// assert.True(suite.T(), true) Esse comando executa a minha suite de test com todos os eventos

	// Registrando 1 evento
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Registrando 2 evento
	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// O valor registrado dos nomes dos handler
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler3, suite.eventDispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {

	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Retorna o erro informando que o handler já existe
	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	// Event 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Event 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	suite.eventDispatcher.Clear()
	suite.Equal(0, len(suite.eventDispatcher.handlers))

}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {

	// Event 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler))
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler2))
	assert.False(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler3))

}

func (suite *EventDispatcherTestSuite) TestEventDipatcher_Remove() {
	// Event 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Event 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	// Removendo o 1 primeiro evento handler
	suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][0])
	
	// Removendo 2 segundo evento handler
	suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Removendo 1 primeiro evento handler3
	suite.eventDispatcher.Remove(suite.event2.GetName(), &suite.handler3)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event IEventInterface) {
	m.Called(event)
}

func (suite *EventDispatcherTestSuite) TestEventDispatch_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &suite.event)	
	
	eh2 := &MockHandler{}
	eh2.On("Handle", &suite.event)

	suite.eventDispatcher.Register(suite.event.GetName(), eh)
	suite.eventDispatcher.Register(suite.event.GetName(), eh2)

	suite.eventDispatcher.Dispatch(&suite.event)	

	eh.AssertExpectations(suite.T())
	eh2.AssertExpectations(suite.T())

	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
	eh2.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
