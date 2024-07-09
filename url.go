package main

import (
    "math/rand"
    "sync"
)

var (
    urls     = make(map[string]string)
    urlsLock sync.RWMutex
    letters  = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func generateShortURL() string {
    b := make([]rune, 6)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func storeURL(shortURL, longURL string) {
    urlsLock.Lock()
    defer urlsLock.Unlock()
    urls[shortURL] = longURL
}

func getURL(shortURL string) (string, bool) {
    urlsLock.RLock()
    defer urlsLock.RUnlock()
    longURL, exists := urls[shortURL]
    return longURL, exists
}

