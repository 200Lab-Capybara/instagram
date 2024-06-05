package hasher

type Hasher interface {
	Hash(password, salt string) (string, error)
	Verify(password, salt, hash string) bool
	GenSalt(len int8) (string, error)
}
