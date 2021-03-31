package diagnostic

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jonas-fan/dsdd/pkg/channel"
	"github.com/jonas-fan/dsdd/pkg/logutil"
)

func newLogReaders(files []string) []*logutil.Reader {
	readers := make([]*logutil.Reader, 0, len(files))

	for _, each := range files {
		file, err := os.Open(each)

		if err != nil {
			panic(err)
		}

		readers = append(readers, logutil.NewReader(file))
	}

	return readers
}

func newLogStreams(files []string) []chan logutil.Log {
	readers := newLogReaders(files)
	streams := make([]chan logutil.Log, 0, len(files))

	for _, each := range readers {
		stream := logutil.NewStream(context.Background(), each, nil, 256)
		streams = append(streams, stream.Chan())
	}

	return streams
}

func ReadLog() {
	pattern := filepath.Join("Agent", "logs", "ds_agent*.log")
	matches, err := filepath.Glob(pattern)

	if err != nil {
		panic(err)
	}

	streams := newLogStreams(matches)
	out := make(chan logutil.Log)

	channel.Multiplex(out, streams)

	for each := range out {
		fmt.Println(each.(string))
	}
}
