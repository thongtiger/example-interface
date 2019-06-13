package main

import (
	"example-interface/service"
)

type IUser interface {
	S()
	M()
	L()
	XL()
	XXL()
}

func main() {
	var userService IUser = service.User{Context: "session_context"}
	userService.M()
	userService.L()
	userService.XXL()
}
