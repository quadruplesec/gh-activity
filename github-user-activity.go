package main

import (
	"fmt"

	"io"

	"os"

	"log"

	"net/http"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("Usage: ./github-user-activity <username>")
	}

	var username string = os.Args[1]
	var request string = fmt.Sprintf("https://api.github.com/users/%s/events", username)

	// TODO : Check if err automatically handles non-existent users. If not, handle that

	resp, err := http.Get(request)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	fmt.Println(body)
}
