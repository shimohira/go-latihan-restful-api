package helper

import "fmt"

func PanicIfError(err error) {
	if err != nil {
		fmt.Println("panic donk")
		fmt.Println(err.Error())
		panic(err)
	}
}
