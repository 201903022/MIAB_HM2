package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello, World!")
	fmt.Println("Ingrese el comando: ")
	fmt.Print("> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer la entrada:", err)

	}
	input = strings.TrimSpace(input)
	fmt.Println("Input: ", input)

}
