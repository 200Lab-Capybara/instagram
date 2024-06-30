package model

type CreateStory struct {
	ImageUrl string `json:"image_url"`
	Size     int    `json:"size"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Content  string `json:"content_story"`
	Storage  string `json:"storage_name"`
}