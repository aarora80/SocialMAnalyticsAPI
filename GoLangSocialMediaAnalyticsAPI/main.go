package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"	
)

type Post struct {
    ID       string `json:"id"`
    Likes    int    `json:"likes"`
    Shares   int    `json:"shares"`
    Comments int    `json:"comments"`
}

var posts = struct {
    sync.RWMutex
    data map[string]Post
}{
    data: make(map[string]Post),
}

func addPost(w http.ResponseWriter, r *http.Request) {
    var post Post
    err := json.NewDecoder(r.Body).Decode(&post)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    posts.Lock()
    posts.data[post.ID] = post
    posts.Unlock()

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintln(w, "Post added successfully")
}

func getPostStats(w http.ResponseWriter, r *http.Request) {
    postID := r.URL.Query().Get("id")
    if postID == "" {
        http.Error(w, "Missing post ID", http.StatusBadRequest)
        return
    }

    posts.RLock()
    post, exists := posts.data[postID]
    posts.RUnlock()

    if !exists {
        http.Error(w, "Post not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(post)
}

func getAggregateStats(w http.ResponseWriter, r *http.Request) {
    var totalLikes, totalShares, totalComments int

    posts.RLock()
    for _, post := range posts.data {
        totalLikes += post.Likes
        totalShares += post.Shares
        totalComments += post.Comments
    }
    posts.RUnlock()

    stats := map[string]int{
        "total_likes":    totalLikes,
        "total_shares":   totalShares,
        "total_comments": totalComments,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(stats)
}

func main() {
    http.HandleFunc("/add", addPost)
    http.HandleFunc("/stats", getPostStats)
    http.HandleFunc("/aggregate", getAggregateStats)

    fmt.Println("Server started on :8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
