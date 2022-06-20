package linux

import (
	"fmt"
	"runtime"
)

func StartLinux() {
	fmt.Println(runtime.GOOS)
}
