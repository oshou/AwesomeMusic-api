package presenter

import "github.com/oshou/AwesomeMusic-api/domain/model"

type UserPresenter interface {
	ResponseUsers(users []*model.User) []*model.User
	ResponseUsers(user *model.User) []*model.User
}
