package usermysql

import (
	"context"
	"github.com/google/uuid"
	usermodel "github.com/nghiatrann0502/instagram-clone/app/internals/services/user/model"
	"time"
)

func (m *mySQLStorage) CreateNewUser(ctx context.Context, user *usermodel.User) (*uuid.UUID, error) {
	stmt := `INSERT INTO users (id, first_name, last_name, email, password, salt, created_at, updated_at)
		values (?, ?, ?, ?, ?,?, ?, ?)`

	_, err := m.db.ExecContext(ctx, stmt,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Salt,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return nil, err
	}

	return &user.ID, nil
}
