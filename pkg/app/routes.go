package app

func (s *Server) setRoutes() {
	s.router.GET("/", s.handleIndex())
}