package repositories

import (
	"context"
	"errors"
	"fmt"
	"knovel/userprofile/domain/common"
	"knovel/userprofile/domain/entities"
)

type UserProfileRepository interface {
	GetUserInfo(ctx context.Context, userId string) (*entities.User, error)
	SignIn(ctx context.Context, email string, password string) (*entities.User, error)
}

type UserProfileRepositoryInstance struct {
	dbContext common.DbContext
	tableName string
}

var _ UserProfileRepository = (*UserProfileRepositoryInstance)(nil)

func NewUserProfileRepository(dbContext common.DbContext) UserProfileRepository {
	return &UserProfileRepositoryInstance{
		dbContext: dbContext,
		tableName: "userprofile",
	}
}

func (p *UserProfileRepositoryInstance) table(query string) string {
	return fmt.Sprintf(query, p.tableName)
}

// GetUserInfo implements UserProfileRepository.
func (repo *UserProfileRepositoryInstance) GetUserInfo(ctx context.Context, userId string) (*entities.User, error) {
	query := "SELECT id, email, first_name, last_name, deleted_at FROM %s WHERE id = $1"
	user := entities.User{}
	err := repo.dbContext.QueryRowContext(ctx, repo.table(query), userId).Scan(&user.Id, &user.Email,
		&user.FirstName, &user.LastName, &user.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Validate implements UserProfileRepository.
func (repo *UserProfileRepositoryInstance) SignIn(ctx context.Context, email string, password string) (*entities.User, error) {
	query := "SELECT id, email, first_name, last_name, password, deleted_at FROM %s WHERE email = $1"
	user := entities.User{}

	err := repo.dbContext.QueryRowContext(ctx, repo.table(query), email).Scan(&user.Id, &user.Email,
		&user.FirstName, &user.LastName, &user.Password, &user.DeletedAt)
	if err != nil {
		return nil, err
	}
	if !common.ValidatePassword(user.Password, password) {
		return nil, errors.New("password is invalid")
	}

	return &user, nil
}
