package cmd

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
)

var (
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
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "A simple http server for debugging purposes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Listening on port: %s\n", port)
		
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
	dumpCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to listen on")
	
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

func dumpHandler (w http.ResponseWriter, r *http.Request) {
		
	var resp []string
	var topLine []string

	if method { topLine = append(topLine, r.Method) }
	if path { topLine = append(topLine, r.URL.Path) }
	if proto { topLine = append(topLine, r.Proto) }
	
	if len(topLine) > 0 {
		resp = append(resp, strings.Join(resp," "))
	}
	
	if headers {
		var headerList []string
		for k, v := range(r.Header) {
			// TODO make sure this formats it like the actual request
			headerList = append(headerList, fmt.Sprintf("%s: %s", k, strings.Join(v,", ")))
		}
		resp = append(resp, strings.Join(headerList, "\n"))
	}
	if body {
		bodyData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		resp = append(resp, string(bodyData))
	}
	
	respString := strings.Join(resp, "\n")
	if stdout {
		fmt.Println(respString)
	}
	if responseBody {
		fmt.Fprintln(w, respString)
	}
}
