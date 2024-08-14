package Commands

import (
	"fmt"
	"regexp"
	"strings"
)

type MKDISK struct {
	size int
	unit string
	fit  string
	path string
}

func AnalyzeMkdisk(tokens []string) (*MKDISK, error) {
	cmd := &MKDISK{}

	args := strings.Join(tokens, " ")
	fmt.Println(args)

	re := regexp.MustCompile(`-size=\d+|-unit=[kKmM]|-fit=[bBfFwW]{2}|-path="[^"]+"|-path=[^\s]+`)
	fmt.Println(re.FindAllString(args, -1))
	return cmd, nil
}
