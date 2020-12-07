package repository

import "github.com/KoteiIto/wire-sample/pkg/model"

type IOnetimeTokenRepository interface {
	Issue(user model.User) (string, error)
	Verify(token string) error
}

type OnetimeTokenRepository struct {

}

var _ IOnetimeTokenRepository = (*OnetimeTokenRepository)(nil)

func NewOnetimeTokenRepository() *OnetimeTokenRepository {
	return &OnetimeTokenRepository{}
}

func (r OnetimeTokenRepository) Issue(user model.User) (string, error) {
	return "", nil
}

func (r OnetimeTokenRepository) Verify(token string) error {
	return nil
}