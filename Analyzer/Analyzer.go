package analyzer

import (
	"errors"
	"fmt"
	"strings"
)

func Analyzer(input string) (interface{}, error) {

	tokens := strings.Fields(input)

	if len(tokens) == 0 {
		return nil, errors.New("No command found")

	}

	switch tokens[0] {
	case "execute":
		fmt.Println("Execute")

	default:
		return nil, errors.New("Command not found")
	}

	return nil, nil
}
