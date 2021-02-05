package gobuildercli

import (
	"fmt"
	"github.com/sadmansakib/gractical/common"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// buildExecuteCmd represents the buildexecute command
var buildExecuteCmd = &cobra.Command{
	Use:   "buildexecute",
	Short: "buildexecute copies the content",
	Long: `build execute copies the content of a folder to another folder.
	Also it can be used to compile binary`,
	Run: execCommand,
}

var src, dest,exe string

func init() {
	rootCmd.AddCommand(buildExecuteCmd)

	dir, err := os.Getwd()
	common.Check(err)

	buildExecuteCmd.Flags().StringVar(
		&dest,
		"builddir",
		dir,
		"provide path of a specific directory where content will be copied to",
	)

	buildExecuteCmd.Flags().StringVar(
		&exe, "exe", "", "compile the code and build a binary")

	buildExecuteCmd.Flags().StringVar(
		&src,
		"copydir",
		"",
		"provide path of a specific directory which will be copied",
	)

	err = buildExecuteCmd.MarkFlagRequired("copydir")
	common.Check(err)
}

func execCommand(_ *cobra.Command, _ []string) {
	if dest == src {
		fmt.Println("Operation can not be done. source and destination are same")
		return
	} else {
		err := common.CopyDir(src, dest)
		common.Check(err)
	}

	if len(exe) > 0{
		cmd := exec.Command("go", "build", "-o", exe)
		cmd.Dir = dest
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(out)
	}
}
