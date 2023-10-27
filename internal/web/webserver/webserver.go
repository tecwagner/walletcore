package webserver

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/tecwagner/walletcore-service/pkg/security"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	return &WebServer{
		Router:        r,
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: webServerPort,
	}
}

func (s *WebServer) AddHandlerPublic(path string, handler http.HandlerFunc, isPublic bool) {
	s.Handlers[path] = handler

	if isPublic {
		s.Router.Post(path, handler)
	} else {
		s.Router.With(security.JWTAuthenticateMiddleware).Post(path, handler)
	}
}

func (s *WebServer) Start() error {

	err := http.ListenAndServe(s.WebServerPort, s.Router)
	if err != nil {
		return err
	}
	return nil
}
