package gonginx

import (
	"errors"
)

//Server represents server block
type Server struct {
	Block    IBlock
	Comments []string
}

//NewServer create a server block from a directive with block
func NewServer(directive IDirective) (*Server, error) {
	if block := directive.GetBlock(); block != nil {
		return &Server{
			Block:    block,
			Comments: directive.GetComment(),
		}, nil
	}

	return nil, errors.New("server directive must have a block")
}

//GetName get directive name to construct the statment string
func (s *Server) GetName() string { //the directive name.
	return "server"
}

//GetParameters get directive parameters if any
func (s *Server) GetParameters() []string {
	return []string{}
}

//GetBlock get block if any
func (s *Server) GetBlock() IBlock {
	return s.Block
}

//GetComment get the comment of directive
func (s *Server) GetComment() []string {
	return s.Comments
}
