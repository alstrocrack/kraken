package concern

import "fmt"

// HTTP response function to separate concerns
func Response(status int, message string) (bool, error) {
	fmt.Println("response")
	return true, nil
}
