package ps

import (
	"fmt"
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
	proc, err := diagnostic.ReadProcess()

	if err != nil {
		panic(err)
	}

	if len(args) > 0 {
		size := 0

		for i := 0; i < len(proc); i++ {
			if sieve(&proc[i], args) {
				proc[size] = proc[i]
				size++
			}
		}

		proc = proc[:size]
	}

	sort.Slice(proc, func(lhs int, rhs int) bool {
		lhsPID, _ := strconv.Atoi(proc[lhs].PID)
		rhsPID, _ := strconv.Atoi(proc[rhs].PID)

		return lhsPID < rhsPID
	})

	formatter := fmtutil.NewFormatter()

	if details {
		formatter.Write("USER", "PID", "PPID", "NAME", "PATH", "COMMAND")
		formatter.Align(1, fmtutil.RightAlign)
		formatter.Align(2, fmtutil.RightAlign)

		for _, each := range proc {
			formatter.Write(each.User, each.PID, each.PPID, each.Name, each.Path, each.CommandLine)
		}
	} else {
		formatter.Write("USER", "PID", "PPID", "COMMAND")
		formatter.Align(1, fmtutil.RightAlign)
		formatter.Align(2, fmtutil.RightAlign)

		for _, each := range proc {
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
