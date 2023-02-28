// TODO: Add native ngrok support.
// TODO: Better styling
// TODO: Add better prefixes for the urls
// TODO: Add a way to add a custom prefix to the urls
// TODO: Add a category of prefixes for the urls
// TODO: Add support for more than one url

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
)

var urls = make(map[string]string)

func main() {
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "style.css")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		longUrl := r.FormValue("url")
		shortUrl := generateShortUrl()

		urls[shortUrl] = longUrl
		fmt.Println(urls)

		fmt.Fprintf(w, "Short URL: %s", shortUrl)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Path[1:]
	if len(shortUrl) == 0 {
		http.ServeFile(w, r, "index.html")
		return
	}

	// Match valid short URLs
	validPath := regexp.MustCompile("^([a-zA-Z]+)-([a-zA-Z0-9]+)$")
	matches := validPath.FindStringSubmatch(shortUrl)
	if matches == nil {
		http.Error(w, "Invalid short URL format", http.StatusBadRequest)
		return
	}
	longUrl, ok := urls[shortUrl]
	if !ok {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, longUrl, http.StatusFound)
}

func generateShortUrl() string {
	prefixes := []string{"cute-animal-game", "funny-game", "work-funny-game", "nsfw-funny-game", "love-game"}
	prefix := prefixes[rand.Intn(len(prefixes))]
	suffix := randomString(8)
	return fmt.Sprintf("%s-%s", prefix, suffix)
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = charset[rand.Intn(len(charset))]
	}
	return string(bytes)
}
