package infra

type Server struct {
	Logger Logger
}

func NewServer(l Logger) *Server {
	l.Log("Initialize server")
	return &Server{
		Logger: l,
	}
}
