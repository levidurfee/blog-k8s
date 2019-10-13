package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func random(length int) string {
	return stringWithCharset(length, charset)
}

func main() {
	http.HandleFunc("/", randomString)
	http.ListenAndServe(":8080", nil)
}

func randomString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", random(64))
}
