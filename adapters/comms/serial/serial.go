package serial

import (
	"io"
	"time"

	"github.com/icez/sofar_g3_lsw3_logger_reader/ports"
	"go.bug.st/serial"
)

type serialPort struct {
	name       string
	serialPort io.ReadWriteCloser

	baud       int
	dataBits   int
	parityMode serial.Parity
	stopBits   serial.StopBits
}

// New creates a new instance of a serial port
func New(name string, baud int, dataBits int, parityMode serial.Parity, stopBits serial.StopBits) ports.CommunicationPort {
	return &serialPort{
		name: name,

		baud:       baud,
		dataBits:   dataBits,
		parityMode: parityMode,
		stopBits:   stopBits,
	}
}

func (s *serialPort) Open() error {
	if s.serialPort != nil {
		s.Close()
	}

	mode := &serial.Mode{
		BaudRate: s.baud,
		Parity:   s.parityMode,
		DataBits: s.dataBits,
		StopBits: s.stopBits,
	}

	var err error

	s.serialPort, err = serial.Open(s.name, mode)

	if err != nil {
		return err
	}

	return nil
}

func (s *serialPort) Close() error {
	if s.serialPort != nil {
		err := s.serialPort.Close()
		s.serialPort = nil
		return err
	}

	return nil

}

func (s *serialPort) Read(buf []byte) (int, error) {
	return s.serialPort.Read(buf)
}

func (s *serialPort) Write(payload []byte) (int, error) {
	return s.serialPort.Write(payload)
}

func (s *serialPort) SetWriteDeadline(t time.Time) error {
	return nil
}

func (s *serialPort) SetReadDeadline(t time.Time) error {
	return nil
}
