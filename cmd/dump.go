package cmd

import (
	"fmt"
	"net/http"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var {
	port string

	method bool
	path bool
	proto bool
	headers bool
	body bool
	all bool

	stdout bool
	responseBody bool
	bothOutputs bool
}

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dump called")
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)
	
	dumpCmd.

}
