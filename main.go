package main

import (
	"os"

	// package shortcuts
	cmdline "./cmdline" // functions related to command line processing
	balloon "./balloon" // main game logic
)

func main () {
	// Read the contents for command line ...
	strArgs, err := cmdline.ReadLimitValuesFromCommandLine()
	if err != nil {
		cmdline.ShowErrorOnConsole("problem parsing balloon size limits.")
		os.Exit(-1)
	}

	// ... and convert them into numbers, checking for validity
	balloonSizeLimits, err := cmdline.CheckNatureOfLimitValues(strArgs)
	if err != nil {
		cmdline.ShowErrorOnConsole(err.Error())
		os.Exit(-2)
	}

	// Start the game
	balloon.PlayGame (balloonSizeLimits)
	
	// just allow results to be seen on console before window closes
	cmdline.GetKeyPress()
}

// end of file
