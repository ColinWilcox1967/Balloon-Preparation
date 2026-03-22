package balloon

import (
	"fmt"

	cmdline "../cmdline"
)

var (
	totalScore int       // Final player game score
	balloonIndex int     // Index of the currently active balloon
	balloonSize int      // Expanded size of current balloon
	numberOfBalloons int // Total number of balloons to blow up
)

func setup() {
	totalScore = 0
	balloonSize = 0
}

func processCommand(command string, balloonSizeLimit int) (bool, bool) {
	
	switch (command) {
		case cmdline.CommandExpand:
			balloonSize++
			if balloonSize > balloonSizeLimit {
				// balloon has popped
				balloonSize = 0

				fmt.Printf ("%s %s\n", cmdline.PromptOutputPrefix, cmdline.PromptBurst)

				return true, false
			}
			
		case cmdline.CommandBank:
			// saved the current size, so move to next balloon
			totalScore += balloonSize
			balloonSize = 0
			return false, true		

		default:
			fmt.Printf("Unknown command - ignored (%s).\n", command)
	}

	return false, false // still playing with current balloon
}

func PlayGame(balloonSizeLimits []int) {
	setup()
	numberOfBalloons = len(balloonSizeLimits)

    var balloonIndex int

	for balloonIndex < numberOfBalloons {
		// Display input prompt and get the action
	
		cmd := cmdline.GetEntryFromCommandLine(cmdline.PromptInputPrefix)
		balloonBurst, balloonBanked := processCommand(cmd, balloonSizeLimits[balloonIndex])
			
		if balloonBurst {
			balloonIndex++
		}

		if balloonBanked {
			balloonIndex++
		}

	}

	// Show player's final score
	fmt.Printf ("%s %s%d\n", cmdline.PromptOutputPrefix, cmdline.PromptScore, totalScore)
}