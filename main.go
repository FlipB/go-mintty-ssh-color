package main

import "os";

import "flag";
import "regexp";

var beforeFlag = flag.String("before", "00FF00", "Color to set before parent process exits. Six digits, 0-F.")
var afterFlag = flag.String("after", "000000", "Color to set after parent process exits. Six digits, 0-F.")
var watchFlag = flag.Int("watch-pid", 0, "The pid to wait for before setting 'after' color.")


func main() {

	flag.Parse()
	if (!isValidColorCodeString(*beforeFlag)) {
		println("Error: 'before' value ("+*beforeFlag+") is not a valid color.");
		panic("Invalid input parameter")
	}
	if (!isValidColorCodeString(*afterFlag)) {
		println("Error: 'after' value ("+*afterFlag+") is not a valid color.");
		panic("Invalid input parameter")
	}
	if (*watchFlag == 0) {
		panic("Invalid value passed to --watch-pid ")
	}

	//ppid := os.Getppid()
	setMinttyBackground(*beforeFlag)

	waitForPidExit(*watchFlag)

	// process ppid has exited.
	setMinttyBackground(*afterFlag)
}

func waitForPidExit(pid int) {
	process, err := os.FindProcess(pid)
	if (err != nil) {
		panic(err)
	}
	_, err = process.Wait()
	if (err != nil) {
		panic(err)
	}
}

func isValidColorCodeString(colorCode string) bool {
	re, _ := regexp.Compile("(?i)[0-9a-f]{6}");
	return re.Match([]byte(colorCode))
}

func setMinttyBackground(colorCode string) {
	print("\x1b]11;#"+colorCode+"\x07"); // set background color
}
