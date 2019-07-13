package cmd

import (
	"fmt"
//	"net/http"
	"github.com/spf13/cobra"
//	flag "github.com/spf13/pflag"

)

var myFlag string

// curlCmd represents the curl command
var curlCmd = &cobra.Command{
	Use:   "curl",
	Short: "A command to copy some curl functionality where it is not available",
	Long: `Sometimes curl isn't available in a container. Or perhaps you're on windows...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(myFlag)
		fmt.Println(args[0])

	},
}

func init() {
	rootCmd.AddCommand(curlCmd)
	
	curlCmd.Flags().StringVarP(&myFlag, "myFlag", "m", "hello default","")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// curlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// curlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
