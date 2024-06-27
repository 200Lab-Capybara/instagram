package common

import "fmt"

func RecoverPanic(message string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf(message)
		}
	}()
}
