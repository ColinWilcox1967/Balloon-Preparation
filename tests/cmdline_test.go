
package main
//cmdline_test.go

// Unit tests for publicly visible functions


import (
	"fmt"
    "testing"

	cmdline "../cmdline"
)

// helper functions
func showTestResults(testSetName string, passedTests, totalTests int) {
	fmt.Printf ("Test Set : %s, ", testSetName)

	fmt.Printf ("Total tests = %d, tests passed = %d.\n", totalTests, passedTests)

//	if (passedTests == totalTests) {
//		fmt.Printf("BLOCK PASSED.\n")
//	} else {
//		fmt.Printf("BLOCK FAILED.\n")
//	}
}

func TestShowErrorOnConsole(t *testing.T) {
    var testCases = []struct {
   	msg        string
    expectedLength int
}{
    {"", 0},
    {"Test message single line", len("Test message single line")},
    {"Test message -\n split line", len("Test message -\n split line")},
}

totalTests := len(testCases)
passedTests := 0
for _, test := range testCases {
        bytesWritten,_ := cmdline.ShowErrorOnConsole(test.msg)
		if bytesWritten != test.expectedLength {
			errStr := fmt.Sprintf("msg = %s, expected %d. Got %d.\n",
                 test.msg, test.expectedLength, bytesWritten)
            t.Errorf(errStr)
        } else
		{
			passedTests++
		}
    }

	showTestResults("\nTestShowErrorOnConsole", passedTests, totalTests)

}