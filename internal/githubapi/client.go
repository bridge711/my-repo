package githubapi

import (
	"context"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	r *resty.Client
}

// New は Github API用クライアントを作成します。
// baseURL は通常 "https://api.github.com" を指定します。
func New(baseURL string) *Client {
	c := resty.New().
		SetBaseURL(baseURL).
		SetTimeout(5*time.Second).
		SetHeader("Accept", "application/vnd.github+json")

	return &Client{r: c}
}

// User は最小限のフィールドだけ持つサンプル構造体です。
type User struct {
	Login string `json:"login"`
	ID    int64  `json:"id`
	Name  string `json:"name"`
}

func (c *Client) GetUser(ctx context.Context, username string) (*User, error) {
	resp, err := c.r.R().
		SetContext(ctx).
		SetResult(&User{}).
		Get("/users/" + username)

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("github api error: status=%d", resp.StatusCode())
	}
	return resp.Request().(*User), nil

}
