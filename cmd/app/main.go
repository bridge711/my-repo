package main

import (
    "context"
    "fmt"
    "os"
    "time"

    githubapi "github.com/bridge711/myrepo/internal/githubapi"
)

func main() {
    // 環境変数 GITHUB_USER があればそれを使う。無ければ "octocat"
    username := os.Getenv("GITHUB_USER")
    if username == "" {
        username = "octocat"
    }

    // クライアント作成（本番は GitHub のパブリック API）
    client := githubapi.New("https://api.github.com")

    // タイムアウト付きコンテキスト
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    user, err := client.GetUser(ctx, username)
    if err != nil {
        fmt.Println("error:", err)
        os.Exit(1)
    }

    fmt.Printf("login=%s id=%d name=%s\n", user.Login, user.ID, user.Name)
}
