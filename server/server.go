package server

import (
	"net/http"
	"server-template/service"
)

const (
	TEXTContentType = "text/plain"
	JSONContentType = "application/json"
	HTMLContentType = "text/html"
)

type Server struct {
	service service.IService
}

func NewServer(service service.IService) (*Server, error) {
	return &Server{service: service}, nil
}

func (s *Server) HandlerForSayHello() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		name := "John Doe"

		nameParam, ok := r.URL.Query()["name"]
		if ok && len(nameParam) > 0 {
			name = nameParam[0]
		}

		w.Header().Set("Content-Type", TEXTContentType)

		sayhello, err := s.service.SayHello(name)

		if err == nil {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(sayhello))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
