package services

import (
	"api/models"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPost(ctx context.Context, postID string) (*models.Post, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://jsonplaceholder.typicode.com/posts/"+postID, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var post models.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func GetAllPosts(ctx context.Context) ([]models.Post, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://jsonplaceholder.typicode.com/posts", nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var posts []models.Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
