package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortenedURL string    `json:"shortened_url"`
	CreatedAt    time.Time `json:"created_at"`
}

var urlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	fmt.Println("Hashing URL:", hasher)
	data := hasher.Sum(nil)
	fmt.Println("Hashed data:", data)
	hash := hex.EncodeToString(data)
	fmt.Println("Hex encoded hash:", hash)
	fmt.Println("Fianl String:", hash[0:8])
	return hash[0:8]
}

func main() {
	fmt.Println("Hello, World!")
	OriginalURL := "https://jotx19.vercel.app/"
	generateShortURL(OriginalURL)
}
