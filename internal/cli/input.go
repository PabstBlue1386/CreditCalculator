package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFloat(prompt string) float64 {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		number, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return number
		} else {
			fmt.Printf("Ввод некоректен!")
			fmt.Print(prompt)
		}

	}

}
