package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFloat(prompt string) float64 {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		number, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return number
		} else {
			fmt.Printf("Ввод некоректен!\n")
			fmt.Print(prompt)
		}

	}
}

func ReadInt(prompt string) int64 {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)

	for {

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		number, err := strconv.ParseInt(input, 10, 64)
		if err == nil {
			return number
		} else {
			fmt.Printf("Ввод некоректен!\n")
			fmt.Print(prompt)
		}

	}

}
