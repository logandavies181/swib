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
	host bool
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
	Short: "A simple http server for debugging purposes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Listening on port: %s", port)
		
		if all {
			method = true
			path = true
			proto = true
			headers = true
			host = true
			body = true
			
		}
		if bothOutputs {
			stdout = true
			responseBody = true
		}
		
		// TODO handle errors here 
		http.HandleFunc("/", dumpHandler)
		http.ListenAndServe(":"+port, nil)
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)
	dumpCmd.Flags().StringVarP(&port, "port", "p", 8080, "Port to listen on")
	
	dumpCmd.Flags().BoolVarP(&method, "method", "m", false, "Print request method")
	dumpCmd.Flags().BoolVarP(&path, "path", "P", false, "Print request path")
	dumpCmd.Flags().BoolVarP(&proto, "proto", "V", false, "Print request http version")
	dumpCmd.Flags().BoolVarP(&headers, "headers", "H", false, "Print request headers")
	dumpCmd.Flags().BoolVarP(&body, "body", "b", true, "Print request body")
	dumpCmd.Flags().BoolVarP(&all, "all", "a", false, "Shortcut to print all")

	dumpCmd.Flags().BoolVarP(&stdout, "stdout", "s", true, "Print to stdout")
	dumpCmd.Flags().BoolVarP(&responseBody, "resp", "r", false, "Print to response body")
	dumpCmd.Flags().BoolVarP(&bothOutputs, "", "A", false, "Shortcut to response to stdout and response body")

}
