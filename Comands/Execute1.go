package Commands

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ParseExecute(tokens []string) error {
	// fmt.Println("Execute")
	fmt.Println("ParseExecute")
	args := strings.Join(tokens, " ")
	//fmt.Println("Args: ")
	//fmt.Println(args)
	re := regexp.MustCompile(`-path="[^"]+"|-path=[^\s]+`)
	matches := re.FindAllString(args, -1)
	//fmt.Println("Matches: ")
	//fmt.Println(matches)

	for _, match := range matches {
		//fmt.Println("Match: ", match)
		parts := strings.SplitN(match, "=", 2)
		//fmt.Println("Parts: ")
		//fmt.Println(parts)
		if len(parts) != 2 {
			return fmt.Errorf("Invalid token: %s", match)

		}
		key, value := strings.ToLower(parts[0]), parts[1]

		//Remove quotes from value if present
		if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
			value = strings.Trim(value, "\"")
		}

		//fmt.Println("Key: ", key)
		//fmt.Println("Value: ", value)

		switch key {
		case "-path":
			//fmt.Println("Path: ", value)
			if value == "" {
				return fmt.Errorf("Invalid path: %s", value)
			}
			if err := ReadFile(value); err != nil {
				return err
			}

		default:
			return fmt.Errorf("Invalid token: %s", match)
		}

	}
	return nil
}

func ReadFile(path string) error {
	fmt.Println("Reading file: ", path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("File not found: %s", path)
	}

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Error opening file: %s", path)
	}
	defer file.Close()
	fileContento := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		fileContento += scanner.Text() + "\n"
		AnalyzeLine(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Error reading file: %s", path)
	}

	//read file

	return nil
}

func AnalyzeLine(line string) error {
	//fmt.Println("Line: ", line)
	tokens := strings.Fields(line)
	//fmt.Println("Tokens: ")
	//fmt.Println(tokens)
	if len(tokens) == 0 {
		return fmt.Errorf("No command found")
	}

	switch tokens[0] {
	case "mkdisk":
		fmt.Println("Mkdisk token")
		CommandMkdisk()
	case "rep":
		fmt.Println("rep token")
		CommandRep()
	default:
		return fmt.Errorf("Command not found")
	}

	return nil
}
