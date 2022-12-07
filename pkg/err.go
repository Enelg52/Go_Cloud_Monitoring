package pkg

import "fmt"

// ErrPrint print the error
func ErrPrint(err error) {
	if err != nil {
		fmt.Printf("[!] %s", err)
	}
}
