package hashtagsql

import (
	"context"
	hashtagmodel "instagram/app/internals/services/hashtag/model"
)

func (m *mySQLStorage) GetHashTags(ctx context.Context, hashtagSlice []string) ([]*hashtagmodel.Hashtag, error) {
	var hashtag hashtagmodel.Hashtag
	var hashtags []*hashtagmodel.Hashtag // Slice to hold hashtags from the database

	// Check if the input slice is empty to prevent an invalid SQL query
	if len(hashtagSlice) == 0 {
		return []*hashtagmodel.Hashtag{}, nil // Return an empty slice if no hashtags are provided
	}

	// Executes a query to fetch records from the 'hashtags' table that are in the hashtagSlice
	if err := m.db.GetConnection().Table(hashtag.TableName()).Where("Hashtag IN ?", hashtagSlice).Find(&hashtags).Error; err != nil {
		return nil, err // Return nil and the error if the query fails
	}
	return hashtags, nil // Return the slice containing matched hashtags
}

//a, b, c
//[a, d, e]
//? in (a, d, e)
//{
//	"a: true"
//}
