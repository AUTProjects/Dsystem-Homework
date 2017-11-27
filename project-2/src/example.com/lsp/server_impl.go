// Contains the implementation of a LSP server.

package lsp

import (
	"encoding/json"
	"errors"
	"fmt"

	net "../lspnet"
)

type server struct {
	udpConn    *net.UDPConn
	lastConnID int
}

// NewServer creates, initiates, and returns a new server. This function should
// NOT block. Instead, it should spawn one or more goroutines (to handle things
// like accepting incoming client connections, triggering epoch events at
// fixed intervals, synchronizing events using a for-select loop like you saw in
// project 0, etc.) and immediately return. It should return a non-nil error if
// there was an error resolving or listening on the specified port number.
func NewServer(port int, params *Params) (Server, error) {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	uc, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}
	s := &server{
		udpConn: uc,
	}
	// Handles incomming messages
	go s.handle()
	return s, nil
}

func (s *server) handle() {
	for {
		buff := make([]byte, 1024)
		nbytes, addr, err := s.udpConn.ReadFromUDP(buff)
		if err != nil || nbytes == 0 {
			continue
		}
		buff = buff[:nbytes]

		// Message parsing
		var mr Message
		err = json.Unmarshal(buff, &mr)
		if err != nil {
			continue
		}

		// Connection setup
		if mr.Type == MsgConnect {
			connID := s.lastConnID
			s.lastConnID++
			mt := NewAck(connID, 0)
			b, _ := json.Marshal(mt)
			s.udpConn.WriteToUDP(b, addr)
		}
	}
}

func (s *server) Read() (int, []byte, error) {
	// TODO: remove this line when you are ready to begin implementing this method.
	select {} // Blocks indefinitely.
	return -1, nil, errors.New("not yet implemented")
}

func (s *server) Write(connID int, payload []byte) error {
	return errors.New("not yet implemented")
}

func (s *server) CloseConn(connID int) error {
	return errors.New("not yet implemented")
}

func (s *server) Close() error {
	return errors.New("not yet implemented")
}