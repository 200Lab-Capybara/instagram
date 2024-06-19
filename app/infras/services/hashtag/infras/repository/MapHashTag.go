package hashtagsql

import (
    "context"
    "database/sql"
    "fmt"
    "github.com/google/uuid"
    hashtagmodel "github.com/nghiatrann0502/instagram-clone/app/internals/services/hashtag/model"
)

type HashtagRepository struct {
    DB *sql.DB
}

// MapHashTag maps a hashtag to a specific post in the database.
func (repo *HashtagRepository) MapHashTag(ctx context.Context, hashtagPost *hashtagmodel.HashtagPost) (*hashtagmodel.HashtagPost, error) {
    if hashtagPost == nil {
        return nil, fmt.Errorf("hashtagPost cannot be nil")
    }

    // Prepare SQL statement for inserting or updating the post-hashtag relationship
    stmt := `
    INSERT INTO post_hashtags (hashtag_id, post_id) 
    VALUES (?, ?) 
    ON DUPLICATE KEY UPDATE 
    hashtag_id = VALUES(hashtag_id), 
    post_id = VALUES(post_id);
    `
    
    _, err := repo.DB.ExecContext(ctx, stmt, hashtagPost.Hashtag_ID, hashtagPost.Post_ID)
    if err != nil {
        return nil, fmt.Errorf("failed to map hashtag to post: %w", err)
    }

    return hashtagPost, nil
}
