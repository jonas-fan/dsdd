package logs

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func pipeAll(writer io.Writer, filenames []string) {
	for _, each := range filenames {
		file, err := os.Open(each)

		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(writer, file); err != nil {
			panic(err)
		}

		file.Close()
	}
}

func run(cmd *cobra.Command, args []string) {
	pattern := filepath.Join("Agent", "logs", "ds_agent*.log")
	matches, err := filepath.Glob(pattern)

	if err != nil {
		panic(err)
	}

	for i := 0; i < len(matches); i++ {
		if strings.Contains(matches[i], "err") {
			matches = append(matches[:i], matches[i+1:]...)
		}
	}

	pipeAll(os.Stdout, matches)
}

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "logs",
		Short: "Show logs",
		Run:   run,
	}

	return command
}
