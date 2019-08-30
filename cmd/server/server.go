package main

import "exp-chat/service"

func main() {
	newChat := service.NewChatService()
	print(newChat)
}
