package main

import (
	"os"
)

func main() {
	API_ACCOUNT_KEY := os.Getenv("API_ACCOUNT_KEY")
	print(API_ACCOUNT_KEY)
}
