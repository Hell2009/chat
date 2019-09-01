package main

import "chat/service"

func main() {
	newChat := service.NewChatService()
	print(newChat)
}
