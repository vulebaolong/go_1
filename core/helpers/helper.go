package helpers

import (
	"bufio"
	"core/model"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func PromptInput(label string) string {
	// bảng mã ASCII

	fmt.Print(label)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString(byte(10))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	input = strings.TrimSpace(input)
	return input
}

func ReadFileJson(expense *[]model.Expense, nameFile string) error {
	dataJsonRead, err := os.ReadFile(nameFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(dataJsonRead, expense)
	if err != nil {
		return err
	}

	return nil
}

func WriteFileJson(expense *[]model.Expense, nameFile string) error {
	dataJsonWrite, err := json.MarshalIndent(expense, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(nameFile, dataJsonWrite, 0600)
	if err != nil {
		return err
	}

	return nil
}
