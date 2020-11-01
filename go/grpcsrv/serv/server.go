package serv


type Servlet struct {
    Conf Conf
}

func New(conf Conf) (s *Servlet, err error) {
	// make server
	s = &Servlet{
		Conf: conf,
	}
    return
}

func (s *Servlet) Free() {
    return
}

func (s *Servlet) StartOrDie() (err error) {
    return
}
