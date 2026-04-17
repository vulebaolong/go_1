package usecase

import (
	"core/helpers"
	"core/model"
	"errors"
	"fmt"
	"slices"
	"strconv"
)

const NAME_FILE_DATABASE = "database.json"

func CreateExpense() (*model.Expense, error) {
	amountString := helpers.PromptInput("Nhập amount: ")
	amountInt, err := strconv.Atoi(amountString)
	if err != nil {
		return nil, err
	}

	noteString := helpers.PromptInput("Nhập ghi chú: ")

	listExpens := []model.Expense{}

	err = helpers.ReadFileJson(&listExpens, NAME_FILE_DATABASE)
	if err != nil {
		return nil, err
	}

	id := 1
	if len(listExpens) > 0 {
		lastExpense := listExpens[len(listExpens)-1]
		id = lastExpense.Id + 1
	}

	expenseNew := model.Expense{
		Id:     id,
		Amount: amountInt,
		Note:   noteString,
	}

	listExpens = append(listExpens, expenseNew)

	// chuyển dữ liệu sang json
	err = helpers.WriteFileJson(&listExpens, NAME_FILE_DATABASE)
	if err != nil {
		return nil, err
	}
	// Chiều 1: từ giá trị -> ra địa chỉ
	// GIÁ TRỊ					|  KIỂU
	// 1						| int
	// "hihi"					| string
	// true / false				| boolean
	// expenseNew				| expense
	// 0x1875e7c4e080 địa chỉ   | (*expense) pointer

	// Chiều 2: từ địa chỉ -> giá trị

	fmt.Printf("%p \n", &expenseNew)

	return &expenseNew, nil
}

func ReadExpense() ([]model.Expense, error) {
	listExpense := []model.Expense{}

	err := helpers.ReadFileJson(&listExpense, NAME_FILE_DATABASE)
	if err != nil {
		return []model.Expense{}, err
	}

	return listExpense, nil
}

func UpdateExpense() (*model.Expense, error) {
	idString := helpers.PromptInput("Nhập id cần chỉnh sửa: ")
	if idString == "" {
		return nil, nil
	}
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		return nil, err
	}

	// lấy ra danh sách
	listExpense := []model.Expense{}

	helpers.ReadFileJson(&listExpense, NAME_FILE_DATABASE)

	var expenseExits *model.Expense

	for index := range listExpense {
		// fmt.Println("idInt", idInt, "index", index, "expense", listExpense[index])
		if listExpense[index].Id == idInt {
			expenseExits = &listExpense[index]
		}
	}

	if expenseExits == nil {
		return nil, errors.New("Không tìm thấy Expense")
	}

	labelAmount := fmt.Sprintf("Chỉnh sửa amount của id: %d thành: ", expenseExits.Id)
	amountString := helpers.PromptInput(labelAmount)

	if amountString != "" {
		amountInt, err := strconv.Atoi(amountString)
		if err != nil {
			return nil, err
		}
		expenseExits.Amount = amountInt
	}

	labelNote := fmt.Sprintf("Chỉnh sửa node của id: %s thành: ", expenseExits.Note)
	noteString := helpers.PromptInput(labelNote)

	if noteString != "" {
		expenseExits.Note = noteString
	}

	err = helpers.WriteFileJson(&listExpense, NAME_FILE_DATABASE)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func DeleteExpense() error {
	idString := helpers.PromptInput("Nhập id cần xoá: ")
	if idString == "" {
		return nil
	}
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	// lấy ra danh sách
	listExpense := []model.Expense{}

	helpers.ReadFileJson(&listExpense, NAME_FILE_DATABASE)

	indexDelete := 0
	for index := range listExpense {
		if listExpense[index].Id == idInt {
			indexDelete = index
		}
	}

	fmt.Println("indexDelete", indexDelete)

	if indexDelete > 0 {
		listExpense = slices.Delete(listExpense, indexDelete, indexDelete+1)
	}

	fmt.Println("listExpense", listExpense)
	fmt.Println("listExpense", &listExpense)

	err = helpers.ReadFileJson(&listExpense, NAME_FILE_DATABASE)
	if err != nil {
		return err
	}

	return nil
}
