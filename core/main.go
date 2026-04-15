package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		// CRUD
		fmt.Println("=====Quản Lý Chi Tiêu=====")
		fmt.Println("1) Thêm chi tiêu")
		fmt.Println("2) Liệt kê chi tiêu")
		fmt.Println("3) Sửa chi tiêu")
		fmt.Println("4) Xoá chi tiêu")
		fmt.Println("0) Thoát")
		fmt.Println("")
		fmt.Print(">>Chọn: ")

		// var reader *bufio.Reader

		// bảng mã ASCII

		reader := bufio.NewReader(os.Stdin)
		choice, err := reader.ReadString(byte(10))
		if err != nil {
			fmt.Println(err)
			return
		}
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Println("Xử lý 1")
		case "2":
			fmt.Println("Xử lý 2")
		case "3":
			fmt.Println("Xử lý 3")
		case "4":
			fmt.Println("Xử lý 5")
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
