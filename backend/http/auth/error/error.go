package error

import "fmt"

type AlreadyExistingError struct{}

func (e *AlreadyExistingError) Error() string {
	return fmt.Sprintln("The User Already Exists")
}
