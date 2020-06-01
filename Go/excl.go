package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

func main() {
	disk := os.Args[1]
	_, err := os.OpenFile(disk, os.O_EXCL, 0444)

	// can be used to check for kernel zfs
	if errors.Is(err, syscall.EBUSY) {
		fmt.Println("cannot open device, in use")
	} else if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("device is free")
	}

}
