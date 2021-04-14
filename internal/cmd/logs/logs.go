package logs

import (
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var pattern string

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
	matches, err := filepath.Glob(filepath.Join("Agent", "logs", pattern))

	if err != nil {
		panic(err)
	}

	for i := 0; i < len(matches); i++ {
		if strings.Contains(matches[i], "err") {
			matches = append(matches[:i], matches[i+1:]...)
			i--
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(matches)))

	matches = append(matches[1:], matches[0])

	pipeAll(os.Stdout, matches)
}

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "logs",
		Short: "Show logs",
		Run:   run,
	}

	flags := command.Flags()
	flags.SetInterspersed(true)
	flags.StringVarP(&pattern, "pattern", "p", "ds_agent*.log", "File pattern to parse")

	return command
}
