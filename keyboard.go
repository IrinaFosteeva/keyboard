package keyboard

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetFloat(keyWord string) (float64, error) {
	fmt.Printf("  Enter %s: ", keyWord)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	return value, nil
}

func GetString(keyWord string) (string, error) {
	fmt.Printf("  Enter %s: ", keyWord)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "0", err
	}

	input = strings.TrimSpace(input)
	return input, nil
}
