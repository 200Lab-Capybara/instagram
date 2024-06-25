package usermodel

type FollowUserPayload struct {
	UserID      string
	FollowingID string
}

func (FollowUserPayload) TableName() string {
	return "users"
}
