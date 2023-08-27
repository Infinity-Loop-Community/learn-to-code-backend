package domain

import "fmt"

func GetSourceIPOutput(sourceIP string) string {
	if sourceIP == "" {
		return "Hello, world!\n"
	} else {
		return fmt.Sprintf("Hello, %s!\n", sourceIP)
	}
}
