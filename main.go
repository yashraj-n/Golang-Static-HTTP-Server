package main

import (
	"fmt"
	"net/http"
	"os"
)

var port string

func main() {
	fmt.Println("\n   _____       _                         _    _ _______ _______ _____                                 \n  / ____|     | |                       | |  | |__   __|__   __|  __ \\                                \n | |  __  ___ | |     __ _ _ __   __ _  | |__| |  | |     | |  | |__) |  ___  ___ _ ____   _____ _ __ \n | | |_ |/ _ \\| |    / _` | '_ \\ / _` | |  __  |  | |     | |  |  ___/  / __|/ _ \\ '__\\ \\ / / _ \\ '__|\n | |__| | (_) | |___| (_| | | | | (_| | | |  | |  | |     | |  | |      \\__ \\  __/ |   \\ V /  __/ |   \n  \\_____|\\___/|______\\__,_|_| |_|\\__, | |_|  |_|  |_|     |_|  |_|      |___/\\___|_|    \\_/ \\___|_|   \n                                  __/ |                                                               \n                                 |___/                                                                \n")
	fmt.Println("by yashraj-n (@yashraj-n) \n")
	static := http.FileServer(http.Dir("."))
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = "8080"
	}

	http.Handle("/", static)
	fmt.Println("[*] - Server Listening on Port " + port)
	fmt.Println("[*] - http://localhost:" + port + "/")
	http.ListenAndServe(":"+port, nil)
}
