package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://google.com", nil)
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("problem in receiving response", err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("issue in reading body", err)
		return
	}

	fmt.Println(string(body))
}
