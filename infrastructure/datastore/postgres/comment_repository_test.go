package postgres

import (
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/stretchr/testify/mock"
)

type postRepositoryMock struct {
	db mock.Mock
}

func (m *postRepositoryMock) GetAll() ([]*model.Post, error) {
	args := m.db.Called()
	return args.Get(0).([]*model.Post)
}
