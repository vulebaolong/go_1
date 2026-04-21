package main

import (
	"core/helpers"
	"core/usecase"
	"fmt"
)

func main() {
	for {
		// CRUD
		fmt.Println("=====Quản Lý Chi Tiêu=====")
		fmt.Println("1) Thêm chi tiêu")
		fmt.Println("2) Liệt kê chi tiêu")
		fmt.Println("3) Sửa chi tiêu")
		fmt.Println("4) Xoá chi tiêu")
		fmt.Println("5) Go-routine")
		fmt.Println("0) Thoát")
		fmt.Println("")

		choice := helpers.PromptInput(">> Chọn: ")

		switch choice {
		case "1":
			result, err := usecase.CreateExpense()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Tạo chi tiêu thành công ", result)
			}
		case "2":
			reuslt, err := usecase.ReadExpense()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Lấy danh sách chi tiêu thành công ")
				for _, value := range reuslt {
					fmt.Printf("%+v \n", value)
				}
			}
		case "3":
			reuslt, err := usecase.UpdateExpense()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Update chi tiêu thành công ", reuslt)
			}
		case "4":
			err := usecase.DeleteExpense()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Xoá chi tiêu thành công")
			}
		case "5":
			// parallel: song song
			// concurent: đồng thời
			// thread

			// usecase.NotGoroutine()
			// usecase.Goroutine()
			// usecase.GoroutineWaiGroupChannel()
			// usecase.GoroutineChannel()
			// usecase.GoroutineWaiGroupBufferChannel()
			usecase.Racecondition()
		case "0":
			fmt.Println("Tạm biệt!")
			return

		default:
			fmt.Println("Lụa chọn không phù hợp, vui lòng chọn laij")
		}

		// fmt.Println("choice", choice)
		// fmt.Println("choice", []byte(choice))
	}
}
