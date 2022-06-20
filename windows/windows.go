package windows

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/gen2brain/dlgs"
)

func WinStart() (os.FileInfo, string) {

	fmt.Println("Windows Detected !")
	file, _, err := dlgs.File("CHOOSE YOUR FILE", "", false)
	if err != nil {
		panic(err)
	}
	info, err := os.Stat(file)
	if err != nil {
		panic(info)
	}
	fmt.Println("File Choosen : ", file)
	fmt.Println("[ LOG WILL APPEAR AT BOTTOM ]")

	return info, file
}

func Notif(file string) error {
	err := beeep.Notify("ALERT", "The file : "+file+" has been opened !", "../warning.png")
	if err != nil {
		panic(err)
	}
	return err
}

func DecodeTime(fi os.FileInfo) time.Time {
	return time.Unix(0, fi.Sys().(*syscall.Win32FileAttributeData).LastAccessTime.Nanoseconds())
}

func Checking(fi os.FileInfo, file string, NewTime time.Time) {

	for {
		info, err := os.Stat(file)
		if err != nil {
			panic(err)
		}

		time.Sleep(1000)
		TimeOpened := DecodeTime(info)
		if TimeOpened != NewTime {
			fmt.Println("Opened at ", TimeOpened)
			Notif(file)
			Process(info, file)
		}
	}
}

func Process(info os.FileInfo, File string) {
	NewTime := DecodeTime(info)
	Checking(info, File, NewTime)
}
