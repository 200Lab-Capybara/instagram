package interfaces

// IUserRepository defines the interface for a user repository
type IUserRepository interface {
    CreateHashTagForPost(text, id string) string
}
