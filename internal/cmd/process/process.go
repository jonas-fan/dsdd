package process

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/jonas-fan/dsdd/pkg/dsa/diagnostic"
	"github.com/jonas-fan/dsdd/pkg/fmtutil"
	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) {
	proc, err := diagnostic.ReadProcess()

	if err != nil {
		panic(err)
	}

	sort.Slice(proc, func(lhs int, rhs int) bool {
		lhsPID, _ := strconv.Atoi(proc[lhs].PID)
		rhsPID, _ := strconv.Atoi(proc[rhs].PID)

		return lhsPID < rhsPID
	})

	formatter := fmtutil.NewFormatter()
	formatter.Write("USER", "PID", "PPID", "COMMAND")
	formatter.Align(1, fmtutil.RightAlign)
	formatter.Align(2, fmtutil.RightAlign)

	for _, each := range proc {
		formatter.Write(each.User, each.PID, each.PPID, each.CommandLine)
	}

	fmt.Println(formatter.String())
}

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "process",
		Short: "Display running process",
		Run:   run,
	}

	flags := command.Flags()
	flags.SetInterspersed(false)

	return command
}
