package githubapi

import (
    "context"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
)

func TestGetUser_Success(t *testing.T) {
    // モックサーバ: /users/testuser に JSON を返す
    mux := http.NewServeMux()
    mux.HandleFunc("/users/testuser", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        _ = json.NewEncoder(w).Encode(User{
            Login: "testuser",
            ID:    12345,
            Name:  "Test User",
        })
    })
    srv := httptest.NewServer(mux)
    defer srv.Close()

    client := New(srv.URL)

    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    u, err := client.GetUser(ctx, "testuser")
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if u.Login != "testuser" || u.ID != 12345 || u.Name != "Test User" {
        t.Fatalf("unexpected user: %+v", u)
    }
}

func TestGetUser_NotFound(t *testing.T) {
    // 404 を返すモック
    mux := http.NewServeMux()
    mux.HandleFunc("/users/unknown", func(w http.ResponseWriter, r *http.Request) {
        http.NotFound(w, r)
    })
    srv := httptest.NewServer(mux)
    defer srv.Close()

    client := New(srv.URL)

    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    _, err := client.GetUser(ctx, "unknown")
    if err == nil {
        t.Fatal("expected error, got nil")
    }
}
