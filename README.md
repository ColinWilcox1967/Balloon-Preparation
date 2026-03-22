# Balloon-Preparation
Coding Problem Around Blowing up Balloons indicating propensity for risk.

# Build steps
(1) Open a terminal window. \
(2) Change into folder that contains the source code.\
(3) From command line, either run _build.bat_ or enter the following ...

_go build -o balloon.exe main.go_ 

and hit return

# Source Code
**main.go**      Main source code and entry point \
**build.bat**    Simple batch file to compile source and create executable.
\
**\cmdline**     Command line related package \
**\balloon**     Main ballon logic and command processing

# Running Unit Tests
A set of unit test has been included for each Golang package. In order to run these test perform the following
from the command line ... 

(1) Change into the _test_ folder \
(2) Run the command _go test -v_

This will execute the unit tests found in each of Golang files in this folder with a __Test_ suffix in order. 

(Written for Arctic Shores)
