package usermysql

import (
	"context"
	"database/sql"
	"errors"
	usermodel "github.com/nghiatrann0502/instagram-clone/app/internals/services/user/model"
)

func (m *mySQLStorage) FindUserByEmail(ctx context.Context, email string) (*usermodel.User, error) {
	//TODO implement me
	var data usermodel.User

	stmt := `SELECT id, email, first_name, last_name, password, salt, role, created_at, updated_at FROM users WHERE email = ?`
	err := m.db.QueryRowContext(ctx, stmt, email).Scan(&data.ID, &data.Email, &data.FirstName, &data.LastName, &data.Password, &data.Salt, &data.Role, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, usermodel.UserNotFound
		}

		return nil, err
	}
	return &data, nil
}
