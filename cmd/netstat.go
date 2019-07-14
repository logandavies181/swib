package cmd

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"strconv"

	"github.com/spf13/cobra"
)

// var all bool from root

var netstatCmd = &cobra.Command{
	Use:   "netstat",
	Short: "Dump listening ports",
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open("/proc/net/tcp")
		defer f.Close()
		if err != nil { fmt.Println (err) }
		
		lineScanner := bufio.NewScanner(f)
		lineScanner.Scan() // first line just lists the columns
		for lineScanner.Scan() {
			wordScanner := bufio.NewScanner(strings.NewReader(lineScanner.Text()))
			wordScanner.Split(bufio.ScanWords)
			wordScanner.Scan()	// first entry isn't interesting
			wordScanner.Scan()
		
			localHex := wordScanner.Text()
			wordScanner.Scan()
	
			remHex := wordScanner.Text()
			wordScanner.Scan()
			// State 0A is listening
			if all {
				fmt.Println(parseHexIp(localHex),parseHexIp(remHex))
			} else {
				if wordScanner.Text() == "0A" {
					fmt.Println(parseHexIp(localHex),parseHexIp(remHex))
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(netstatCmd)
	netstatCmd.Flags().BoolVarP(&all, "all", "a", false, "Print all states instead of just listening")

}

func parseHexIp(hex string) string {
	// expecting ip:port in weird format from /proc/net/tcp
	stringParts := strings.Split(hex, ":")
	var ip []string
	port, err := strconv.ParseInt(stringParts[1],16,64)
	if err != nil { fmt.Println (err) }
	
	// Read through part
	for i := 3; i+1 > 0; i-- {
		strHexRep := stringParts[0][i*2:i*2+2]
		strHex, err := strconv.ParseInt(strHexRep,16,32)
		if err != nil { fmt.Println(err) }
		ip = append(ip,fmt.Sprintf("%d",strHex))
	}
	return fmt.Sprintf("%s:%s",strings.Join(ip,"."),fmt.Sprintf("%d",port))
}
