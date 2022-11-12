package tcpip

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"github.com/icez/sofar_g3_lsw3_logger_reader/ports"
)

type tcpIpPort struct {
	name string
	conn net.Conn
}

func New(portName string) ports.CommunicationPort {
	return &tcpIpPort{
		name: portName,
	}
}

func (s *tcpIpPort) Open() error {

	var err error

	d := net.Dialer{Timeout: 3 * time.Second}
	s.conn, err = d.Dial("tcp", s.name)

	if err != nil {
		return err
	}

	return nil
}

func (s *tcpIpPort) Close() error {

	if s.conn != nil {
		err := s.conn.Close()
		s.conn = nil
		return err
	}
	return nil
}

func (s *tcpIpPort) Read(buf []byte) (int, error) {
	if s.conn == nil {
		return 0, fmt.Errorf("connection is not open")
	}

	reader := bufio.NewReader(s.conn)
	return reader.Read(buf)
}

func (s *tcpIpPort) Write(payload []byte) (int, error) {
	if s.conn == nil {
		return 0, fmt.Errorf("connection is not open")
	}
	s.conn.SetWriteDeadline(time.Now().Add(20 * time.Second))
	return s.conn.Write(payload)
}

func (s *tcpIpPort) SetWriteDeadline(t time.Time) error {
	if s.conn == nil {
		return fmt.Errorf("connection is not open")
	}
	return s.conn.SetWriteDeadline(t)
}

func (s *tcpIpPort) SetReadDeadline(t time.Time) error {
	if s.conn == nil {
		return fmt.Errorf("connection is not open")
	}
	return s.conn.SetReadDeadline(t)
}
