package handler

import (
	"github.com/sample-tech-test/usecase"
)

type Server struct {
	UseCase usecase.UsecaseInterface
}

func NewServer(usecase usecase.UsecaseInterface) *Server {
	return &Server{UseCase: usecase}
}
