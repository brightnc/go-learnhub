package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Title         string `json:"title"`
	Thumbnail_url string `json:"thumbnail_url"`
	Author_name   string `json:"author_name"`
	Author_url    string `json:"author_url"`
}

func Oembed(videoUrl string) (*Data, error) {
	var result *Data
	url := fmt.Sprintf("https://noembed.com/embed?url=%v", videoUrl)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to get noembed : %v", err)
		return nil, fmt.Errorf("ooopsss an error occurred, please try again")
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("failed to decode : %v", err)
		return nil, fmt.Errorf("ooopsss! can not decode")
	}
	return result, nil
}
