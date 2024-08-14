package analyzer

import (
	Commands "MIAB_HM2/Comands"
	"errors"
	"strings"
)

func Analyzer(input string) (interface{}, error) {

	tokens := strings.Fields(input)

	if len(tokens) == 0 {
		return nil, errors.New("No command found")

	}

	switch tokens[0] {
	case "execute":
		//fmt.Println("Execute")
		Commands.ParseExecute(tokens[1:])
	default:
		return nil, errors.New("Command not found")
	}

	return nil, nil
}
