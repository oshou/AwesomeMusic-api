package presenter

type UserPresenterInterface interface{
	ResponseUsers(us []*model.User) []*mode.User
}
