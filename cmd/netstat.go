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
var state []string
var quiet bool

var statesMap map[int64]string = map[int64]string{
	1:	"ESTABLISHED",
	2:	"SYN SENT",
	3:	"SYN RECEIVED",
	4:	"FIN WAIT 1",
	5:	"FIN WAIT 2",
	6:	"TIME WAIT",
	7:	"CLOSE",
	8:	"CLOSE WAIT",
	9:	"LAST ACKNOWLEDGEMENT",
	10:	"LISTENING",
	11:	"CLOSING",
	12:	"NEW SYN RECEIVED", 
}

var netstatCmd = &cobra.Command{
	Use:   "netstat",
	Short: "Dump listening ports",
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open("/proc/net/tcp")
		defer f.Close()
		if err != nil { fmt.Println (err) }
		
		if !quiet { fmt.Println("LOCAL\t\tREMOTE\t\tSTATE") }
		
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
		
			stateHex := wordScanner.Text()

			// State 0A is listening
			if all {
				fmt.Printf("%s\t%s\t%s\n",parseHexIp(localHex),parseHexIp(remHex),parseStateHex(stateHex))
			} else if stateHexInStates(stateHex)  {
				fmt.Printf("%s\t%s\t%s\n",parseHexIp(localHex),parseHexIp(remHex),parseStateHex(stateHex))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(netstatCmd)
	netstatCmd.Flags().BoolVarP(&all, "all", "a", false, "Print all states instead of just listening")
	netstatCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Print header or not")
	netstatCmd.Flags().StringSliceVarP(&state, "state", "s", []string{"0A"}, "States to print")

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

	var portString string
	if port == 0 { 
		portString = "*"
	} else {
		portString = fmt.Sprintf("%d",port)
	}
	return fmt.Sprintf("%s:%s",strings.Join(ip,"."),portString)
}

func parseStateHex(stateHex string) string {
	stateInt, err := strconv.ParseInt(stateHex,16,8)
	if err != nil { fmt.Println(err); return "unknown" }
	return statesMap[stateInt]
}

func stateHexInStates(stateHex string) bool {
	
	for _, s := range state {
		if stateHex == s {
			return true
		}
	}
	return false
}
