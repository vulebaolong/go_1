package usecase

import (
	"core/helpers"
	"core/model"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"sync"
	"time"
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

	indexDelete := -1
	for index := range listExpense {
		if listExpense[index].Id == idInt {
			indexDelete = index
		}
	}

	fmt.Println("indexDelete", indexDelete)

	if indexDelete > -1 {
		listExpense = slices.Delete(listExpense, indexDelete, indexDelete+1)
	}

	fmt.Println("listExpense", listExpense)
	fmt.Println("listExpense", &listExpense)

	err = helpers.WriteFileJson(&listExpense, NAME_FILE_DATABASE)
	if err != nil {
		return err
	}

	return nil
}

func jobListExpense() []model.Expense {
	fmt.Println("jobListExpense bắt đầu")
	time.Sleep(5 * time.Second)
	listExpense, err := ReadExpense()
	if err != nil {
		return []model.Expense{}
	}
	return listExpense
}
func jobTotal() int {
	fmt.Println("jobTotal bắt đầu")
	time.Sleep(5 * time.Second)
	listExpense, err := ReadExpense()
	if err != nil {
		return 0
	}
	total := len(listExpense)
	return total
}

func NotGoroutine() {
	startTime := time.Now()

	list := jobListExpense()
	total := jobTotal()

	result := map[string]any{
		"list":  list,
		"total": total,
	}

	endTime := time.Since(startTime)

	fmt.Println("NotGoroutine", result, endTime)
}

func Goroutine() {
	startTime := time.Now()

	var wg sync.WaitGroup

	wg.Add(1)

	list := []model.Expense{}
	go func(wg *sync.WaitGroup) {
		defer func() {
			fmt.Println("List Done")
			wg.Done()
		}()

		list = jobListExpense()
	}(&wg)

	wg.Add(1)
	total := 0
	go func(wg *sync.WaitGroup) {
		defer func() {
			fmt.Println("Total Done")
			wg.Done()
		}()

		total = jobTotal()
	}(&wg)

	wg.Wait()

	result := map[string]any{
		"list":  list,
		"total": total,
	}

	endTime := time.Since(startTime)

	fmt.Println("NotGoroutine", result, endTime)
}

func GoroutineWaiGroupChannel() {
	startTime := time.Now()

	var wg sync.WaitGroup

	// cơ chế channel unbufer
	// phải có GỬI và NHẬN
	// GỬI và NHẬN cái nào tới trước thì sẽ bị block (ngừng đợi)
	// GỬI tới trước: NHẬN được chạy => mới mở block
	// NHẬN tới trước: GỬI được chạy => mới mở block
	channelList := make(chan []model.Expense)
	channelTotal := make(chan int)

	wg.Add(1)

	go func(wg *sync.WaitGroup, channelList chan []model.Expense) {
		defer func() {
			fmt.Println("List Done")
			wg.Done()
		}()

		list := jobListExpense()

		fmt.Println("jobListExpense GỬI")
		channelList <- list
	}(&wg, channelList)

	wg.Add(1)
	go func(wg *sync.WaitGroup, channelTotal chan int) {
		defer func() {
			fmt.Println("Total Done")
			wg.Done()
		}()

		total := jobTotal()

		fmt.Println("jobTotal GỬI")
		channelTotal <- total
	}(&wg, channelTotal)

	fmt.Println("NHẬN")
	list := <-channelList
	total := <-channelTotal

	wg.Wait()

	result := map[string]any{
		"list":  list,
		"total": total,
	}

	endTime := time.Since(startTime)

	fmt.Println("NotGoroutine", result, endTime)
}

// tận dụng cơ chế block của channel để đạt kết quả wg.Group()
func GoroutineChannel() {
	startTime := time.Now()

	// cơ chế channel unbufer
	// phải có GỬI và NHẬN (code phải được chạy)
	// GỬI và NHẬN cái nào tới trước thì sẽ bị block (ngừng đợi)
	// GỬI tới trước: NHẬN được chạy => mới mở block
	// NHẬN tới trước: GỬI được chạy => mới mở block
	channelList := make(chan []model.Expense)
	channelTotal := make(chan int)

	go func(channelList chan []model.Expense) {
		list := jobListExpense()

		fmt.Println("jobListExpense GỬI")
		channelList <- list
	}(channelList)

	go func(channelTotal chan int) {
		total := jobTotal()

		fmt.Println("jobTotal GỬI")
		channelTotal <- total
	}(channelTotal)

	fmt.Println("NHẬN")
	list := <-channelList
	total := <-channelTotal

	result := map[string]any{
		"list":  list,
		"total": total,
	}

	endTime := time.Since(startTime)

	fmt.Println("NotGoroutine", result, endTime)
}

func GoroutineWaiGroupBufferChannel() {
	startTime := time.Now()

	var wg sync.WaitGroup

	// cơ chế channel bufer
	// GỬI chỉ block khi bufer đã đầy
	// NHẬN chỉ block khi bufer còn rỗng
	channelList := make(chan []model.Expense, 1)
	channelTotal := make(chan int, 1)

	wg.Add(1)
	go func(wg *sync.WaitGroup, channelList chan []model.Expense) {
		defer func() {
			fmt.Println("List Done")
			wg.Done()
		}()

		list := jobListExpense()

		fmt.Println("jobListExpense GỬI")
		channelList <- list // [...] -> [list]
		channelList <- list // [list] ❌ không chạy KHÔNG HOÀN THÀNH ĐƯỢC DÒNG CODE => BLOCK
	}(&wg, channelList)

	wg.Add(1)
	go func(wg *sync.WaitGroup, channelTotal chan int) {
		defer func() {
			fmt.Println("Total Done")
			wg.Done()
		}()

		total := jobTotal()

		fmt.Println("jobTotal GỬI")
		channelTotal <- total
	}(&wg, channelTotal)

	wg.Wait()

	fmt.Println("NHẬN")
	list := <-channelList
	total := <-channelTotal

	result := map[string]any{
		"list":  list,
		"total": total,
	}

	endTime := time.Since(startTime)

	fmt.Println("NotGoroutine", result, endTime)
}

// racecondition
// một goroutine đang xem giá trị cũ, trong khi có 1 goroutine đang sửa giá trị đó
// go run --race .
func Racecondition() {
	// var mu sync.Mutex
	var wg sync.WaitGroup
	counter := 0

	totalRoutine := 100

	for i := 0; i < totalRoutine; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			// mu.Lock()
			counter = counter + 1
			// mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("Mong muốn counter =", 100)
	fmt.Println("Thực tế counter =", counter)
}
