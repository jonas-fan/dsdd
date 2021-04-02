package ps

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/dsa/diagnostic"
	"github.com/jonas-fan/dsdd/pkg/fmtutil"
	"github.com/spf13/cobra"
)

var details bool

func shake(task *diagnostic.Task, filters []string) bool {
	for _, each := range filters {
		switch {
		case task.PID == each:
			return true
		case task.User == each:
			return true
		case task.Name == each:
			return true
		case strings.HasSuffix(task.Path, each):
			return true
		}
	}

	return false
}

func sieve(tasks []diagnostic.Task, filters []string) <-chan *diagnostic.Task {
	out := make(chan *diagnostic.Task)

	go func() {
		fastpath := (len(filters) == 0)

		for _, each := range tasks {
			if fastpath || shake(&each, filters) {
				out <- &each
			}
		}

		close(out)
	}()

	return out
}

func run(cmd *cobra.Command, args []string) {
	file, err := os.Open(filepath.Join("Agent", "RunningProcesses.xml"))

	if err != nil {
		panic(err)
	}

	defer file.Close()

	tasks, err := diagnostic.ReadProcess(file)

	if err != nil {
		panic(err)
	}

	formatter := fmtutil.NewFormatter()

	if details {
		formatter.Write("USER", "PID", "PPID", "NAME", "PATH", "COMMAND")
	} else {
		formatter.Write("USER", "PID", "PPID", "COMMAND")
	}

	formatter.Align(1, fmtutil.RightAlign)
	formatter.Align(2, fmtutil.RightAlign)

	for each := range sieve(tasks, args) {
		if details {
			formatter.Write(each.User, each.PID, each.PPID, each.Name, each.Path, each.CommandLine)
		} else {
			formatter.Write(each.User, each.PID, each.PPID, each.CommandLine)
		}
	}

	fmt.Println(formatter.String())
}

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "ps",
		Short: "List running processes",
		Run:   run,
	}

	flags := command.Flags()
	flags.SetInterspersed(false)
	flags.BoolVarP(&details, "details", "d", false, "Show details")

	return command
}
