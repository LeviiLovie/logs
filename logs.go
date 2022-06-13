package logs

import "os"

func LogOne() {
	os.Create("logs")
}
