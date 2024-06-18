package main

import (
	"repository"
)
var postWithHashtag = []map[string]interface{}{
	{
		"ID":      1,
		"Content": "This is my first post",
		"Hashtags": []map[string]interface{}{
			{"ID": 1, "Tag": "#golang"},
			{"ID": 2, "Tag": "#programming"},
		},
	},
	{
		"ID":      2,
		"Content": "Learning Go is fun!",
		"Hashtags": []map[string]interface{}{
			{"ID": 3, "Tag": "#golang"},
			{"ID": 4, "Tag": "#learning"},
		},
	},
}


func InMemoryRepo() *InMemoryRepository {
	return &InMemoryRepository{}
}