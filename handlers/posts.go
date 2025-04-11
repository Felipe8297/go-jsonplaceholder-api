package handlers

import (
	"api/services"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func GetPostsHandler(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
	defer cancel()

	path := req.URL.Path

	if path == "/" {
		allPosts, err := services.GetAllPosts(ctx)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("500 - Internal Server Error: " + err.Error()))
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(allPosts)
		return
	}

	postID := path[1:]

	postData, err := services.GetPost(ctx, postID)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("500 - Internal Server Error: " + err.Error()))
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(postData)
}
