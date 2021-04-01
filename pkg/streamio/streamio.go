package streamio

import (
	"bufio"
	"context"
	"io"
)

type Stream struct {
	ctx     context.Context
	scanner *bufio.Scanner
	channel chan []byte
}

func (s *Stream) activate() {
	defer close(s.channel)

	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			if s.scanner.Scan() {
				s.channel <- s.scanner.Bytes()
			} else {
				return
			}
		}
	}
}

// Chan returns the stream channel.
func (s *Stream) Chan() chan []byte {
	return s.channel
}

// NewStream returns a new stream.
func NewStream(ctx context.Context, reader io.Reader) *Stream {
	stream := &Stream{
		ctx:     ctx,
		scanner: bufio.NewScanner(reader),
		channel: make(chan []byte, 1024),
	}

	go stream.activate()

	return stream
}
