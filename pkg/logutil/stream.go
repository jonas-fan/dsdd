package logutil

import (
	"context"
)

type Log interface{}

type Processor func(string) Log

type Stream struct {
	ctx       context.Context
	reader    *Reader
	processor Processor
	channel   chan Log
}

func fastpath(data string) Log {
	return data
}

func (s *Stream) move() {
	defer close(s.channel)

	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			if s.reader.HasNext() {
				s.channel <- s.processor(s.reader.Next())
			} else {
				return
			}
		}
	}
}

// Chan resturns the stream channel.
func (s *Stream) Chan() chan Log {
	return s.channel
}

// NewStream returns a new log stream.
func NewStream(ctx context.Context, reader *Reader, processor Processor, buffered int) *Stream {
	if processor == nil {
		processor = fastpath
	}

	stream := &Stream{
		ctx:       ctx,
		reader:    reader,
		processor: processor,
		channel:   make(chan Log, buffered),
	}

	go stream.move()

	return stream
}
