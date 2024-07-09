package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/shorten", shortenHandler)
    http.HandleFunc("/u/", redirectHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the URL Shortener!")
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
    var req struct {
        URL string `json:"url"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    shortURL := generateShortURL()
    storeURL(shortURL, req.URL)

    res := struct {
        ShortURL string `json:"short_url"`
    }{
        ShortURL: fmt.Sprintf("http://localhost:8080/u/%s", shortURL),
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
    shortURL := r.URL.Path[len("/u/"):]

    if longURL, exists := getURL(shortURL); exists {
        http.Redirect(w, r, longURL, http.StatusFound)
    } else {
        http.NotFound(w, r)
    }
}
