package ps

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/dsa/diagnostic"
	"github.com/jonas-fan/dsdd/pkg/fmtutil"
	"github.com/spf13/cobra"
)

var details bool

func sieve(task *diagnostic.Task, filters []string) bool {
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

	if len(args) > 0 {
		size := 0

		for i := 0; i < len(tasks); i++ {
			if sieve(&tasks[i], args) {
				tasks[size] = tasks[i]
				size++
			}
		}

		tasks = tasks[:size]
	}

	sort.Slice(tasks, func(lhs int, rhs int) bool {
		lhsPID, _ := strconv.Atoi(tasks[lhs].PID)
		rhsPID, _ := strconv.Atoi(tasks[rhs].PID)

		return lhsPID < rhsPID
	})

	formatter := fmtutil.NewFormatter()

	if details {
		formatter.Write("USER", "PID", "PPID", "NAME", "PATH", "COMMAND")
		formatter.Align(1, fmtutil.RightAlign)
		formatter.Align(2, fmtutil.RightAlign)

		for _, each := range tasks {
			formatter.Write(each.User, each.PID, each.PPID, each.Name, each.Path, each.CommandLine)
		}
	} else {
		formatter.Write("USER", "PID", "PPID", "COMMAND")
		formatter.Align(1, fmtutil.RightAlign)
		formatter.Align(2, fmtutil.RightAlign)

		for _, each := range tasks {
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
