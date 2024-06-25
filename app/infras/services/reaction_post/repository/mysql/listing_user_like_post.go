package reactionpostmysql

func (store *mySQLStorage) ListingUserLikePost(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (bool, error) {
	db := store.db.GetConnection()
	data := reactionpostmodel.ReactionPost{}
	if err := db.Table(reactionpostmodel.ReactionPost{}.TableName()).
		Where("post_id = ? AND user_id = ?", postId, userId).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, common.ErrDB(err)
	}
	return true, nil
}