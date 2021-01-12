package messages

import "fmt"

func ArgErrorMessage(p string) string {
	return fmt.Sprintf("There is no such Argumentsument in packet %s", p)
}
