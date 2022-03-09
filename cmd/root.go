package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

var delay time.Duration
var times int32

var rootCmd = &cobra.Command{
	Use:   "jfc",
	Short: "Jesus Fucking Christ - retry it",
	Run: func(cmd *cobra.Command, args []string) {
		lastErrStatus := -1
		for times > 0 {
			var c *exec.Cmd
			if len(args) > 1 {
				c = exec.Command(args[0], args[1:]...)
			} else {
				c = exec.Command(args[0])
			}
			c.Stdout = os.Stdout
			c.Stdin = os.Stdin
			c.Stderr = os.Stderr

			err := c.Run()
			lastErrStatus = c.ProcessState.ExitCode()
			if err != nil || lastErrStatus != 0 {
				<-time.After(delay)
				times--
			} else {
				if err != nil {
					fmt.Fprintf(os.Stderr, "%+v", err)
				}
				os.Exit(0)
			}
		}

		os.Exit(lastErrStatus)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Int32VarP(&times, "times", "t", 0, "How many fucking times to retry?")
	rootCmd.Flags().DurationVarP(&delay, "delay", "d", 0, "How long to fucking wait? (go duration format)")

}
