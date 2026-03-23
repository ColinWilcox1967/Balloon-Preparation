
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
}

//
// Block of unit tests to exercise the 'ShowErrorOnConsole()' function
//
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


//
// Block of unit tests to exercise the 'ReadLimitValueFromCommandLine()' function
//
func TestCheckNatureOfLimitValues(t *testing.T) {
	var testCases = []struct {
   		msg             []string
		values          []int
   		expectedStatus  bool // set to false if there is an error
}{
    {[]string{""}, []int{0}, false},			// Empty stlice is not allowed
	{[]string{"A"}, []int{0}, false},			// Dont allow non-numeric characters
	{[]string{"A","B"}, []int{0,0}, false},		// Multiple non-numeric characetrs no allowed
	{[]string{"1"}, []int{1}, true},			// Single number is permissable
	{[]string{"1","2"}, []int{1,2}, true},		// .. as is several numbers
	{[]string{"10","0"}, []int{10,0}, true},	// numbers can be zero
	{[]string{"5","-1"}, []int{5,-1}, true},	// and negative
	{[]string{"2","X"}, []int{2,0}, false},		// cannot allow mix and match with value types
	{[]string{"1.0"}, []int{1}, false},			// Non integer values not allowed	
	{[]string{"1.0", "1"}, []int{1,1}, false},  // All values must be integers
}

totalTests := len(testCases)
passedTests := 0
for _, test := range testCases {
        values, err := cmdline.CheckNatureOfLimitValues(test.msg)
		
		if test.expectedStatus != (err == nil) {

			// build errStr
			errStr := "Params: "
			for idx := 0; idx < len(test.msg);idx++ {
				errStr += test.msg[idx]
				errStr += " "
			}

			// just dump the value slice on error
			fmt.Printf ("Values: ")
			for value,_ := range(values) {
				fmt.Printf ("%d ", value)
			}
			fmt.Println()

            t.Errorf(errStr)
        } else
		{
			passedTests++
		}
    }

	showTestResults("\nTestCheckNatureOfLimitValues", passedTests, totalTests)
}

