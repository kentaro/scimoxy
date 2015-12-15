package scimoxy

type Server struct {
}

func newServer() *Server {
	return &Server{}
}

func (s *Server) serviceHandlerFor(path string) ServiceHandler {
	return newSlackHandler()
}
