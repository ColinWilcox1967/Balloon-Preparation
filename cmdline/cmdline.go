package cmdline

import (
	"strconv"
	"bufio"
	"strings"
	"os"
    "fmt"
)

func ShowErrorOnConsole(err string) {
    fmt.Printf ("ERROR: %s\n", err)
}

func GetKeyPress() []byte {
    b := make([]byte, 1)
	os.Stdin.Read(b); // Ignore erro returns as its just for superficial use
    
    return b
}

func GetEntryFromCommandLine(prompt string) string {
    fmt.Printf("%s ", prompt)

    reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
    if err != nil {
        return ""
    }

    return strings.TrimSpace(line) // trim and trailing or leading whitespace
}

func ReadLimitValuesFromCommandLine() ([]string, error) {
    line := GetEntryFromCommandLine ("")
    strs := strings.Fields(line)  // Split string on space character into balloon limits slice

	return strs, nil
}

func CheckNatureOfLimitValues(strs []string) ([]int, error) {
 	nums := make([]int, len(strs))
    for i, strLimit := range strs {
        sizeLimit, err := strconv.Atoi(strLimit) // convert string to numeric integer, if possible
     
        // Supplied balloon size limit is non numeric.
        if err != nil {
            fmt.Printf ("balloon %d supplied with a non-numeric size limit (%s)", i+1, strs[i])
            return nil, err  
        }

        // Reject any none positive numbers
        if sizeLimit < 1 {
            fmt.Printf ("balloon %d, size limit is not a positive number (%d)", i+1, sizeLimit)
            return nil, err
        }

        // passed checks so added it to integer return slice
        nums[i] = sizeLimit
    }

    return nums, nil
}

// end of file
